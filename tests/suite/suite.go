package suite

import (
	"context"
	"sso/internal/config"
	"testing"

	ssov1 "github.com/Guitarkeepsme/protos/gen/go/sso"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	return context.TODO(), &Suite{}
}
