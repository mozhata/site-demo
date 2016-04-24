package lib

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/golang/glog"
)

// RunCommandAndWait runs the command and return error attached with output.
func RunCommandAndWait(cmd *exec.Cmd) error {
	if cmd.Stdout == nil {
		cmd.Stdout = os.Stdout
	}
	if cmd.Stderr == nil {
		cmd.Stderr = os.Stderr
	}

	buf := bytes.Buffer{}
	cmd.Stdout = io.MultiWriter(cmd.Stdout, &buf)
	cmd.Stderr = io.MultiWriter(cmd.Stderr, &buf)
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("Output:\n%s\nError: %v", buf.Bytes(), err)
	}
	return nil
}

// BuildBinary builds the binary and return its path.
func BuildBinary(projectDir ProjectDir, name string, mainFile string) (string, error) {
	glog.Infof("Building %s", name)
	c := exec.Command("godep", "go", "install", mainFile)
	c.Dir = projectDir.String()
	binDir := path.Join(projectDir.TmpDir(), name)
	c.Env = append(os.Environ(), fmt.Sprint("GOBIN=", binDir))
	err := RunCommandAndWait(c)
	if err != nil {
		return "", err
	}

	return path.Join(binDir, strings.TrimSuffix(path.Base(mainFile), ".go")), nil
}
