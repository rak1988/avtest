package lrucache

import (
	"container/list"
	"fmt"
	"sync"
)

/*
 *	Implements an LRU cache
 */
type LRUCache struct {
	hashMap  map[string]interface{}
	dlist    *list.List
	capacity int64
	count    int64
	mutex    *sync.Mutex
}

// CacheElement is the element struct in the LRU queue
type CacheElement struct {
	Key   string
	Value interface{}
}

//Init function for LRU Cache
func (lc *LRUCache) Init(capacity int64) *LRUCache {
	lc.hashMap = make(map[string]interface{})
	lc.capacity = capacity
	lc.count = 0
	lc.dlist = new(list.List)
	lc.mutex = new(sync.Mutex)
	return lc
}

// Get value of key from cache
func (lc *LRUCache) Get(key string) interface{} {
	if lc.hashMap[key] != nil {
		ele := lc.hashMap[key].(*list.Element)
		value := ele.Value.(*CacheElement).Value
		// apply a lock before moving the element to front
		lc.mutex.Lock()
		//fmt.Println(lc.hashMap[key].(*list.Element).Value)
		lc.dlist.MoveToFront(ele)
		//lc.PrintCache()
		lc.mutex.Unlock()

		return value
	}
	return nil
}

func (lc *LRUCache) evictIfRequired() {
	if int64(lc.dlist.Len()) > (lc.capacity) {
		lc.mutex.Lock()
		delete(lc.hashMap, lc.dlist.Back().Value.(*CacheElement).Key)
		lc.dlist.Remove(lc.dlist.Back())
		lc.count = lc.count - 1
		lc.mutex.Unlock()
	}
}

// Put new key, value into cache
func (lc *LRUCache) Put(key string, value interface{}) bool {
	newCacheElement := new(CacheElement)
	newCacheElement.Key = key
	newCacheElement.Value = value
	lc.mutex.Lock()
	newElement := lc.dlist.PushFront(newCacheElement)
	lc.hashMap[key] = newElement
	lc.count = lc.count + 1
	lc.mutex.Unlock()
	lc.evictIfRequired()
	return true
}

// PrintCache prints elements in cache
func (lc *LRUCache) PrintCache() {
	nodeptr := lc.dlist.Front()
	for nodeptr.Next() != nil {
		fmt.Println(nodeptr.Value.(*CacheElement).Value)
		nodeptr = nodeptr.Next()
	}
}

func (lc *LRUCache) GetDListLength() int {
	return lc.dlist.Len()
}

func (lc *LRUCache) GetFrontElement() interface{} {
	return lc.dlist.Front().Value.(*CacheElement).Key
}
