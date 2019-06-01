package lrucache

import (
	"testing"
)

func TestLRUCacheSimpleGet(t *testing.T) {
	cache := new(LRUCache).Init(3)
	cache.Put("a", 1)
	if cache.Get("a") != 1 {
		t.Errorf("value incorrect for key a")
	}
}

func TestLRUCacheCapacityCheck(t *testing.T) {
	cache := new(LRUCache).Init(3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	cache.Put("d", 4)

	if cache.GetDListLength() > 3 {
		t.Errorf("Dlist length should not be more than specified capacity")
	}
}

func TestLRUCacheLeastRecentlyUsedEvicted(t *testing.T) {
	cache := new(LRUCache).Init(3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	cache.Put("d", 4)
	if cache.Get("a") != nil {
		t.Errorf("least recently used is not evicted after adding new element")
	}
	if cache.Get("b") == nil || cache.Get("c") == nil || cache.Get("d") == nil {
		t.Errorf("except the least recently used eviction, rest of the elements should be in the cache")
	}
}

func TestLRUCacheCheckLastElementPutIstheTopElement(t *testing.T) {
	cache := new(LRUCache).Init(3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	cache.Put("d", 4)

	if cache.GetFrontElement() != "d" {
		t.Errorf("front element is not the last inserted element")
	}

}
