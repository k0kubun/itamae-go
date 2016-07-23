package specinfra

import (
	"fmt"
)

func CheckPackageIsInstalled(name string, version string) string {
	return fmt.Sprintf("pacman -Q %s || pacman -Qg %s", name, name)
}

func Install(name string, version string, options string) string {
	return fmt.Sprintf("pacman -S --noconfirm %s %s", options, name)
}
