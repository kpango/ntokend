package ntokend

import (
	"context"
	"io/ioutil"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/kpango/fastime"
	"github.com/kpango/glg"
	"github.com/yahoo/athenz/libs/go/zmssvctoken"
)

// TokenService represents a interface for user to get the token, and automatically update the token
type TokenService interface {
	StartTokenUpdater(context.Context) TokenService
	Update() error
	TokenExists() bool
	GetTokenProvider() TokenProvider
}

type token struct {
	tokenFilePath   string
	token           *atomic.Value
	validateToken   bool
	tokenExpiration time.Duration
	refreshDuration time.Duration
	builder         zmssvctoken.TokenBuilder

	// token builder parameters
	athenzDomain string
	serviceName  string
	keyVersion   string
	keyData      []byte
	hostname     string
	ipAddr       string
}

type rawToken struct {
	domain     string
	name       string
	signature  string
	expiration time.Time
}

// TokenProvider represents a token provider function to get the role token
type TokenProvider func() (string, error)

// New return TokenService
// This function will initialize the TokenService object with the tokenOptions
func New(opts ...Option) (TokenService, error) {
	tok := &token{
		token: new(atomic.Value),
	}

	for _, to := range opts {
		if err := to(tok); err != nil {
			return nil, err
		}
	}

	// create token builder
	tokBuilder, err := zmssvctoken.NewTokenBuilder(tok.athenzDomain, tok.serviceName, tok.keyData, tok.keyVersion)
	if err != nil || tokBuilder == nil {
		return nil, ErrTokenBuilder(tok.athenzDomain, tok.serviceName, tok.keyVersion, err)
	}

	if tok.hostname != "" {
		tokBuilder.SetHostname(tok.hostname)
	}

	if tok.ipAddr != "" {
		tokBuilder.SetIPAddress(tok.ipAddr)
	}

	tok.builder = tokBuilder

	return tok, nil
}

// StartTokenUpdater return TokenService
// This function will start a goroutine to update the token periodically, and store the token into memory
func (t *token) StartTokenUpdater(ctx context.Context) TokenService {
	go func() {
		var err error
		err = t.Update()
		fch := make(chan struct{})
		if err != nil {
			glg.Error(err)
			fch <- struct{}{}
		}

		ticker := time.NewTicker(t.refreshDuration)
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-fch:
				err = t.Update()
				if err != nil {
					glg.Error(err)
					time.Sleep(time.Second)
					fch <- struct{}{}
				}
			case <-ticker.C:
				err = t.Update()
				if err != nil {
					glg.Error(err)
					fch <- struct{}{}
				}
			}
		}
	}()
	return t
}

// GetTokenProvider returns a function pointer to get the token.
func (t *token) GetTokenProvider() TokenProvider {
	return t.getToken
}

// GetToken return a token string or error
// This function is thread-safe. This function will return the token stored in the atomic variable, or return the error when the token is not initialized or cannot be generated
func (t *token) getToken() (string, error) {
	tok := t.token.Load()
	if tok == nil {
		return "", ErrTokenNotFound
	}
	return tok.(string), nil
}

func (t *token) TokenExists() bool {
	_, err := t.getToken()
	return err == nil
}

// loadToken return a n-token string, or error
// This function return n-token, which is generated with the token builder. If the ntoken_path is set in the yaml (Copper Argos),
// this function will directly return the token file content.
// If ntoken_path is not set (k8s secret), the builder will read the private key from environment variable (private_key_env_name), and generate and sign a new token and return.
// This function can also validate the token generated or read. If validate_token flag is on, this function will verify the token first before this function return.
func (t *token) loadToken() (ntoken string, err error) {
	if t.tokenFilePath == "" {
		// k8s secret
		t.builder.SetExpiration(t.tokenExpiration)

		ntoken, err = t.builder.Token().Value()
		if err != nil {
			return "", err
		}

	} else {
		// Copper Argos
		var tok []byte
		tok, err = ioutil.ReadFile(t.tokenFilePath)
		if err != nil {
			return "", err
		}

		ntoken = strings.TrimRight(*(*string)(unsafe.Pointer(&tok)), "\r\n")
	}

	if t.validateToken {
		err = newRawToken(ntoken).isValid()
		if err != nil {
			return "", ErrInvalidToken(err)
		}
	}
	return ntoken, nil
}

// Update will load the ntoken and set it to the cache.
func (t *token) Update() error {
	token, err := t.loadToken()
	if err != nil {
		return err
	}
	t.setToken(token)
	return nil
}

func (t *token) setToken(token string) {
	t.token.Store(token)
}

// newRawToken returns the rawToken pointer.
// This function parse the token string, and transform to rawToken struct.
func newRawToken(token string) *rawToken {
	t := new(rawToken)
	for _, field := range strings.Split(token, ";") {
		parts := strings.SplitN(field, "=", 2)
		if len(parts) != 2 {
			continue
		}

		switch parts[0] {
		case "d": // Domain
			t.domain = parts[1]
		case "n": // Name
			t.name = parts[1]
		case "s": // Signature
			t.signature = parts[1]
		case "e": // Expiration
			parsed, err := strconv.ParseInt(parts[1], 0, 64)
			if err != nil {
				t.expiration = fastime.Now().Add(time.Second * 30)
			} else {
				t.expiration = time.Unix(parsed, 0)
			}
		}
	}
	return t
}

// isValid returns error from validating the rawToken struct.
func (r *rawToken) isValid() error {
	switch {
	case r.domain == "":
		return ErrDomainNotFound
	case r.name == "":
		return ErrServiceNameNotFound
	case r.signature == "":
		return ErrSignatureNotFound
	case r.expiration.Before(fastime.Now()):
		return ErrTokenExpired
	}
	return nil
}
