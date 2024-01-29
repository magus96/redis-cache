package main

type Cacher interface {
	Set(int, string) error
	Get(int) (string, bool)
	Remove(int) error
}

type NopCache struct{}

func (c NopCache) Get(key int) (string, bool) {
	return "", false
}

func (c NopCache) Set(key int, val string) error {
	return nil
}

func (c NopCache) Remove(key int) error {
	return nil
}
