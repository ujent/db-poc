package policy

type Policies interface {
	AddPolicy(*Policy) error
	Policies() ([]*Policy, error)
}

type Policy struct {
	ID string
}

type PolicyStorage struct {
}

func (st *PolicyStorage) AddPolicy(*Policy) error {
	return nil
}

func (st *PolicyStorage) Policies() ([]*Policy, error) {
	return nil, nil
}
