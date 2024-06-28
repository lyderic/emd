package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "emd",
	Short:   "A basic markdown viewer for the command line",
	Version: VERSION,
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetBool("list-themes") {
			listThemes()
			return nil
		}
		if viper.GetBool("pager") {
			color.Yellow("pager not implemented yet")
		}
		if len(args) < 1 {
			return fmt.Errorf("missing file argument!")
		}
		return render(args[0])
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

const VERSION = "0.0.2"

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.emd.yaml)")
	rootCmd.PersistentFlags().StringP("theme", "t", "dark", "theme")
	rootCmd.PersistentFlags().BoolP("list-themes", "l", false, "list available themes")
	rootCmd.PersistentFlags().IntP("width", "w", 80, "word wrap width")
	rootCmd.PersistentFlags().BoolP("pager", "p", false, "use pager")
	viper.BindPFlag("theme", rootCmd.PersistentFlags().Lookup("theme"))
	viper.BindPFlag("list-themes", rootCmd.PersistentFlags().Lookup("list-themes"))
	viper.BindPFlag("width", rootCmd.PersistentFlags().Lookup("width"))
	viper.BindPFlag("pager", rootCmd.PersistentFlags().Lookup("pager"))
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
