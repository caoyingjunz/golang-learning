package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A short versionCmd demo",
	Long:  `A Long versionCmd demo`,
	// PersistentPreRun(global)： 子命令会继承 rootCmd
	// PreRun(local): 子命令不会继承 rootCmd
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Run versionCmd PersistentPreRun with args: %v\n", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("run versionCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Exec versionCmd")
	},
}
