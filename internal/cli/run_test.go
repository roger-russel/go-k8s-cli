package cli

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_setupClientset(t *testing.T) {
	os.Setenv("FAKE_EXIT", "true")

	type args struct {
		cmd *cobra.Command
	}
	tests := []struct {
		name     string
		args     args
		wantExit bool
	}{
		{
			name: "simple",
			args: args{
				cmd: func() *cobra.Command {
					var kconfig string
					cmd := &cobra.Command{}
					kubeconfigFlag(cmd, &kconfig)
					return cmd
				}(),
			},
			wantExit: false,
		},
		{
			name: "flag error",
			args: args{
				cmd: &cobra.Command{},
			},
			wantExit: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(t *testing.T, wantExit bool) {
				if r := recover(); r != nil {
					assert.True(t, wantExit, "it should not call core.Exit")
				}
			}(t, tt.wantExit)

			if tt.args.cmd != nil {
				tt.args.cmd.Execute()
			}

			_ = setupClientset(tt.args.cmd)
		})
	}
}
