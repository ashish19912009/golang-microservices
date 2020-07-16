package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization:
	els := []int{9, 8, 7, 6, 5}

	// execution:
	BubbleSort(els)

	//Validation:
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubbleSortBestCase(t *testing.T) {
	// Initialization:
	els := []int{5, 6, 7, 8, 9}

	// execution:
	BubbleSort(els)

	//Validation:
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubbleSortNilSplice(t *testing.T) {
	// execution:
	BubbleSort(nil)
}

func GetElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := GetElements(5)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
}

func BenchmarkBubbleSort(b *testing.B) {
	els := GetElements(20)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
func BenchmarkBubbleSort10000(b *testing.B) {
	els := GetElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSortBuildIn10000(b *testing.B) {
	els := GetElements(10000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
	// BenchmarkBubbleSortBuildIn10000-4   	    1735	    678586 ns/op	      79 B/op	       1
}
