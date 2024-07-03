package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var rootCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	//SilenceUsage:          true,
	Use:     "emd [-n] [-t <theme>] [-w <width>] file.md",
	Short:   DESCRIPTION, // globals.go
	Long:    longHelp(),  // utils.go
	Version: VERSION,
	Args:    cobra.ExactArgs(1),
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
	cobra.OnInitialize(initConfig)
	w := initTerm()
	debug(fmt.Sprintf("w=%d", w))
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "",
		"config `file` (default is $HOME/.emd.yaml)")
	rootCmd.PersistentFlags().StringP("theme", "t", "dark",
		"`name` of the theme")
	rootCmd.PersistentFlags().IntP("width", "w", 0,
		"word wrap `width`")
	rootCmd.PersistentFlags().BoolP("no-pager", "n", false,
		"don't use pager")
	rootCmd.PersistentFlags().Bool("debug", false,
		"show debugging information")
	rootCmd.PersistentFlags().SortFlags = false
	viper.BindPFlag("theme",
		rootCmd.PersistentFlags().Lookup("theme"))
	viper.BindPFlag("width",
		rootCmd.PersistentFlags().Lookup("width"))
	viper.BindPFlag("no-pager",
		rootCmd.PersistentFlags().Lookup("no-pager"))
	viper.BindPFlag("debug",
		rootCmd.PersistentFlags().Lookup("debug"))
}

func initTerm() (w int) {
	var err error
	w, _, err = term.GetSize(0)
	if err != nil {
		color.Yellow("cannot get terminal size")
		w = 80
	}
	viper.SetDefault("termwidth", w)
	return
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
	viper.AutomaticEnv() // EMD_* envvars
	viper.ReadInConfig()
	debug(fmt.Sprintf("%#v\n", viper.AllSettings()))
}
