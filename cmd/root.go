/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

const (
	layersPath     = "/var/lib/vessel/images/layers"
	containersPath = "/var/run/vessel/containers"
	netnsPath      = "/var/run/vessel/netns"
)

var ErrNotPermitted = errors.New("operation not permitted")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                   "vessel [OPTIONS] COMMAND",
	Short:                 "A tiny tool for managing containers",
	Long:                  `A tool for managing containers, including ps, run, images`,
	TraverseChildren:      true,
	DisableFlagsInUseLine: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Make vessel directories first.
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mydocker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	os.MkdirAll(netnsPath, 0700)
	os.MkdirAll(layersPath, 0700)
	os.MkdirAll(containersPath, 0700)
}

// isRoot implements a cobra acceptable function and
// returns ErrNotPermitted if user is not root.
func isRoot(_ *cobra.Command, _ []string) error {
	if os.Getuid() != 0 {
		return ErrNotPermitted
	}
	return nil
}
