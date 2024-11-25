package utils

import (
	"hash/fnv"
)

type Pair struct {
	Key   string
	Value int
	Next  *Pair // Добавляем поле для следующего элемента в связном списке
}

type Bucket struct {
	Head *Pair
}

type HashMap struct {
	buckets []Bucket
	size    int
}

func NewHashMap(initialCapacity int) *HashMap {
	buckets := make([]Bucket, initialCapacity)
	for i := range buckets {
		buckets[i] = Bucket{}
	}
	return &HashMap{buckets: buckets, size: 0}
}

func (h *HashMap) hash(key string) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int(hash.Sum32()) % len(h.buckets)
}

func (h *HashMap) Set(key string, value int) {
	index := h.hash(key)
	bucket := &h.buckets[index]

	for p := bucket.Head; p != nil; p = p.Next {
		if p.Key == key {
			p.Value = value
			return
		}
	}

	newPair := &Pair{Key: key, Value: value, Next: bucket.Head}
	bucket.Head = newPair
	h.size++
}

func (h *HashMap) Get(key string) (int, bool) {
	index := h.hash(key)
	bucket := &h.buckets[index]

	for p := bucket.Head; p != nil; p = p.Next {
		if p.Key == key {
			return p.Value, true
		}
	}

	return 0, false
}

func (h *HashMap) Remove(key string) bool {
	index := h.hash(key)
	bucket := &h.buckets[index]

	prev := &bucket.Head
	for p := bucket.Head; p != nil; p = p.Next {
		if p.Key == key {
			*prev = p.Next
			h.size--
			return true
		}
		prev = &p.Next
	}

	return false
}

func (h *HashMap) Size() int {
	return h.size
}

// func main() {
// 	hashMap := NewHashMap(32)
// 	hashMap.Set("one", 1)
// 	hashMap.Set("two", 2)
// 	fmt.Println(hashMap.Get("one")) // Вывод: 1, true
// 	fmt.Println(hashMap.Get("two")) // Вывод: 2, true
// 	fmt.Println(hashMap.Size())     // Вывод: 2
// 	hashMap.Remove("one")
// 	fmt.Println(hashMap.Get("one")) // Вывод: 0, false
// 	fmt.Println(hashMap.Size())     // Вывод: 1
// }
