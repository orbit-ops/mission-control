package controller

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/actiontokens"
	"github.com/orbit-ops/launchpad-core/ent/audit"
	"github.com/orbit-ops/launchpad-core/ent/ogent"
	"github.com/orbit-ops/launchpad-core/ent/request"
	"github.com/orbit-ops/launchpad-core/internal/notifications"
	"github.com/orbit-ops/launchpad-core/providers"
	"github.com/orbit-ops/launchpad-core/utils"
)

type AccessController struct {
	*ogent.OgentHandler

	client *ent.Client
	// mc       *missions.MissionController
	notifier notifications.Notifier
	provider providers.Provider
}

func NewAccessController(prov providers.Provider, client *ent.Client) (*AccessController, error) {
	ac := &AccessController{
		OgentHandler: ogent.NewOgentHandler(client),
		client:       client,
		provider:     prov,
		// mc:     missions.NewMissionController(),
	}

	// logs and audits all mutation operations of all types
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (val ent.Value, err error) {
			if _, ok := m.(*ent.AuditMutation); ok {
				return next.Mutate(ctx, m)
			}
			start := time.Now()

			defer func() {
				log.Debugf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\tError=%v\n", m.Op(), m.Type(), time.Since(start), m, err)
			}()

			val, err = next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			user := ctx.Value(utils.ContextUserKey{}).(string)

			if _, err := client.Audit.Create().
				SetAction(audit.ActionApproveRequest).
				SetAuthor(user).
				SetTimestamp(start).Save(ctx); err != nil {
				return nil, fmt.Errorf("creating audit entry for approval: %w", err)
			}

			return
		})
	})

	ac.setAccessHooks()
	ac.setRequestHooks()
	ac.setApprovalHooks()
	ac.setAuditHooks()
	ac.setMissionHooks()

	return ac, nil
}

func (c *AccessController) Shutdown() error {
	if err := c.client.Close(); err != nil {
		return fmt.Errorf("closing ent: %w", err)
	}
	return nil
}

func (c *AccessController) CancelAccess(ctx context.Context, acc *ent.Access, author string) error {
	if err := acc.Update().SetRolledBack(true).SetRollbackTime(time.Now()).Exec(ctx); err != nil {
		return err
	}

	if err := c.removeAccess(ctx, acc); err != nil {
		return fmt.Errorf("removing access: %w", err)
	}

	if _, err := c.client.Audit.Create().
		SetAction(audit.ActionRevokeApprovalRequest).
		SetAuthor(author).
		SetTimestamp(time.Now()).Save(ctx); err != nil {
		return fmt.Errorf("creating audit entry for approval: %w", err)
	}

	return nil
}

func (c *AccessController) FinishAccess(ctx context.Context, acc *ent.Access) error {
	if err := acc.Update().SetRolledBack(true).SetRollbackTime(time.Now()).Exec(ctx); err != nil {
		return err
	}

	if err := c.removeAccess(ctx, acc); err != nil {
		return fmt.Errorf("removing access: %w", err)
	}

	if _, err := c.client.Audit.Create().
		SetAction(audit.ActionRevokeApprovalRequest).
		SetAuthor("engine").
		SetTimestamp(time.Now()).Save(ctx); err != nil {
		return fmt.Errorf("creating audit entry for approval: %w", err)
	}
	return nil
}

func (c *AccessController) CreateAccess(ctx context.Context, mission *ent.Mission, req *ent.Request) error {
	acc, err := c.client.Access.Create().Save(ctx)
	if err != nil {
		return fmt.Errorf("creating access db entry for %#v: %w", req, err)
	}

	token, err := c.client.ActionTokens.Create().
		SetAction(actiontokens.ActionCreate).
		SetExpiration(time.Now().Add(5 * time.Minute)).
		SetToken(utils.RandomSeq(32)).
		SetAccessID(acc.ID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating token for provisioning %#v: %w", req, err)
	}

	for _, rocket := range mission.Edges.Rockets {
		if err := c.provider.CreateAccess(ctx, token.Token, rocket, req); err != nil {
			return fmt.Errorf("provisioning access: %w", err)
		}
	}

	if err := c.notifier.Notify(""); err != nil {
		return fmt.Errorf("failing to notify %s of access created: %w", req.Requester, err)
	}

	return c.scheduleDeletion(ctx, acc)
}

func (c *AccessController) removeAccess(ctx context.Context, acc *ent.Access) error {
	if err := acc.Update().SetRolledBack(true).Exec(ctx); err != nil {
		return err
	}

	// mission, err := c.client.Mission.Query().Select().Where(mission.HasRequestsWith(request.IDEQ(acc.RequestID))).First(ctx)
	req, err := c.client.Request.Query().Where(request.IDEQ(acc.RequestID)).WithMission(func(mq *ent.MissionQuery) {
		mq.WithRockets()
	}).First(ctx)
	if err != nil {
		return err
	}

	token, err := c.client.ActionTokens.Create().
		SetAction(actiontokens.ActionCreate).
		SetExpiration(time.Now().Add(5 * time.Minute)).
		SetToken(utils.RandomSeq(32)).
		SetAccessID(acc.ID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating token for provisioning %#v: %w", req, err)
	}

	for _, rocket := range req.Edges.Mission.Edges.Rockets {
		if err := c.provider.RemoveAccess(ctx, token.Token, rocket, req); err != nil {
			return fmt.Errorf("provisioning access: %w", err)
		}
	}

	if err := c.notifier.Notify(""); err != nil {
		return fmt.Errorf("failing to notify %s of access removed: %w", req.Requester, err)
	}

	return nil
}

func (c *AccessController) scheduleDeletion(ctx context.Context, acc *ent.Access) error {
	return nil
}
