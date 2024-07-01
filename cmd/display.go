package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	if _, err := exec.LookPath(PAGER); err != nil {
		debug(fmt.Sprintf("Command %q not found!", PAGER))
		viper.SetDefault("no-pager", true)
		return
	}
}

func display(buffer []byte) (err error) {
	if viper.GetBool("no-pager") {
		fmt.Println(string(buffer))
		return
	}
	return less(string(buffer))
}

func less(message string) error {
	cmd := exec.Command(PAGER, "-FRIX")
	cmd.Stdin = strings.NewReader(message)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}
