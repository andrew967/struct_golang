package hashTable

import (
	"fmt"
	"hash/fnv"
)

type Node struct {
	key   string
	value interface{}
}

type HashMap struct {
	size   int
	bucket []*[]Node
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		size:   size,
		bucket: make([]*[]Node, size),
	}
}

func (hm *HashMap) hash(key string) int {
	h := fnv.New32a()
	if _, err := h.Write([]byte(key)); err != nil {
		panic(err)
	}
	return int(h.Sum32()) % hm.size
}

func (h *HashMap) Add(key string, value interface{}) {
	index := h.hash(key)
	if h.bucket[index] == nil {
		h.bucket[index] = &[]Node{}
	}
	*h.bucket[index] = append(*h.bucket[index], Node{key: key, value: value})
}

func (h *HashMap) Remove(key string) error {
	index := h.hash(key)
	if h.bucket[index] != nil {
		for i, node := range *h.bucket[index] {
			if node.key == key {
				*h.bucket[index] = append((*h.bucket[index])[:i], (*h.bucket[index])[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("Key not found to delete: %s", key)
}

func (h *HashMap) Get(key string) (value interface{}, err error) {
	index := h.hash(key)
	if h.bucket[index] != nil {
		for _, node := range *h.bucket[index] {
			if node.key == key {
				return node.value, nil
			}
		}
	}
	return nil, fmt.Errorf("Key not found: %s", key)
}
