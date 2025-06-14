package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set <flutter_version>",
	Short: "Switches Flutter version using FVM",
	Long:  `Switches your Flutter version using the FVM (Flutter Version Manager).`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: pralex set <flutter_version>")
			return
		}
		version := args[0]
		switchFlutterVersion(version)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

// switchFlutterVersion runs the `fvm use <version>` command
func switchFlutterVersion(version string) {
	fmt.Printf("Switching Flutter to version %s...\n", version)

	output, err := exec.Command("fvm", "use", version).CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Error: %v\nOutput: %s\n", err, string(output))
		return
	}

	fmt.Printf("✅ Success: %s\n", string(output))
}
