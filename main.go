package main

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	data  map[int]string
	cache Cacher
}

func (s *Store) GetFromCache(key int) (string, bool) {
	val, ok := s.cache.Get(key)
	if ok {
		fmt.Println("Cache hit")
		return val, ok
	}
	return "", false
}

func (s *Store) Get(key int) (string, error) {
	val, ok := s.cache.Get(key)
	if ok {
		// bust cache
		s.cache.Remove(key)
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

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:5432",
		Password: "",
		DB:       0,
	})

	s := NewStore(NewRedisCache(client))
	// for q := 1; q < 10; q++ {
	t1, err := s.Get(1)
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	} else {
		fmt.Println(t1)
	}
	// }

}