package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v3"
	"testing"
)

func TestGetCommand(t *testing.T) {
	ctx := context.TODO()
	GetCLI(
		ctx, func(_ context.Context, cmd *cli.Command) error {
			assert.Equal(t, "global", cmd.String("scope"))
			assert.Equal(t, "TEST", cmd.String("mode"))
			assert.Equal(t, "vim", cmd.String("editor"))
			return nil
		},
	).Run(ctx, []string{"mushi", "--s", "global", "--m", "TEST", "--e", "vim"})
}
