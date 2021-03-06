package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"os/exec"
	"path/filepath"
)

var _ = Describe("Main", func() {
	var pathToFactorialsBinary string

	BeforeSuite(func() {
		var err error
		pathToFactorialsBinary, err = gexec.Build("github.com/benlaplanche/theregister-bluemix-challenge/factorials")

		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("Reading from ARGS", func() {

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

	})

	Describe("Reading from a provided file path", func() {
		Context("That has valid input", func() {
			It("returns the correct values", func() {
				input, _ := filepath.Abs("assets/valid_input.txt")

				command := exec.Command(pathToFactorialsBinary, input)
				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

				Expect(err).ShouldNot(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("1"))
				Eventually(session.Out).Should(gbytes.Say("24"))
				Eventually(session.Out).Should(gbytes.Say("6"))
			})
		})

		Context("That has valid input with additional values after the seperator", func() {
			It("Should only return values before seperator", func() {
				input, _ := filepath.Abs("assets/additional_input.txt")

				command := exec.Command(pathToFactorialsBinary, input)
				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

				Expect(err).ShouldNot(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("1"))
				Eventually(session.Out).Should(gbytes.Say("24"))
				Eventually(session.Out).Should(gbytes.Say("6"))
				Eventually(session.Out).ShouldNot(gbytes.Say("120"))
			})
		})

		Context("That has values greater than 15", func() {
			It("Should ignore values greater than 15", func() {
				input, _ := filepath.Abs("assets/sixteen_input.txt")

				command := exec.Command(pathToFactorialsBinary, input)
				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

				Expect(err).ShouldNot(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("1307674368000"))
				Eventually(session.Out).ShouldNot(gbytes.Say("20922789888000"))
			})
		})

		Context("That has invalid input", func() {
			It("Should strip the invalid characters and return correct results", func() {
				input, _ := filepath.Abs("assets/invalid_input.txt")

				command := exec.Command(pathToFactorialsBinary, input)
				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

				Expect(err).ShouldNot(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("1"))
				Eventually(session.Out).Should(gbytes.Say("24"))
				Eventually(session.Out).Should(gbytes.Say("6"))
			})
		})
	})

	Describe("Error scenarios - ", func() {
		Context("No args provided", func() {
			It("Provides an error message", func() {
				command := exec.Command(pathToFactorialsBinary)
				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

				Expect(err).ShouldNot(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("Please provide either a file path or a number input"))
			})
		})
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

})
