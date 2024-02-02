package main

import (
	"fmt"
	"log"
)

func main() {

	cache := NewMCache()
	store := NewStore(cache)
	for q := 0; q < 10; q++ {
		val, err := store.Get(1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(val)
	}
}
