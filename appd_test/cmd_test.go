package appdtest

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Printf("Dir %s\n", rootDir)
	makeBuild := exec.Command("make", fmt.Sprintf("-C%s", rootDir), "build")

	err := makeBuild.Run()

	if err != nil {
		fmt.Printf("make error: %s\n", err)
		fmt.Printf("aborting tests")
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestNoArgs(t *testing.T) {
	out := exec_appd()

	require.Equal(t, "", out, "runs with no output")
}
