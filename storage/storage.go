package storage

type Storage interface {
}

type Settings struct {
	Dialect string
}

func NewStorage(settings *Settings) {

}
