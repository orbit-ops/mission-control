package controller

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/hook"
	"github.com/orbit-ops/launchpad-core/ent/rocket"
)

func (ac *AccessController) setMissionHooks() {
	ac.client.Mission.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.MissionFunc(func(ctx context.Context, am *ent.MissionMutation) (ent.Value, error) {
			min_approvers, _ := am.MinApprovers()
			possible_approvers, _ := am.PossibleApprovers()

			// no duplicate people in approvers
			approvers := make(map[string]bool)
			for _, a := range possible_approvers {
				if _, ok := approvers[a]; ok {
					return nil, errors.New("approvers must be unique")
				} else {
					approvers[a] = true
				}
			}

			// possible approvers must be at least as big as minimum approvers
			if min_approvers > len(possible_approvers) {
				return nil, errors.New("amount of possible approvers must not be smaller than minimum approvers")
			}

			// rockets must exist
			rockets := am.RocketsIDs()
			existing, err := am.Client().Rocket.Query().Where(rocket.IDIn(rockets...)).All(ctx)
			if err != nil {
				return nil, err
			}

			if len(existing) < len(rockets) {
				dont_exist := make([]string, 0)
				for _, r := range rockets {

					if !slices.ContainsFunc(existing, func(er *ent.Rocket) bool {
						return er.ID == r
					}) {
						dont_exist = append(dont_exist, r.String())
					}
				}

				return nil, fmt.Errorf("mission has non existing rockets: %s", strings.Join(dont_exist, ", "))
			}

			return next.Mutate(ctx, am)
		})
	}, ent.OpCreate|ent.OpUpdate))
}
