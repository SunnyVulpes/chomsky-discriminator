/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"chomsky-discriminator/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chomsky-discriminator",
	Short: "Classify grammars into Chomsky hierarchy categories based on non-terminal symbols and grammar rules.",
	Long: `The chomsky-discriminator is a command-line tool designed to analyze and classify grammars into the Chomsky hierarchy, including Type 0 (Recursively Enumerable), Type 1 (Context-Sensitive), Type 2 (Context-Free), and Type 3 (Regular) grammars. Users can input non-terminal symbols and a set of grammar rules, and the tool will determine the most restrictive category that the grammar fits into.

This tool is ideal for computational linguistics, language design, and formal language theory studies. For example:

Using the tool, you might input the following grammar rules:
  S -> ABC | a
  A -> a | ε
  B -> b
  C -> c

And receive an analysis indicating that the grammar is Context-Free.

The chomsky-discriminator leverages the Cobra CLI library to provide a robust, user-friendly interface. Its functionality aids in educational and practical applications where understanding the limitations and capabilities of different grammar types is crucial.`,
	Run: discriminate,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func discriminate(cmd *cobra.Command, args []string) {
	g := pkg.BuildGrammar()
	fmt.Printf("%v", *g)
}
