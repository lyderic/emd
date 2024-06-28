package cmd

import "fmt"

var themes = [4]string{
	"dark",
	"light",
	"dracula",
	"pink",
}

func listThemes() {
	for _, theme := range themes {
		fmt.Println("•", theme)
	}
}
