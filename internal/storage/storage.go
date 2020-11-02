package storage

type Storage interface {
	Store(short string, url string)
	Get(short string) *string
}
