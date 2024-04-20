package cmd

import (
	"chomsky-discriminator/pkg"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test a case",
	Long:  ``,
	Run:   testCase,
}

func testCase(cmd *cobra.Command, args []string) {
	g := pkg.TestCase()

	g.Print()
}
