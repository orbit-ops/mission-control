package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/access"
	"github.com/orbit-ops/launchpad-core/ent/actiontokens"
	"github.com/orbit-ops/launchpad-core/ent/approval"
	"github.com/orbit-ops/launchpad-core/ent/audit"
	"github.com/orbit-ops/launchpad-core/ent/request"
	"github.com/orbit-ops/launchpad-core/internal/notifications"
	"github.com/orbit-ops/launchpad-core/providers"
	"github.com/orbit-ops/launchpad-core/utils"
)

type AccessController struct {
	client *ent.Client
	// mc       *missions.MissionController
	notifier notifications.Notifier
	provider providers.Provider
}

func NewAccessController(prov providers.Provider, client *ent.Client) (*AccessController, error) {
	return &AccessController{
		client:   client,
		provider: prov,
		// mc:     missions.NewMissionController(),
	}, nil
}

func (c *AccessController) Shutdown() error {
	if err := c.client.Close(); err != nil {
		return fmt.Errorf("closing ent: %w", err)
	}
	return nil
}

func (c *AccessController) CreateApproval(ctx context.Context, req *ent.Request, approver string) error {
	approved, err := c.client.Approval.Query().Where(approval.RequestIDEQ(req.ID)).Count(ctx)
	if err != nil {
		return fmt.Errorf("checking approvals for request: %w", err)
	}

	r, err := c.client.Request.Query().Where(request.IDEQ(req.ID)).WithMission().First(ctx)
	if err != nil {
		return fmt.Errorf("retrieving request for approval: %w", err)
	}

	if _, err = c.client.Approval.Create().
		SetApproved(true).
		SetApprovedTime(time.Now()).
		SetPerson(approver).
		Save(ctx); err != nil {

		return fmt.Errorf("creating approval: %w", err)
	}

	if _, err := c.client.Audit.Create().
		SetAction(audit.ActionApproveRequest).
		SetAuthor(approver).
		SetTimestamp(time.Now()).Save(ctx); err != nil {
		return fmt.Errorf("creating audit entry for approval: %w", err)
	}

	if approved+1 < r.Edges.Mission.MinApprovers {
		return nil
	}

	return c.CreateAccess(ctx, r.Edges.Mission, req)
}

func (c *AccessController) RemoveApproval(ctx context.Context, req *ent.Request, approver string) error {
	exists, err := c.client.Access.Query().Where(access.RequestIDEQ(req.ID)).Exist(ctx)
	if err != nil {
		return fmt.Errorf("checking access already provisioned: %w", err)
	}
	if exists {
		return fmt.Errorf("cannot revoke approval, access provisioned")
	}

	if _, err := c.client.Audit.Create().
		SetAction(audit.ActionRevokeApprovalRequest).
		SetAuthor(approver).
		SetTimestamp(time.Now()).Save(ctx); err != nil {
		return fmt.Errorf("creating audit entry for approval: %w", err)
	}

	_, err = c.client.Approval.Delete().Where(approval.PersonEQ(approver), approval.ApprovedEQ(true), approval.RequestIDEQ(req.ID)).Exec(ctx)
	if err != nil {
		return fmt.Errorf("executing approval removal: %w", err)
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
