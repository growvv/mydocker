/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"mydocker/internal"

	"github.com/spf13/cobra"
)

// imagesCmd represents the images command
var imagesCmd = &cobra.Command{
	Use:                   "images",
	Short:                 "List local images",
	DisableFlagsInUseLine: true,
	SilenceUsage:          true,
	Args:                  cobra.NoArgs,
	PreRunE:               isRoot,
	RunE:                  internal.Images,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("images called")
	},
}

func init() {
	rootCmd.AddCommand(imagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
