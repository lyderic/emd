package cmd

import (
	"fmt"
	"os"
)

var themes = [5]string{
	"dark",
	"light",
	"dracula",
	"pink",
	"notty",
}

func warn(message string) {
	fmt.Fprintf(os.Stderr, "\033[7mWarning: %s\033[m\n", message)
}
