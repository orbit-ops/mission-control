package providers

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"

	"github.com/orbit-ops/launchpad-core/ent"
)

type ProviderCommand string

const (
	CreateAccess ProviderCommand = "create"
	RemoveAccess ProviderCommand = "remove"
)

type Provider interface {
	CreateAccess(ctx context.Context, token string, rocket *ent.Rocket, req *ent.Request) error
	RemoveAccess(ctx context.Context, token string, rocket *ent.Rocket, req *ent.Request) error
}

type ProviderConfig struct {
	ManagerExecutable string
	ApiUrl            string
}

type BaseProvider struct{}

func (BaseProvider) EncodeRocketConfig(payload any) (string, error) {
	bs, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	return b64.StdEncoding.EncodeToString(bs), nil
}
