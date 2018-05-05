package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the current status of the directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tdoc: status")
	},
}
