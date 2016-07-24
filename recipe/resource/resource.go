package resource

type Resource interface {
	Apply()
	DryRun()
}

var resources = make([]Resource, 0)

func Register(resource Resource) {
	resources = append(resources, resource)
}

func Registered() []Resource {
	return resources
}
