/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"mydocker/internal"
	"mydocker/pkg/container"

	"github.com/spf13/cobra"
)

// forkCmd represents the fork command

// NewForkCommand implements and returns fork command.
// fork command is called by reexec to apply namespaces.
//
// It is a hidden command and requires root path and
// container id to run.

func init() {
	ctr := container.NewContainer()
	var detach bool

	var forkCmd = &cobra.Command{

		Use:          "fork",
		Short:        " It is a hidden command and requires root path and container id to run",
		Hidden:       true,
		SilenceUsage: true,
		PreRunE:      isRoot,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ctr.LoadConfig(); err != nil {
				return err
			}
			return internal.Fork(ctr, args, detach)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("fork called")
		},
	}

	flags := forkCmd.Flags()
	flags.StringVar(&ctr.Digest, "container", "", "")
	flags.StringVar(&ctr.RootFS, "root", "", "")
	flags.StringVar(&ctr.Config.Hostname, "host", "", "")
	flags.BoolVar(&detach, "detach", false, "")
	mem := flags.Int("memory", 100, "")
	swap := flags.Int("swap", 20, "")
	cpu := flags.Float64("cpus", 1, "")
	pids := flags.Int("pids", 128, "")
	ctr.SetMemorySwapLimit(*mem, *swap)
	ctr.SetCPULimit(*cpu)
	ctr.SetProcessLimit(*pids)

	forkCmd.MarkFlagRequired("root")
	forkCmd.MarkFlagRequired("container")

	rootCmd.AddCommand(forkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
