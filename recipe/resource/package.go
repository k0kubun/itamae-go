package resource

import (
	"fmt"

	. "github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/k0kubun/itamae-go/specinfra"
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
	if Execute(specinfra.CheckPackageIsInstalled(p.Name, p.Version)) {
		fmt.Println("Skip installation:", p.Name)
	} else {
		fmt.Println("Install:", p.Name)
		Execute(specinfra.Install(p.Name, p.Version, p.Options))
	}
}
