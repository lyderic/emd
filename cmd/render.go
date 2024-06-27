package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
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
		glamour.WithStandardStyle(c.Style),
		glamour.WithWordWrap(c.WordWrap),
	)
	out, err := r.RenderBytes(raw)
	if err != nil {
		return
	}
	fmt.Print(string(out))
	return
}
