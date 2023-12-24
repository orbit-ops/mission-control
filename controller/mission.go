package controller

import (
	"context"
	"errors"

	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/hook"
)

func (ac *AccessController) setMissionHooks() {
	// possible approvers must be at least as big as minimum approvers
	ac.client.Mission.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.MissionFunc(func(ctx context.Context, am *ent.MissionMutation) (ent.Value, error) {
			min_approvers, _ := am.MinApprovers()
			possible_approvers, _ := am.PossibleApprovers()
			if min_approvers > len(possible_approvers) {
				return nil, errors.New("amount of possible approvers must not be smaller than minimum approvers")
			}
			return next.Mutate(ctx, am)
		})
	}, ent.OpCreate|ent.OpUpdate))

	ac.client.Mission.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.MissionFunc(func(ctx context.Context, am *ent.MissionMutation) (ent.Value, error) {
			min_approvers, _ := am.MinApprovers()
			possible_approvers, _ := am.PossibleApprovers()
			if min_approvers > 0 && len(possible_approvers) == 0 {
				return nil, errors.New("mission requires approvers but no possible approvers set")
			}
			return next.Mutate(ctx, am)
		})
	}, ent.OpCreate|ent.OpUpdate))
}
