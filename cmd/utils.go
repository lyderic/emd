package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var themes = [5]string{
	"dark",
	"light",
	"dracula",
	"pink",
	"notty",
}

func longHelp() (output string) {
	var b strings.Builder
	fmt.Fprintf(&b, "\n\033[1m%s %s: %s\033[m\n",
		PROGNAME, VERSION, DESCRIPTION)
	fmt.Fprintln(&b, "\nAvailable themes:")
	for _, theme := range themes {
		fmt.Fprintln(&b, " •", theme)
	}
	return b.String()
}

func debug(message string) {
	if viper.GetBool("debug") {
		fmt.Fprintf(os.Stderr, "\033[7m%s\033[m\n", message)
	}
}
