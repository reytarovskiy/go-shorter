package storage

type MemoryStorage struct {
	shorts map[string]string
}

func NewMemoryStorage(urls map[string]string) *MemoryStorage {
	return &MemoryStorage{
		shorts: urls,
	}
}

func (m *MemoryStorage) Store(short string, url string) {
	m.shorts[short] = url
}

func (m *MemoryStorage) Get(short string) *string {
	url, exists := m.shorts[short]
	if !exists {
		return nil
	}

	return &url
}
