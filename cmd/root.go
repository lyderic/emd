package cmd

import (
	"fmt"
	"os"

	"golang.org/x/term"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:                   "emd [-n] [-t <theme>] [-w <width>] file.md",
	DisableFlagsInUseLine: true,
	SilenceUsage:          true,
	Short:                 "A basic markdown viewer for the command line",
	Version:               VERSION,
	Args:                  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return render(args[0])
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile string

func init() {
	w, _, err := term.GetSize(0)
	if err != nil {
		color.Yellow("cannot get terminal size")
	}
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.emd.yaml)")
	rootCmd.PersistentFlags().StringP("theme", "t", "dark", "theme")
	rootCmd.PersistentFlags().IntP("width", "w", w, "word wrap width")
	rootCmd.PersistentFlags().BoolP("no-pager", "n", false, "don't use pager")
	rootCmd.PersistentFlags().SortFlags = false
	viper.BindPFlag("theme", rootCmd.PersistentFlags().Lookup("theme"))
	viper.BindPFlag("width", rootCmd.PersistentFlags().Lookup("width"))
	viper.BindPFlag("no-pager", rootCmd.PersistentFlags().Lookup("no-pager"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".emd")
	}
	viper.SetEnvPrefix("EMD")
	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
