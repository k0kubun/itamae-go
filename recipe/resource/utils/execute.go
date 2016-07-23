package utils

import (
	"os"
	"os/exec"

	"github.com/k0kubun/itamae-go/logger"
)

func Execute(str string) bool {
	logger.Debug("Executing `" + str + "`")
	cmd := exec.Command(shell(), "-c", str)
	err := cmd.Run()
	return err == nil
}

func Run(str string) {
	result := Execute(str)
	if !result {
		logger.Error("Command `" + str + "` failed.")
	}
}

func shell() string {
	ret := os.Getenv("SHELL")
	if len(ret) == 0 {
		ret = "/bin/sh"
	}
	return ret
}
