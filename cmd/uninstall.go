/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		uninstall()
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

func uninstall() {
	var cmd *exec.Cmd
	fmt.Println("🧹 Cleaning up Flutter project...")
	cmd = exec.Command("flutter", "clean")

	cmd.Stdout = cmd.Stdout
	cmd.Stderr = cmd.Stderr
	cmd.Stdin = cmd.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Command failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Done!")
}
