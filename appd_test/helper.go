package appdtest

import (
	"os/exec"
	"path"
	"path/filepath"
)

var rootDir string

func init() {
	var err error
	rootDir, err = filepath.Abs("..")

	if err != nil {
		panic(err)
	}
}

func exec_appd() string {
	cmdPath := path.Join(rootDir, "out", "appd")
	cmd := exec.Command(cmdPath)

	out, err := cmd.CombinedOutput()

	if err != nil {
		panic("could not run")
	}

	return string(out)

}
