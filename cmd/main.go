package main

import (
	"fmt"

	"github.com/me4hit/distributed-go-cache/internal/cache"
)

func main() {
	fmt.Println("Hello, World!")

	lru := cache.NewLRUCache(2)

	lru.Set("key1", "value1")

	lru.Set("key2", "value2")

	if value, ok := lru.Get("key1"); ok {
		fmt.Println("key1:", value)
	} else {
		fmt.Println("key1 not found")
	}

	lru.Set("key3", "value3")

	if value, ok := lru.Get("key2"); ok {
		fmt.Println("key2:", value)
	} else {
		fmt.Println("key2 not found")
	}

	if value, ok := lru.Get("key3"); ok {
		fmt.Println("key3:", value)
	} else {
		fmt.Println("key3 not found")
	}
}
