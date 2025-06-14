package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var showVersion bool

var rootCmd = &cobra.Command{
	Use:   "pralex",
	Short: "A CLI tool for Flutter utilities",
	Long:  "Pralex is a CLI tool that helps with Flutter development tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			out, err := exec.Command("flutter", "--version").CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error running flutter --version: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(out))
		} else {
			_ = cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add the -v flag
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show Flutter version (flutter --version)")
}
