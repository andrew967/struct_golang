package hashSet

import (
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	hs := NewHashSet(5)

	hs.Add("key1")
	if !hs.Contains("key1") {
		t.Error("Expected key1 to be in the set")
	}

	hs.Add("key1")
	if !hs.Contains("key1") {
		t.Error("Expected key1 to still be in the set after adding a duplicate")
	}

	keys := []string{"key2", "key3", "key4"}
	for _, key := range keys {
		hs.Add(key)
		if !hs.Contains(key) {
			t.Errorf("Expected %s to be in the set", key)
		}
	}
}

func TestHashSet_Remove(t *testing.T) {
	hs := NewHashSet(5)

	err := hs.Remove("nonexistentKey")
	if err == nil || err.Error() != "Key not found to delete: nonexistentKey" {
		t.Errorf("Expected error: Key not found, Got: %v", err)
	}

	hs.Add("key1")
	err = hs.Remove("key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if hs.Contains("key1") {
		t.Error("Expected key1 to be removed from the set")
	}
}

func TestHashSet_Contains(t *testing.T) {
	hs := NewHashSet(5)

	if hs.Contains("nonexistentKey") {
		t.Error("Expected key not to be in the set")
	}

	hs.Add("key1")
	if !hs.Contains("key1") {
		t.Error("Expected key1 to be in the set")
	}
}

func TestHashSet_IsEmpty(t *testing.T) {
	hs := NewHashSet(5)

	if !hs.IsEmpty() {
		t.Error("Expected set to be empty")
	}

	hs.Add("key1")
	if hs.IsEmpty() {
		t.Error("Expected set not to be empty")
	}
}

func TestHashSet_Size(t *testing.T) {
	hs := NewHashSet(5)

	if size := hs.Size(); size != 5 {
		t.Errorf("Expected size: 5, Got: %v", size)
	}

	hs.Add("key1")
	if size := hs.Size(); size != 5 {
		t.Errorf("Expected size: 5, Got: %v", size)
	}
}

func TestHashSet_Clear(t *testing.T) {
	hs := NewHashSet(5)

	hs.Clear()
	if !hs.IsEmpty() {
		t.Error("Expected set to be empty after clearing")
	}

	hs.Add("key1")
	hs.Clear()
	if !hs.IsEmpty() {
		t.Error("Expected set to be empty after clearing")
	}
}
