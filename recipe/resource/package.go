package resource

import (
	"fmt"
)

type Package struct {
	Action []string
	User   string
	Cwd    string
	OnlyIf string
	NotIf  string

	Name    string
	Version string
	Options string
}

func (p *Package) Apply() {
	fmt.Printf("Install package: %s\n", p.Name)
}
