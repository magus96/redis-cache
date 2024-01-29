package main

import "fmt"

type Store struct {
	data  map[int]string
	cache Cacher
}

func (s *Store) Get(key int) (string, error) {
	val, ok := s.cache.Get(key)
	if ok {
		// bust cache
		if err := s.cache.Remove(key); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Returning from cache")
		return val, nil
	}
	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("Key not found: %d", key)
	}
	if err := s.cache.Set(key, val); err != nil {
		return "", err
	}
	fmt.Println("Returning from internal storage")
	return val, nil
}

func NewStore(c Cacher) *Store {
	data := map[int]string{
		1: "Hellow",
		2: "World",
	}
	return &Store{
		data:  data,
		cache: c,
	}
}
