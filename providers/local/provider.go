package local

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/providers"
)

type LocalProvider struct {
	providers.BaseProvider

	conf *providers.ProviderConfig
}

func NewLocalProvider(c *providers.ProviderConfig) (*LocalProvider, error) {
	return &LocalProvider{
		conf: c,
	}, nil
}

func (lp *LocalProvider) CreateAccess(ctx context.Context, token string, rocket *ent.Rocket, req *ent.Request) error {
	return lp.runJob(ctx, providers.CreateAccess, token, rocket, req)
}

func (lp *LocalProvider) RemoveAccess(ctx context.Context, token string, rocket *ent.Rocket, req *ent.Request) error {
	return lp.runJob(ctx, providers.RemoveAccess, token, rocket, req)
}

func (lp *LocalProvider) runJob(ctx context.Context, command providers.ProviderCommand, token string, rocket *ent.Rocket, req *ent.Request) error {
	rc, err := lp.EncodeRocketConfig(req)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "run", rocket.Image)
	cmd.Env = []string{
		fmt.Sprintf("LAUNCHPAD_API_URL=%s", lp.conf.ApiUrl),
		fmt.Sprintf("LAUNCHPAD_TOKEN=%s", token),
		fmt.Sprintf("LAUNCHPAD_CONFIG=%s", rc),
		fmt.Sprintf("LAUNCHPAD_COMMAND=%s", string(command)),
	}

	return cmd.Run()
}
