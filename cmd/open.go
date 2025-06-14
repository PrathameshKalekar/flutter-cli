package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open current directory in VS Code",
	Long: `Opens the current working directory in Visual Studio Code.

Example:
  pralex open

This will launch the current directory in VS Code using the 'code' command.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := openDirectoryInVSCode(); err != nil {
			fmt.Println("❌ Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func openDirectoryInVSCode() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current working directory: %w", err)
	}

	cmd := exec.Command("code", cwd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to open VS Code: %w", err)
	}

	fmt.Println("✅ Opened current directory in VS Code")
	return nil
}
