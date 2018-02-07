package upgradestatus

import (
	pb "gp_upgrade/idl"
)

type Seginstall struct {
	seginstallPath string
}

func NewSeginstall(seginstallPath string) Seginstall {
	return Seginstall{seginstallPath: seginstallPath}
}

func (c Seginstall) GetStatus() (*pb.UpgradeStepStatus, error) {
	return &pb.UpgradeStepStatus{Step: pb.UpgradeSteps_SEGINSTALL, Status: pb.StepStatus_PENDING}, nil
}
