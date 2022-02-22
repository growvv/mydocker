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
	// runCmd represents the run command
	var runCmd = &cobra.Command{
		Use:                   "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
		Short:                 "Run a command inside a new Container.",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Args:                  cobra.MinimumNArgs(1),
		PreRunE:               isRoot,
		RunE:                  internal.Run,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run called")
		},
	}

	flags := runCmd.Flags()
	flags.StringP("host", "", "", "Container Hostname")
	flags.IntP("memory", "m", 100, "Limit memory access in MB")
	flags.IntP("swap", "s", 20, "Limit swap access in MB")
	flags.Float64P("cpus", "c", 2, "Limit CPUs")
	flags.IntP("pids", "p", 128, "Limit number of processes")
	flags.BoolP("detach", "d", false, "run command in the background")

	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
