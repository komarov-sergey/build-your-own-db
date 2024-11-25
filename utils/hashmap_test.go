package utils

import (
	"testing"
)

func TestHashMapSetAndGet(t *testing.T) {
	hashMap := NewHashMap(32)
	hashMap.Set("one", 1)
	hashMap.Set("two", 2)

	if value, exists := hashMap.Get("one"); !exists || value != 1 {
		t.Errorf("Expected value 1 for key 'one', but got %v", value)
	}

	if value, exists := hashMap.Get("two"); !exists || value != 2 {
		t.Errorf("Expected value 2 for key 'two', but got %v", value)
	}

	if value, exists := hashMap.Get("three"); exists {
		t.Errorf("Expected no value for key 'three', but got %v", value)
	}
}

func TestHashMapRemove(t *testing.T) {
	hashMap := NewHashMap(32)
	hashMap.Set("one", 1)
	hashMap.Set("two", 2)

	hashMap.Remove("one")
	if value, exists := hashMap.Get("one"); exists {
		t.Errorf("Expected no value for key 'one', but got %v", value)
	}

	if value, exists := hashMap.Get("two"); !exists || value != 2 {
		t.Errorf("Expected value 2 for key 'two', but got %v", value)
	}
}

func TestHashMapSize(t *testing.T) {
	hashMap := NewHashMap(32)
	hashMap.Set("one", 1)
	hashMap.Set("two", 2)

	if hashMap.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", hashMap.Size())
	}

	hashMap.Remove("one")
	if hashMap.Size() != 1 {
		t.Errorf("Expected size 1, but got %d", hashMap.Size())
	}
}

func TestHashMapOverwrite(t *testing.T) {
	hashMap := NewHashMap(32)
	hashMap.Set("one", 1)
	hashMap.Set("one", 2)

	if value, exists := hashMap.Get("one"); !exists || value != 2 {
		t.Errorf("Expected value 2 for key 'one', but got %v", value)
	}
}

func TestHashMapEmpty(t *testing.T) {
	hashMap := NewHashMap(32)

	if hashMap.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", hashMap.Size())
	}

	if value, exists := hashMap.Get("one"); exists {
		t.Errorf("Expected no value for key 'one', but got %v", value)
	}
}
