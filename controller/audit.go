package controller

import (
	"context"
	"fmt"

	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/hook"
)

func (ac *AccessController) setAuditHooks() {
	ac.client.Audit.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AuditFunc(func(ctx context.Context, am *ent.AuditMutation) (ent.Value, error) {
			return nil, fmt.Errorf("audit entries cannot be " + am.Op().String())
		})
	}, ent.OpDelete|ent.OpUpdate))
}
