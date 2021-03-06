package quick

import (
	sort2 "godev/basic/algorithm/sort"
	"godev/utils"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	a, b := make(sort2.IntSlice, 100), make(sort2.IntSlice, 100)
	v := 0
	for i := range a {
		v = utils.GenerateRandomInt()
		a[i], b[i] = v, v
	}
	sort.Sort(a)
	Sort(b)

	if !sort2.Equal(a, b) {
		t.Fatal(a, b)
	}
}

// BenchmarkSort-8   	10000000	       225 ns/op
func BenchmarkSort(b *testing.B) {
	data := make(sort2.IntSlice, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = utils.GenerateRandomInt()
	}

	b.ResetTimer()

	Sort(data)
}

// BenchmarkOfficialSort-8   	10000000	       206 ns/op
func BenchmarkOfficialSort(b *testing.B) {
	data := make(sort2.IntSlice, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = utils.GenerateRandomInt()
	}

	b.ResetTimer()

	sort.Sort(data)
}
