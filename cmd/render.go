package cmd

import (
	"fmt"
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
	width := getOptimalWidth()
	r, _ := glamour.NewTermRenderer(
		glamour.WithStandardStyle(viper.GetString("theme")),
		glamour.WithWordWrap(width),
	)
	rendered, err := r.RenderBytes(content)
	if err != nil {
		return
	}
	return display(rendered)
}

/*
 * Given:
 *  tw = terminal width (computed by x/term lib in root.go), and
 *  cw = width given by configuration (--width flag,
 *    EMD_WIDTH envvar, etc.)
 *
 * If cw not set: use tw
 *
 * If cw is set:
 *   - if cw > tw: use tw
 *   - if cw < tw: use cw
 */
func getOptimalWidth() (w int) {
	tw := viper.GetInt("termwidth")
	cw := viper.GetInt("width")
	debug(fmt.Sprintf("Width:\n Terminal = %d\n Configuration = %d", tw, cw))
	if cw > tw {
		return tw
	}
	if cw < tw {
		return cw
	}
	return cw
}
