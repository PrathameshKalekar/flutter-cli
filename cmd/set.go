/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: pralex set <flutter_version>")
			return
		}
		version := args[0]
		fmt.Printf("Switching Flutter to version %s...\n", version)

		// Execute the fvm command
		out, err := exec.Command("fvm", "use", version).CombinedOutput()
		if err != nil {
			fmt.Printf("Error: %v\nOutput: %s\n", err, string(out))
			return
		}
		fmt.Printf("Success: %s\n", string(out))
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
