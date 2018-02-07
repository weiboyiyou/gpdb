package services

import (
	"gp_upgrade/utils"
	"path"

	"os"

	"github.com/pkg/errors"
)

type ClusterSsher struct {
	checklistWriter ChecklistWriter
}

type ChecklistWriter interface {
	MarkInProgress(string) error
	ResetStateDir(string) error
}

type ChecklistWriterImpl struct {
	pathToStateDir string
}

func (c *ChecklistWriterImpl) MarkInProgress(step string) error {
	_, err := utils.System.OpenFile(path.Join(c.pathToStateDir, "in.progress"), os.O_RDONLY|os.O_CREATE, 0700)
	if err != nil {
		return err
	}

	return nil
}

func (c *ChecklistWriterImpl) ResetStateDir(step string) error {
	homeDirectory := utils.System.Getenv("HOME")
	if homeDirectory == "" {
		return errors.New("Could not find the home directory environment variable")

	}
	pathToStateDir := path.Join(homeDirectory, ".gp_upgrade", step)
	err := utils.System.RemoveAll(pathToStateDir)
	if err != nil {
		return err
	}
	err = utils.System.MkdirAll(pathToStateDir, 0700)
	if err != nil {
		return err
	}
	c.pathToStateDir = pathToStateDir
	return nil

}

func NewChecklistWriterImpl() *ChecklistWriterImpl {
	return &ChecklistWriterImpl{}
}

func NewClusterSsher(cw ChecklistWriter) *ClusterSsher {
	return &ClusterSsher{checklistWriter: cw}
}

func (c *ClusterSsher) VerifySoftware([]string) {
	c.checklistWriter.ResetStateDir("seginstall")
	c.checklistWriter.MarkInProgress("seginstall")
}
