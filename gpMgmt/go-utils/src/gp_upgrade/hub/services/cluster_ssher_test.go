package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gp_upgrade/hub/services"
)

var _ = Describe("ClusterSsher", func() {
	It("indicates that it is in progress on the hub filesystem", func() {
		cw := newSpyChecklistWriter()
		clusterSsher := services.NewClusterSsher(cw)
		clusterSsher.VerifySoftware([]string{"doesn't matter"})
		Expect(cw.freshStateDirs).To(ContainElement("seginstall"))
		Expect(cw.stepsMarkedInProgress).To(ContainElement("seginstall"))
	})
})

func newSpyChecklistWriter() *spyChecklistWriter {
	return &spyChecklistWriter{}
}

type spyChecklistWriter struct {
	freshStateDirs        []string
	stepsMarkedInProgress []string
}

func (s *spyChecklistWriter) MarkInProgress(step string) error {
	s.stepsMarkedInProgress = append(s.stepsMarkedInProgress, step)
	return nil
}

func (s *spyChecklistWriter) ResetStateDir(step string) error {
	s.freshStateDirs = append(s.freshStateDirs, step)
	return nil
}
