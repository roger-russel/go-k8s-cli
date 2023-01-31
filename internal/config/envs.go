package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Environments struct {
	APIAddress    string `env:"API_ADDRESS,default=:80"`
	CandidateName string `env:"CANDIDATE_NAME,required"`
}

var Envs Environments

func LoadEnvs(ctx context.Context) error {
	if err := envconfig.Process(ctx, &Envs); err != nil {
		return err
	}
	return nil
}
