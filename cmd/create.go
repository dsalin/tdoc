package cmd

import (
	"fmt"

	"strings"

	"github.com/spf13/cobra"
)

func init() {
	createCmd.AddCommand(createTagCmd)
	createCmd.AddCommand(createTagsCmd)
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create tdoc entity: document, tag, reference.",
}

var createTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create tdoc tag",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tdoc: create tag: %s\n", args[0])
	},
}

var createTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Create tdoc tags",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tdoc: create tags \n\n%s\n", strings.Join(args, "\n"))
	},
}
