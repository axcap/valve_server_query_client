package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "vmon",
	Short: "Query Valve servers",
}

var Verbose bool

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	if Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output")
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
