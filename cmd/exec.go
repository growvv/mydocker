/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"mydocker/internal"

	"github.com/spf13/cobra"
)

func init() {

	// execCmd represents the exec command
	var detach bool
	var execCmd = &cobra.Command{
		Use:                   "exec [OPTIONS] CONTAINER COMMAND [ARG...]",
		Short:                 "Run a command inside a existing Container.",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Args:                  cobra.MinimumNArgs(2),
		PreRunE:               isRoot,
		RunE: func(cmd *cobra.Command, args []string) error {
			return internal.Exec(args[0], args[1:], detach)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("exec called")
		},
	}

	flags := execCmd.Flags()
	flags.BoolVarP(&detach, "detach", "d", false, "run command in the background")

	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
