package main

import (
	"avitest/lrucache"
	"strconv"
)

func main() {
	cache := new(lrucache.LRUCache).Init(100000)

	// cache.Put("asdf1", 1)
	// cache.Put("asdf2", 2)
	// cache.Put("asdf3", 3)
	// cache.Put("asdf4", 4)
	// cache.Put("asdf3", 4)
	// cache.Put("asdf5", 5)
	// cache.Put("asdf6", 6)
	// cache.Get("asdf3")
	// fmt.Println(cache.Get("asdf4"))
	// cache.PrintCache()

	for i := 0; i < 100000; i++ {
		//fmt.Println(strconv.Itoa(i))
		cache.Put(strconv.Itoa(i), i)
	}

	cache.PrintCache()
}
