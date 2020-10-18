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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Exec versionCmd")
	},
}
