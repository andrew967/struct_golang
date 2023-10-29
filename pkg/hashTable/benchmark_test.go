package hashTable

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func BenchmarkHashMap(b *testing.B) {
	b.Run(fmt.Sprint("Device characteristics"), func(b *testing.B) {})
	fmt.Println("------------------------")
	for i := 9; i < 19; i++ {
		size := int(math.Pow(2, float64(i)))
		fmt.Printf("Running benchmarks for size %d\n", size)

		b.Run(fmt.Sprintf("Get_%d", size), func(b *testing.B) { BenchHashMap_Get(b, size) })
		b.Run(fmt.Sprintf("Add_%d", size), func(b *testing.B) { BenchHashMap_Add(b, size) })
		b.Run(fmt.Sprintf("Remove_%d", size), func(b *testing.B) { BenchHashMap_Remove(b, size) })

		fmt.Println("------------------------")
	}
}

func BenchHashMap_Add(b *testing.B, size int) {
	hm := NewHashMap(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i)
		hm.Add(key, i)
	}
}

func BenchHashMap_Get(b *testing.B, size int) {
	hm := NewHashMap(size)
	for i := 0; i < size; i++ {
		key := "key" + strconv.Itoa(i)
		hm.Add(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i%size)
		_, _ = hm.Get(key)
	}
}

func BenchHashMap_Remove(b *testing.B, size int) {
	hm := NewHashMap(size)
	for i := 0; i < size; i++ {
		key := "key" + strconv.Itoa(i)
		hm.Add(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i%size)
		_ = hm.Remove(key)
	}
}

func BenchmarkHashSet(b *testing.B) {
	b.Run(fmt.Sprint("Device characteristics"), func(b *testing.B) {})
	fmt.Println("------------------------")
	for i := 9; i < 19; i++ {
		size := int(math.Pow(2, float64(i)))
		fmt.Printf("Running benchmarks for size %d\n", size)

		b.Run(fmt.Sprintf("Contains_%d", size), func(b *testing.B) { BenchHashSet_Contains(b, size) })
		b.Run(fmt.Sprintf("Add_%d", size), func(b *testing.B) { BenchHashSet_Add(b, size) })
		b.Run(fmt.Sprintf("Remove_%d", size), func(b *testing.B) { BenchHashSet_Remove(b, size) })

		fmt.Println("------------------------")
	}
}

func BenchHashSet_Add(b *testing.B, size int) {
	hs := NewHashSet(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		hs.Add(key)
	}
}

func BenchHashSet_Contains(b *testing.B, size int) {
	hs := NewHashSet(size)
	for i := 0; i < size; i++ {
		key := fmt.Sprintf("key%d", i)
		hs.Add(key)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i%size)
		hs.Contains(key)
	}

}

func BenchHashSet_Remove(b *testing.B, size int) {
	hs := NewHashSet(size)
	for i := 0; i < size; i++ {
		key := fmt.Sprintf("key%d", i)
		hs.Add(key)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", int(i%size))
		_ = hs.Remove(key)
	}
}
