package controller

import (
	"testing"

	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/ogent"
	"github.com/orbit-ops/launchpad-core/utils"
)

func (ctrl *testAccessController) getValidRequest() *ogent.CreateRequestReq {
	mission := ctrl.createValidMission()

	return ctrl.getValidRequestWithMission(mission.ID)
}

func (ctrl *testAccessController) getValidRequestWithMission(mid uuid.UUID) *ogent.CreateRequestReq {
	return &ogent.CreateRequestReq{
		Requester: "test-" + utils.RandomSeq(6),
		Reason:    "test",
		Mission:   mid,
	}
}

func (ctrl *testAccessController) createValidRequest() *ogent.RequestCreate {
	resp, err := ctrl.Access.CreateRequest(ctrl.Context, ctrl.getValidRequest())
	if err != nil {
		ctrl.t.Fatalf("creating base Request: %v", err)
	}
	return resp.(*ogent.RequestCreate)
}

func (ctrl *testAccessController) createValidRequestWithMission(mid uuid.UUID) *ogent.RequestCreate {
	resp, err := ctrl.Access.CreateRequest(ctrl.Context, ctrl.getValidRequestWithMission(mid))
	if err != nil {
		ctrl.t.Fatalf("creating base Request: %v", err)
	}
	return resp.(*ogent.RequestCreate)
}

func TestRequestCreate(t *testing.T) {
	ctrl := newTestController(t)
	ctrl.createValidRequest()
}

func TestRequestWithNonExistingMissionThrowsError(t *testing.T) {
	ctrl := newTestController(t)
	_, err := ctrl.Access.CreateRequest(ctrl.Context, &ogent.CreateRequestReq{
		Requester: "test-" + utils.RandomSeq(6),
		Reason:    "test",
		Mission:   uuid.New(),
	})
	if err == nil {
		t.Fatal("non-existing mission did not throw error")
	}
}

func TestRequestWithEmptyReasonThrowsError(t *testing.T) {
	ctrl := newTestController(t)
	req := ctrl.getValidRequest()
	req.Reason = ""
	_, err := ctrl.Access.CreateRequest(ctrl.Context, req)
	if err == nil {
		t.Fatal("empty reason did not throw error")
	}
}

func TestRequestWithEmptyRequestThrowsError(t *testing.T) {
	ctrl := newTestController(t)
	req := ctrl.getValidRequest()
	req.Requester = ""
	_, err := ctrl.Access.CreateRequest(ctrl.Context, req)
	if err == nil {
		t.Fatal("empty reason did not throw error")
	}
}
