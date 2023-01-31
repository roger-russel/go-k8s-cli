package config

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvs(t *testing.T) {
	tests := []struct {
		name           string
		apiAddress     string
		candidateName  string
		expectedResult string
	}{
		{
			name:           "Test case 1: With default values",
			apiAddress:     "",
			candidateName:  "John Doe",
			expectedResult: ":80",
		},
		{
			name:           "Test case 2: With defined values",
			apiAddress:     ":8000",
			candidateName:  "Jane Doe",
			expectedResult: ":8000",
		},
		{
			name:           "Test case 3: Missing required env",
			apiAddress:     "",
			candidateName:  "",
			expectedResult: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt := tt
			os.Unsetenv("API_ADDRESS")
			os.Unsetenv("CANDIDATE_NAME")

			if tt.apiAddress != "" {
				os.Setenv("API_ADDRESS", tt.apiAddress)
			}

			if tt.candidateName != "" {
				os.Setenv("CANDIDATE_NAME", tt.candidateName)
			}

			ctx := context.Background()
			err := LoadEnvs(ctx)
			envs := Envs

			if tt.expectedResult == "" {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, envs.APIAddress)
				assert.Equal(t, tt.candidateName, envs.CandidateName)
			}
		})
	}
}
