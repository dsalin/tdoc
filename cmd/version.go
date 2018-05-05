package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "tdoc: terminal based tool for documentation",
	Long:  `Create your personal documentation in terminal. Edit, search, tag whatever you want.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tdoc: version 0.1 (dev)")
	},
}
