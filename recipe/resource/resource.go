package resource

type Resource interface {
	Apply()
	DryRun()
	ShouldSkip() bool
}

var resources = make([]Resource, 0)

func Register(resource Resource) {
	resources = append(resources, resource)
}

func Registered() []Resource {
	return resources
}
