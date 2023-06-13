package ntokend

import (
	"errors"
	"fmt"
)

var (
	// ErrTokenNotFound represents a error the the token is not found
	ErrTokenNotFound       = errors.New("error:\ttoken not found")
	ErrDomainNotFound      = errors.New("error:\tno domain in token")
	ErrServiceNameNotFound = errors.New("error:\tno service name in token")
	ErrSignatureNotFound   = errors.New("error:\tno signature in token")
	ErrTokenExpired        = errors.New("error:\ttoken has expired")
	ErrTokenBuilder        = func(dom, svc, kv string, err error) error {
		return Wrapf(err, "failed to create ZMS SVC Token Builder\nAthenzDomain:\t%s\nServiceName:\t%s\nKeyVersion:\t%s", dom, svc, kv)
	}
	ErrInvalidToken = func(err error) error {
		return Wrap(err, "invalid server identity token")
	}
	Wrap = func(err error, msg string) error {
		if err != nil {
			if msg != "" && err.Error() != msg {
				return fmt.Errorf("%s: %w", msg, err)
			}
			return err
		}
		return New(msg)
	}
)
