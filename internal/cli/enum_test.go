package cli

import (
	"testing"
)

func TestReadEnum_String(t *testing.T) {
	tests := []struct {
		name string
		r    ReadEnum
		want string
	}{
		{
			name: "Test case 1: String of Pods",
			r:    Pods,
			want: "pods",
		},
		{
			name: "Test case 2: String of Nodes",
			r:    Nodes,
			want: "nodes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("ReadEnum.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadEnum_Type(t *testing.T) {
	r := Pods
	want := "enum"
	if got := r.Type(); got != want {
		t.Errorf("ReadEnum.Type() = %v, want %v", got, want)
	}
}

func TestReadEnum_Values(t *testing.T) {
	want := []string{"pods", "nodes"}
	var readEnum ReadEnum
	if got := readEnum.Values(); !equal(got, want) {
		t.Errorf("ReadEnum.Values() = %v, want %v", got, want)
	}
}

func TestReadEnum_Set(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "Test case 1: Valid resource type 'pods'",
			value:   "pods",
			wantErr: false,
		},
		{
			name:    "Test case 2: Valid resource type 'nodes'",
			value:   "nodes",
			wantErr: false,
		},
		{
			name:    "Test case 3: Invalid resource type",
			value:   "invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReadEnum("")
			if err := r.Set(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("ReadEnum.Set() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
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
