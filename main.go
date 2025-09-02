package main

import (
	"hash/fnv"
	"sync"
)

type InMemoryCache struct {
	shards []Shard
}
type Shard struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewInMemoryCache(numShards int) *InMemoryCache {
	shards := make([]Shard, 0, numShards)
	for i := 0; i < numShards; i++ {
		shards = append(shards, Shard{
			data: make(map[string]string),
		})
	}
	return &InMemoryCache{
		shards: shards,
	}
}

func hasher(k string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(k))
	return int(h.Sum32())
}

func (c *InMemoryCache) Set(k string, v string) {
	shardID := hasher(k) % len(c.shards)
	println(shardID)
	c.shards[shardID].Set(k, v)
}
func (c *InMemoryCache) Get(k string) (string, bool) {
	shardID := hasher(k) % len(c.shards)
	return c.shards[shardID].Get(k)
}

type Cache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}

func (s *Shard) Set(k string, v string) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.data[k] = v

}

func (s *Shard) Get(k string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	data, ok := s.data[k]
	return data, ok
}

func main() {
	cache := NewInMemoryCache(5)
	cache.Set("foo", "bar")
	cache.Set("foo1", "bar")
	cache.Set("baz", "qux")
	cache.Set("quux", "quuz")
}
