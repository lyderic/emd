package cmd

import (
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/viper"
)

func render(path string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return
	}
	var content []byte
	if content, err = os.ReadFile(path); err != nil {
		return
	}
	r, _ := glamour.NewTermRenderer(
		glamour.WithStandardStyle(viper.GetString("theme")),
		glamour.WithWordWrap(viper.GetInt("width")),
	)
	rendered, err := r.RenderBytes(content)
	if err != nil {
		return
	}
	return display(rendered)
}
