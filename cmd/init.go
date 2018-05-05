package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const permissionBits os.FileMode = 0755

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize current directory",
	Run: func(cmd *cobra.Command, args []string) {
		os.MkdirAll("./.tdoc", permissionBits)

		// help directories
		os.MkdirAll("./.tdoc/refs", permissionBits)
		os.MkdirAll("./.tdoc/objects", permissionBits)
		os.MkdirAll("./.tdoc/hooks", permissionBits)

		// help files
		createFileOrPanic("./.tdoc/HEAD")
		createFileOrPanic("./.tdoc/index")
		createFileOrPanic("./.tdoc/config")

		fmt.Println("tdoc: directory initialized")
	},
}

func createFileOrPanic(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
