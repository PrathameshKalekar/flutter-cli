package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set <flutter_version>",
	Short: "Switch Flutter version using FVM",
	Long: `Switch your Flutter version using FVM (Flutter Version Manager).

Example:
  pralex set 3.27.0`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: pralex set <flutter_version>")
			return
		}
		switchFlutterVersion(args[0])
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

func switchFlutterVersion(version string) {
	fmt.Printf("🔁 Switching Flutter to version %s...\n", version)

	cmd := exec.Command("fvm", "use", version)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Failed to switch Flutter version: %v\n", err)
		return
	}

	fmt.Println("✅ Flutter version switched successfully")
}
