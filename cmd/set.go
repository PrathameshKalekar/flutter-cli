package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set <flutter_version>",
	Short: "Switch Flutter version using FVM",
	Long: `Use this command to switch your Flutter version via FVM (Flutter Version Manager).

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

	output, err := exec.Command("fvm", "use", version).CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Failed to switch Flutter version:\n%v\n%s\n", err, string(output))
		return
	}

	fmt.Printf("✅ Flutter version switched successfully:\n%s\n", string(output))
}
