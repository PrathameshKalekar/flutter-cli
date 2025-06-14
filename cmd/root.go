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
	Long:  `Pralex is a lightweight CLI tool that simplifies common Flutter development tasks like version switching, branch handling, and Git operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			output, err := exec.Command("flutter", "--version").CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "❌ Error running 'flutter --version': %v\n", err)
				os.Exit(1)
			}
			fmt.Print(string(output))
			return
		}
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show Flutter version using 'flutter --version'")
}
