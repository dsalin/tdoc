package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "tdoc",
	Short: "tdoc: terminal based tool for documentation",
	Long:  `Create your personal documentation in terminal. Edit, search, tag whatever you want.`,
}
