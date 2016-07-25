package specinfra

import (
	"fmt"
)

func CreateAsDirectory(file string) string {
	return fmt.Sprintf("mkdir -p %s", file)
}

func RemoveFile(file string) string {
	return fmt.Sprintf("rm -rf %s", file)
}

func CheckFileIsFile(file string) string {
	return fmt.Sprintf("test -f %s", file)
}

func CheckFileIsDirectory(file string) string {
	return fmt.Sprintf("test -d %s", file)
}

func LinkFileTo(link string, target string, force bool) string {
	option := "-s"
	if force {
		option += "f"
	}
	return fmt.Sprintf("ln %s %s %s", option, target, link)
}
