package main

import (
	"container/list"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CacheItem struct {
	key        string
	value      interface{}
	expiration *time.Time
}

type LRUCache struct {
	capacity  int
	mutex     sync.Mutex
	items     map[string]*list.Element
	evictList *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity:  capacity,
		items:     make(map[string]*list.Element),
		evictList: list.New(),
	}
}

func (c *LRUCache) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// If the item already exists, update its value and move it to the front
	if elem, ok := c.items[key]; ok {
		c.evictList.MoveToFront(elem)
		item := elem.Value.(*CacheItem)
		item.value = value
		expiration := time.Now().Add(duration)
		item.expiration = &expiration
		return
	}

	// Add a new item. If the cache is full, remove the least recently used item
	if c.evictList.Len() == c.capacity {
		c.evictOldest()
	}
	expiration := time.Now().Add(duration)
	item := &CacheItem{key: key, value: value, expiration: &expiration}
	elem := c.evictList.PushFront(item)
	c.items[key] = elem
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.items[key]; ok {
		item := elem.Value.(*CacheItem)

		// Check if the item has expired
		if item.expiration != nil && item.expiration.Before(time.Now()) {
			c.removeElement(elem)
			return nil, false
		}

		c.evictList.MoveToFront(elem)
		return item.value, true
	}
	return nil, false
}

func (c *LRUCache) evictOldest() {
	elem := c.evictList.Back()
	if elem != nil {
		c.removeElement(elem)
	}
}

func (c *LRUCache) removeElement(elem *list.Element) {
	c.evictList.Remove(elem)
	item := elem.Value.(*CacheItem)
	delete(c.items, item.key)
}

func main() {
	cache := NewLRUCache(1024) // Initialize the cache with a capacity of 1024 items

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/set", func(c *gin.Context) {
		key := c.Query("key")
		value := c.Query("value")
		duration, err := time.ParseDuration(c.Query("duration"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duration format"})
			return
		}

		cache.Set(key, value, duration)
		c.JSON(http.StatusOK, gin.H{"message": "Value set successfully"})
	})

	r.GET("/get", func(c *gin.Context) {
		key := c.Query("key")
		value, ok := cache.Get(key)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found or expired"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	r.Run()
}
