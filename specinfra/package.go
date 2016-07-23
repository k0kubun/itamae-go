package specinfra

import (
	"fmt"
)

func CheckPackageIsInstalled(name string, version string) string {
	switch platform {
	case "arch":
		return fmt.Sprintf("pacman -Q %s || pacman -Qg %s", name, name)
	default:
		panic("specinfra.CheckPackageIsInstalled is not supported")
	}
}

func Install(name string, version string, options string) string {
	switch platform {
	case "arch":
		return fmt.Sprintf("pacman -S --noconfirm %s %s", options, name)
	default:
		panic("specinfra.Install is not supported")
	}
}
