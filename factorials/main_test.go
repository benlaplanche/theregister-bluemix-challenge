package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"os/exec"
)

var _ = Describe("Main", func() {

	Describe("Reading from ARGS", func() {

		var pathToFactorialsBinary string

		BeforeSuite(func() {
			var err error
			pathToFactorialsBinary, err = gexec.Build("github.com/benlaplanche/theregister-bluemix-challenge/factorials")

			Expect(err).ShouldNot(HaveOccurred())
		})

		Context("With valid input params", func() {

			It("should return 1 when the input is 1", func() {
				input := "1"
				command := exec.Command(pathToFactorialsBinary, input)
				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

				Expect(err).ShouldNot(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("1"))
			})

			It("should return the correct answer when greater than 1", func() {
				input := "4"
				command := exec.Command(pathToFactorialsBinary, input)
				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

				Expect(err).ShouldNot(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("24"))
			})
		})

		AfterSuite(func() {
			gexec.CleanupBuildArtifacts()
		})

	})

})
