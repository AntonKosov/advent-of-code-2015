package aoc_test

import (
	"reflect"
	"testing"

	"github.com/AntonKosov/advent-of-code-2015/aoc"
)

func TestPremutation(t *testing.T) {
	values := []int{0, 1, 2}
	p := aoc.HeapPermutation(values)
	if len(p) != 6 {
		t.Fatalf("Expected length: 6, actual: %v", len(p))
	}
	expectedValues := [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 1, 0},
		{2, 0, 1},
	}
nextPermutation:
	for _, ep := range expectedValues {
		for _, permutation := range p {
			if reflect.DeepEqual(ep, permutation) {
				continue nextPermutation
			}
		}
		t.Fatalf("Permutation not found: %v", ep)
	}
}
