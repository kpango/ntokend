package ntokend

import "time"

// Option represents a functional options pattern interface
type Option func(*token) error

var (
	defaultOpts = []Option{
		FailureSleepDuration(time.Second),
		TokenExpiration(time.Second * 30),
	}
)

// TokenFilePath represents a functional options pattern setter method to set the token file path value
func TokenFilePath(path string) Option {
	return func(tok *token) error {
		tok.tokenFilePath = path
		return nil
	}
}

// EnableValidate represents a functional options pattern setter method to enable validate token flag value
func EnableValidate() Option {
	return func(tok *token) error {
		tok.validateToken = true
		return nil
	}
}

// DisableValidate represents a functional options pattern setter method to disable validate token flag value
func DisableValidate() Option {
	return func(tok *token) error {
		tok.validateToken = false
		return nil
	}
}

// TokenExpiration represents a functional options pattern setter method to set the token expiration period
func TokenExpiration(dur time.Duration) Option {
	return func(tok *token) error {
		if dur > time.Millisecond {
			tok.tokenExpiration = dur
		}
		return nil
	}
}

// RefreshDuration represents a functional options pattern setter method to set the token refresh duration
func RefreshDuration(dur time.Duration) Option {
	return func(tok *token) error {
		tok.refreshDuration = dur
		return nil
	}
}

// FaiFailureSleepDuration represents a functional options pattern setter method to set the token fetch failure retry sleep duration
func FailureSleepDuration(dur time.Duration) Option {
	return func(tok *token) error {
		tok.failureSleepDuration = dur
		return nil
	}
}

// AthenzDomain represents a functional options pattern setter method to set the domain name of the token builder
func AthenzDomain(domain string) Option {
	return func(tok *token) error {
		tok.athenzDomain = domain
		return nil
	}
}

// ServiceName represents a functional options pattern setter method to set the service name of the token builder
func ServiceName(name string) Option {
	return func(tok *token) error {
		tok.serviceName = name
		return nil
	}
}

// KeyVersion represents a functional options pattern setter method to set the key version of the token builder
func KeyVersion(ver string) Option {
	return func(tok *token) error {
		tok.keyVersion = ver
		return nil
	}
}

// KeyData represents a functional options pattern setter method to set the key data of the token builder
func KeyData(keyData []byte) Option {
	return func(tok *token) error {
		tok.keyData = keyData
		return nil
	}
}

// Hostname represents a functional options pattern setter method to set the hostname of the token builder
func Hostname(hostname string) Option {
	return func(tok *token) error {
		tok.hostname = hostname
		return nil
	}
}

// IPAddr represents a functional options pattern setter method to set the IP address of the token builder
func IPAddr(ipAddr string) Option {
	return func(tok *token) error {
		tok.ipAddr = ipAddr
		return nil
	}
}
