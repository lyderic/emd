package cmd

import (
	"fmt"
	"os"

	. "github.com/lyderic/tools"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/viper"
)

func render(path string) (err error) {

	if _, err = os.Stat(path); os.IsNotExist(err) {
		return
	}

	var raw []byte
	if raw, err = os.ReadFile(path); err != nil {
		return
	}
	r, _ := glamour.NewTermRenderer(
		glamour.WithStandardStyle(viper.GetString("theme")),
		glamour.WithWordWrap(viper.GetInt("width")),
	)
	out, err := r.RenderBytes(raw)
	if err != nil {
		return
	}
	if viper.GetBool("no-pager") {
		fmt.Println(string(out))
		return
	}
	return Less(string(out))
}
