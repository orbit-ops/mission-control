package controller

import (
	"testing"

	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent"
	"github.com/orbit-ops/launchpad-core/ent/ogent"
	"github.com/orbit-ops/launchpad-core/utils"
)

func (ctrl *testAccessController) getValidApproval() *ogent.CreateApprovalReq {
	req := ctrl.createValidRequest()

	return &ogent.CreateApprovalReq{
		Person:   "test-" + utils.RandomSeq(6),
		Approved: true,
		Request:  req.ID,
	}
}

func (ctrl *testAccessController) createValidApproval() *ogent.ApprovalCreate {
	resp, err := ctrl.Access.CreateApproval(ctrl.Context, ctrl.getValidApproval())
	ctrl.ok(err, "creating base Approval")

	return resp.(*ogent.ApprovalCreate)
}

func TestApprovalCreate(t *testing.T) {
	ctrl := newTestController(t)
	ctrl.createValidApproval()
}

func TestApprovalWithNonExistingRequestThrowError(t *testing.T) {
	ctrl := newTestController(t)

	_, err := ctrl.Access.CreateApproval(ctrl.Context, &ogent.CreateApprovalReq{
		Person:   "test-" + utils.RandomSeq(6),
		Approved: true,
		Request:  uuid.New(),
	})

	ctrl.assert(err == nil, "non-existing request did not throw error")
}

func TestApprovalCorrectApprovers(t *testing.T) {
	tests := []struct {
		name    string
		request *ent.Request
		wantErr bool
	}{
		{
			name:    "doubleApprovers",
			wantErr: true,
		},
	}
}

func TestApprovalDoubleApproveThrowError(t *testing.T) {
	ctrl := newTestController(t)
	appr := ctrl.getValidApproval()

	_, _ = ctrl.Access.CreateApproval(ctrl.Context, appr)
	_, err := ctrl.Access.CreateApproval(ctrl.Context, appr)
	ctrl.assert(err == nil, "double approval did not throw error")
}

func TestApprovalAllDoneCreatesAccess(t *testing.T) {
	ctrl := newTestController(t)
	appr := ctrl.getValidApproval()

	_, err := ctrl.Access.CreateApproval(ctrl.Context, appr)
	ctrl.ok(err)

	_, err = ctrl.Access.CreateApproval(ctrl.Context, appr)
	ctrl.assert(err == nil, "double approval did not throw error")
}

func TestApprovalWithoutAllDoneNoAccess(t *testing.T) {
	ctrl := newTestController(t)
	mission := ctrl.getValidMission()
	mission.MinApprovers = 2
	mission.PossibleApprovers = []string{"test1", "test2"}
	m, err := ctrl.Access.CreateMission(ctrl.Context, mission)
	ctrl.ok(err, "creating misson")
	mc, _ := m.(*ogent.MissionCreate)

	req := ctrl.createValidRequestWithMission(mc.ID)
	appr := &ogent.CreateApprovalReq{
		Person:   "test1",
		Approved: true,
		Request:  req.ID,
	}

	_, err = ctrl.Access.CreateApproval(ctrl.Context, appr)
	ctrl.ok(err, "creating first approval")
	_, err = ctrl.Access.CreateApproval(ctrl.Context, appr)
	ctrl.ok(err, "creating second approval")
}
