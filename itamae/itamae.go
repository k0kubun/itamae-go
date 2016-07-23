package itamae

import (
	"fmt"

	. "github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/k0kubun/pp"
)

func Apply(resources []Resource) {
	fmt.Println("--- itamae apply (stubbed) ---")
	for _, resource := range resources {
		resource.Apply()
	}
	pp.Println(resources)
}

func DryRun(resources []Resource) {
	fmt.Println("--- itamae dry-run (stubbed) ---")
	pp.Println(resources)
}
