package storage

import "truesrc/core"

type MemoryStore struct {
	data map[string]core.Narrative
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: map[string]core.Narrative{},
	}
}

func (m *MemoryStore) Save(n core.Narrative) {
	m.data[n.ID] = n
}

func (m *MemoryStore) Get(id string) (core.Narrative, bool) {
	n, ok := m.data[id]
	return n, ok
}
