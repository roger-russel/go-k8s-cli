package k8s

import (
	"testing"
)

func TestAuthType_Set(t *testing.T) {
	tests := []struct {
		name    string
		a       AuthType
		value   string
		wantErr bool
	}{
		{
			name:    "valid inClusterConfig value",
			value:   string(InClusterConfig),
			wantErr: false,
		},
		{
			name:    "valid buildConfigFromFlags value",
			value:   string(BuildConfigFromFlags),
			wantErr: false,
		},
		{
			name:    "invalid value",
			value:   "invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.Set(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("AuthType.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got := tt.a.String(); got != tt.value {
				t.Errorf("AuthType.String() = %v, want %v", got, tt.value)
			}
		})
	}
}

func TestAuthType_Values(t *testing.T) {
	want := []string{string(InClusterConfig), string(BuildConfigFromFlags)}
	var a AuthType
	if got := a.Values(); !equal(got, want) {
		t.Errorf("AuthType.Values() = %v, want %v", got, want)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
