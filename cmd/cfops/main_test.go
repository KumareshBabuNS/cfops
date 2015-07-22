package main

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewApp", func() {
	Describe("`cfops backup` command", func() {
		runTestSuiteFor("backup")
	})

	Describe("`cfops restore` caommand", func() {
		runTestSuiteFor("restore")
	})
})

func runTestSuiteFor(command string) {
	var (
		app                 = NewApp()
		requiredArgs        []string
		allArgs             []string
		invalidArgs         []string
		missingRequiredArgs []string
	)

	BeforeEach(func() {
		ExitCode = cleanExitCode
		app = NewApp()
		requiredArgs = []string{
			"cfops",
			command,
			"--opsmanagerhost", "<host>",
			"--adminuser", "<usr>",
			"--adminpass", "<pass>",
			"--opsmanageruser", "<opsuser>",
			"--opsmanagerpass", "<opspass>",
			"-d", "<dir>",
		}
		allArgs = append(requiredArgs, "-tl", "'opsmanager, er'")
		invalidArgs = append(requiredArgs, "--fakearg", "blah")
		missingRequiredArgs = []string{
			"cfops",
			command,
			"--opsmanagerhost", "<host>",
			"--adminuser", "<usr>",
			"--opsmanagerpass", "<opspass>",
			"-d", "<dir>",
		}
	})

	Context("When given all required arguments", func() {
		It("Should not throw an error", func() {
			err := app.Run(requiredArgs)
			Ω(err).Should(BeNil())
		})
	})

	Context("When given all available arguments", func() {
		It("Should not throw an error", func() {
			err := app.Run(missingRequiredArgs)
			Ω(err).Should(BeNil())
		})
	})

	Context("When given invalid arguments", func() {
		It("Should throw an error", func() {
			fmt.Println(invalidArgs)
			err := app.Run(invalidArgs)
			Ω(err).ShouldNot(BeNil())
		})
	})

	Context("When missing a required argument", func() {
		It("Should throw an error", func() {
			fmt.Println(missingRequiredArgs)
			app.Run(missingRequiredArgs)
			Ω(ExitCode).Should(Equal(helpExitCode))
		})
	})
}
