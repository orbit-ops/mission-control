package missions

import (
	"fmt"
	"os/exec"

	"github.com/orbit-ops/launchpad-core/ent"
)

type Mission struct {
	Name string
	Path string
}

type MissionController struct {
	missions map[string]*Mission
}

func NewMissionController() *MissionController {
	return &MissionController{
		missions: make(map[string]*Mission),
	}
}

func (pc *MissionController) CreateAccess(missionID string, msg *ent.Access) error {
	cmd := exec.Command(pc.missions[missionID].Path)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("exec create access: %w", err)
	}

	return nil
}

func (pc *MissionController) RemoveAccess(missionID string, msg *ent.Access) error {
	cmd := exec.Command(pc.missions[missionID].Path)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("exec remove access: %w", err)
	}

	return nil
}
