package specinfra

import (
	"fmt"
)

func StartService(service string) string {
	switch platform {
	case "arch":
		return fmt.Sprintf("systemctl start %s", service)
	case "debian":
		return fmt.Sprintf("service %s start", service)
	default:
		panic("specinfra.StartServie is not supported")
	}
}

func StopService(service string) string {
	switch platform {
	case "arch":
		return fmt.Sprintf("systemctl stop %s", service)
	case "debian":
		return fmt.Sprintf("service %s stop", service)
	default:
		panic("specinfra.StartServie is not supported")
	}
}

func EnableService(service string) string {
	switch platform {
	case "arch":
		return fmt.Sprintf("systemctl enable %s", service)
	case "debian":
		return fmt.Sprintf("update-rc.d %s defaults", service)
	default:
		panic("specinfra.StartServie is not supported")
	}
}

func DisableService(service string) string {
	switch platform {
	case "arch":
		return fmt.Sprintf("systemctl enable %s", service)
	case "debian":
		return fmt.Sprintf("update-rc.d -f %s remove", service)
	default:
		panic("specinfra.StartServie is not supported")
	}
}
