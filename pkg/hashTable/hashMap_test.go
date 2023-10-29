package hashTable

import (
	"testing"
)

func TestHashMap_Add(t *testing.T) {
	hm := NewHashMap(5)

	hm.Add("key1", 42)
	if value, err := hm.Get("key1"); err != nil || value.(int) != 42 {
		t.Errorf("Expected value: 42, Got : %v", value)
	}

	hm.Add("key2", "value2")
	if value, err := hm.Get("key2"); err != nil || value.(string) != "value2" {
		t.Errorf("Expected value: value2, Got : %v", value)
	}

}

func TestHashMap_Remove(t *testing.T) {
	hm := NewHashMap(5)

	err := hm.Remove("nonexistentKey")
	if err == nil || err.Error() != "Key not found to delete: nonexistentKey" {
		t.Errorf("Expected errorL Key not found, Got : %v", err)
	}

	hm.Add("key1", 42)
	err = hm.Remove("key1")
	if err != nil {
		t.Errorf("Expected error nil, Got: %v", err)
	}

	if _, err := hm.Get("key1"); err == nil || err.Error() != "Key not found: key1" {
		t.Errorf("Expected error: Key not found, Got: %v", err)
	}
}

func TestHashMap_Get(t *testing.T) {
	hm := NewHashMap(5)

	value, err := hm.Get("nonexistentKey")
	if err == nil || err.Error() != "Key not found: nonexistentKey" {
		t.Errorf("Expected error: Key not found, Got: %v", err)
	}
	if value != nil {
		t.Errorf("Expected value: nil, Got: %v", value)
	}

	hm.Add("key1", 42)
	value, err = hm.Get("key1")
	if err != nil || value.(int) != 42 {
		t.Errorf("Expected value: 42, Got: %v", value)
	}
}
