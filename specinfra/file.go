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

func CheckFileIsDirectory(file string) string {
	return fmt.Sprintf("test -d %s", file)
}
