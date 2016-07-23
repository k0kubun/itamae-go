package utils

import (
	"os"
	"os/exec"
)

func Execute(str string) bool {
	cmd := exec.Command(shell(), "-c", str)
	err := cmd.Run()
	return err == nil
}

func shell() string {
	ret := os.Getenv("SHELL")
	if len(ret) == 0 {
		ret = "/bin/sh"
	}
	return ret
}
