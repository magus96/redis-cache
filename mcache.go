package main

import (
	"sync"
)

type mcache struct {
	lock sync.RWMutex
	data map[int]string
}

func NewMCache() *mcache {
	return &mcache{data: make(map[int]string)}
}

func (m *mcache) Get(key int) (string, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	val, ok := m.data[key]
	if !ok {
		return "", false
	}
	return val, true
}

func (m *mcache) Set(key int, val string) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data[key] = val
	return nil
}

func (m *mcache) Remove(key int) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.data, key)
	return nil
}
