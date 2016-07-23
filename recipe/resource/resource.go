package resource

type Resource interface {
	Apply()
}

var resources = make([]Resource, 0)

func Register(resource Resource) {
	resources = append(resources, resource)
}

func Registered() []Resource {
	return resources
}
