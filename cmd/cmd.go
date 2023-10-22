package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type globalFlags struct {
	directory string
	output string
}

var flag *globalFlags

var rootCmd = &cobra.Command{
	Use: "maphash",
	Short: "A simple CLI program to get all files in that directory IEEE CRC32 Hashes.",
	RunE: func (cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	flag = &globalFlags{}

	rootCmd.PersistentFlags().StringVar(
		&flag.directory, 
		"dir", 
		"./", 
		"The directory to read from.",
	)

	rootCmd.PersistentFlags().StringVar(
		&flag.output, 
		"output", 
		"./maps.json", 
		"The file to put all the hashes in.",
	)
}

func Execute() error {
	start := time.Now()

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	done := time.Now()
	fmt.Println(done.Sub(start))

	return nil
}