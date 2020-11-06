package storage

type Storage interface {
	Store(short string, url string) error
	Get(short string) (*string, error)
}
