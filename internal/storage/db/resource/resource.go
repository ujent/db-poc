package resource

type Resources interface {
	AddResource(*Resource) error
	Resourcies() ([]*Resource, error)
}

type Resource struct {
	ID string
}

type ResourceStorage struct{}

func (st *ResourceStorage) AddResource(*Resource) error {
	return nil
}

func (st *ResourceStorage) Resourcies() ([]*Resource, error) {
	return nil, nil
}
