package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var themes = [5]string{
	"dark",
	"light",
	"dracula",
	"pink",
	"notty",
}

func debug(message string) {
	if viper.GetBool("debug") {
		fmt.Fprintf(os.Stderr, "\033[7m%s\033[m\n", message)
	}
}
