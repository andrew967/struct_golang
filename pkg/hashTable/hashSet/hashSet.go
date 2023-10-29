package hashSet

import (
	"fmt"
	"hash/fnv"
)

type HashSet struct {
	size   int
	bucket []*[]string
}

func NewHashSet(size int) *HashSet {
	return &HashSet{
		size:   size,
		bucket: make([]*[]string, size),
	}
}

func (hs *HashSet) hash(key string) int {
	h := fnv.New32a()
	if _, err := h.Write([]byte(key)); err != nil {
		panic(err)
	}
	return int(h.Sum32()) % hs.size
}

func (h *HashSet) Add(key string) {
	index := h.hash(key)
	if h.bucket[index] == nil {
		h.bucket[index] = &[]string{}
	}
	for _, str := range *h.bucket[index] {
		if key == str {
			return
		}
	}
	*h.bucket[index] = append(*h.bucket[index], key)
}

func (h *HashSet) Remove(key string) error {
	index := h.hash(key)
	if h.bucket[index] != nil {
		for i, str := range *h.bucket[index] {
			if str == key {
				*h.bucket[index] = append((*h.bucket[index])[:i], (*h.bucket[index])[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("Key not found to delete: %s", key)
}

func (h *HashSet) Contains(key string) bool {
	index := h.hash(key)
	if h.bucket[index] != nil {
		for _, str := range *h.bucket[index] {
			if str == key {
				return true
			}
		}
	}
	return false
}

func (h *HashSet) IsEmpty() bool {
	for _, bucket := range h.bucket {
		if bucket != nil && len(*bucket) > 0 {
			return false
		}
	}
	return true
}

func (h *HashSet) Size() int {
	return len(h.bucket)
}

func (h *HashSet) Clear() {
	if !h.IsEmpty() {
		for i := range h.bucket {
			h.bucket[i] = nil
		}
		h.bucket = []*[]string{}
	}
}
