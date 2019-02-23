package ntokend

import (
	"reflect"
	"testing"
	"time"
)

func TestTokenFilePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokenFilePath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnableValidate(t *testing.T) {
	tests := []struct {
		name string
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnableValidate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnableValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisableValidate(t *testing.T) {
	tests := []struct {
		name string
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DisableValidate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisableValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenExpiration(t *testing.T) {
	type args struct {
		dur time.Duration
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokenExpiration(tt.args.dur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenExpiration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefreshDuration(t *testing.T) {
	type args struct {
		dur time.Duration
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RefreshDuration(tt.args.dur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefreshDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAthenzDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AthenzDomain(tt.args.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AthenzDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ServiceName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyVersion(tt.args.ver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyData(t *testing.T) {
	type args struct {
		keyData []byte
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyData(tt.args.keyData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHostname(t *testing.T) {
	type args struct {
		hostname string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hostname(tt.args.hostname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hostname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPAddr(t *testing.T) {
	type args struct {
		ipAddr string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IPAddr(tt.args.ipAddr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IPAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
