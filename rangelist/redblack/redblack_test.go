package redblack

import (
	. "cezhang/rangelist/rangelist/core"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedBlack(t *testing.T) {

	rb := New()
	data := Range{Start: 1, End: 5}
	err := rb.Put(data)
	assert.Nilf(t, err, "Put should succeed")

	unionRes, err := rb.Union(data)
	assert.Equal(t, data, unionRes, fmt.Sprintf("Union result Should be the same with %v", data))
	assert.Nilf(t, err, "Union should succeed")

	boundRes, err := rb.Bound(data)
	boundExp := []int{1}
	assert.Equal(t, boundExp, boundRes, fmt.Sprintf("Bound result Should be the same with %v", boundExp))
	assert.Nilf(t, err, "Bound should succeed")

	diffRes, err := rb.Diff(data)
	diffExp := []Range{
		{5, 5},
	}
	assert.Equal(t, diffExp, diffRes, fmt.Sprintf("Diff result Should be the same with %v", diffExp))
	assert.Nilf(t, err, "Diff should succeed")

	err = rb.Delete(data.Start)
	assert.Nilf(t, err, "Delete should succeed")

	listRes, err := rb.List()
	listExp := []Range{}
	assert.Equal(t, listExp, listRes, fmt.Sprintf("List result Should be the same with %v", diffExp))
	assert.Nilf(t, err, "List should succeed")

}

func TestRedBlackPut(t *testing.T) {
	putTestcases := map[string]struct {
		input []int
		err   error
	}{
		"common":    {[]int{1, 5}, nil},
		"intersect": {[]int{2, 3}, nil},
		"disjoint":  {[]int{6, 7}, nil},
	}

	putRB := New()
	for name, tc := range putTestcases {
		t.Run(name, func(t *testing.T) {
			actual := putRB.Put(Range{tc.input[0], tc.input[1]})
			assert.Nilf(t, actual, "Put should succeed")
		})
	}
}
