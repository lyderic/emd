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
	return renderBytes(content)
}

func renderBytes(content []byte) (err error) {
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
 * If cw not set: use tw (minus a little margin)
 *
 * If cw is set:
 *   - if cw > tw: use tw (minus a little margin)
 *   - if cw < tw: use cw
 */
func getOptimalWidth() (ow int) {
	tw := viper.GetInt("termwidth")
	cw := viper.GetInt("configwidth")
	ow = tw - 2              // -2 = little margin, default case
	if cw > 0 && cw < tw-1 { // -1 for margin
		ow = cw
	}
	debug(fmt.Sprintf("Width: tw=%d cw=%d ow=%d\n", tw, cw, ow))
	return
}
