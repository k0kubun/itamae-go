package specinfra

import (
	"log"
	"os/exec"
	"regexp"
)

var platform string

func init() {
	platform = Platform()
}

func Platform() string {
	if regexp.MustCompile("(?i)darwin").MatchString(unameSr()) {
		return "darwin"
	} else if execute("ls /etc/arch-release") {
		return "arch"
	} else {
		panic("Sorry, this distribution is not supported yet.")
	}
}

func unameSr() string {
	buf, err := exec.Command("uname", "-sr").Output()
	if err != nil {
		log.Fatal()
	}
	return string(buf)
}

func execute(str string) bool {
	cmd := exec.Command("/bin/sh", "-c", str)
	err := cmd.Run()
	return err == nil
}
