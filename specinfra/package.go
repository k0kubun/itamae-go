package specinfra

import (
	"fmt"
)

func CheckPackageIsInstalled(name string, version string) string {
	// XXX: Consider version...
	switch platform {
	case "arch":
		return fmt.Sprintf("pacman -Q %s || pacman -Qg %s", name, name)
	case "darwin":
		return fmt.Sprintf("brew list | grep %s", name)
	case "debian":
		return fmt.Sprintf("dpkg-query -f '${Status}' -W %s | grep -E '^(install|hold) ok installed$'", name)
	default:
		panic("specinfra.CheckPackageIsInstalled is not supported")
	}
}

func Install(name string, version string, options string) string {
	// XXX: Consider version...
	switch platform {
	case "arch":
		return fmt.Sprintf("pacman -S --noconfirm %s %s", options, name)
	case "darwin":
		return fmt.Sprintf("brew install %s '%s'", options, name)
	case "debian":
		return fmt.Sprintf("DEBIAN_FRONTEND='noninteractive' apt-get -y -o Dpkg::Options::='--force-confdef' -oDpkg::Options::='--force-confold' %s install %s", options, name)
	default:
		panic("specinfra.Install is not supported")
	}
}

func Remove(name string, options string) string {
	switch platform {
	case "arch":
		return fmt.Sprintf("pacman -R --noconfirm %s %s", options, name)
	case "darwin":
		return fmt.Sprintf("brew uninstall %s '%s'", options, name)
	case "debian":
		return fmt.Sprintf("DEBIAN_FRONTEND='noninteractive' apt-get -y %s remove %s", options, name)
	default:
		panic("specinfra.Install is not supported")
	}
}
