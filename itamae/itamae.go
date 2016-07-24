package itamae

import (
	. "github.com/k0kubun/itamae-go/recipe/resource"
)

func Apply(resources []Resource) {
	for _, resource := range resources {
		resource.Apply()
	}
}

func DryRun(resources []Resource) {
	for _, resource := range resources {
		resource.DryRun()
	}
	pp(resources)
}
