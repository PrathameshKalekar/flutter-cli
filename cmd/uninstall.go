package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Clean the Flutter project",
	Long: `Run 'flutter clean' to remove the build directory and cached data 
from your Flutter project.`,
	Run: func(cmd *cobra.Command, args []string) {
		uninstall()
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

func uninstall() {
	fmt.Println("🧹 Cleaning up Flutter project...")

	cmd := exec.Command("flutter", "clean")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to clean project: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Project cleaned successfully.")
}
