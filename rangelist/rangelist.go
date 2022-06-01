package rangelist

import (
	. "cezhang/rangelist/rangelist/core"
	"cezhang/rangelist/rangelist/redblack"
	"fmt"
	"strings"
)

var print = fmt.Println

// internal is an abstract algorithm interface for RangeList
type internal interface {

	// Put inserts new range to underlying implementation
	Put(Range) error

	// Delete deletes start of range as key from underlying implementation
	Delete(int) error

	// Union returns a new range union with existing range(s)
	Union(Range) (Range, error)

	// Diff returns a new range difference with existing range(s)
	Diff(Range) ([]Range, error)

	// Bound returns list of start of range within a given range
	Bound(Range) ([]int, error)

	// List returns all range with ascending order
	List() ([]Range, error)
}

// RangeList struct with an underlying algorithm
type RangeList struct {
	internal
}

// New returns RangeList pointer with default RedBlack underlying implementation
func New() *RangeList {
	return &RangeList{
		redblack.New(),
	}
}

// NewWith return RangeList pointer with a given underlying implementation
func NewWith(algo internal) *RangeList {
	return &RangeList{
		algo,
	}
}

// Add inserts a new range or merges with existing range(s) with given range
// It processes with 3 steps:
// (1) union: if the given range does not intersect with existing range(s), that's a new range, else union with existing
//            range(s) and return the union one
// (2) insert: insert the unions range as a new one into underlying implementation
// (3) delete: delete those ranges that within the union range
func (rangeList *RangeList) Add(rangeElement [2]int) error {
	r, err := rangeList.checkRange(rangeElement)
	if err != nil {
		return err
	}

	if union, err := rangeList.Union(r); err != nil {
		return err
	} else {
		if err := rangeList.Put(union); err != nil {
			return err
		}
	}

	r.Start += 1 // delete keys within(not include start) the given range
	if keys2Delete, err := rangeList.Bound(r); err != nil {
		return err
	} else {
		for _, k := range keys2Delete {
			if err := rangeList.Delete(k); err != nil {
				return err
			}
		}
	}

	return nil
}

// Remove update existing range(s) by delete or cut by the given range
// It processes with 3 steps:
// (1) diff: if the given range does not intersect with existing range(s), just delete it, else cut the existing
//            range(s) with given range and return list of result ranges
// (2) insert: insert ranges from above step as new ranges into underlying implementation
// (3) delete: delete those ranges that within the given range
func (rangeList *RangeList) Remove(rangeElement [2]int) error {
	r, err := rangeList.checkRange(rangeElement)
	if err != nil {
		return err
	}

	if range2Update, err := rangeList.Diff(r); err != nil {
		return err
	} else {
		for _, r2u := range range2Update {
			rangeList.Put(r2u)
		}
	}

	r.End -= 1 // delete keys within(not include end) the given range
	if keys2Delete, err := rangeList.Bound(r); err != nil {
		return err
	} else {
		for _, k := range keys2Delete {
			if err := rangeList.Delete(k); err != nil {
				return err
			}
		}
	}

	return nil
}

// Print prints all the ranges in ascending order with format [xx, xx)
func (rangeList *RangeList) Print() error {

	if ls, err := rangeList.List(); err != nil {
		return err
	} else {
		s := make([]string, len(ls))
		for i, r := range ls {
			s[i] = fmt.Sprintf("[%d, %d)", r.Start, r.End)
		}
		print(strings.Join(s, " "))
	}
	return nil
}

// checkRange validates the array argument to make sure 0-index int less than or equal to 1-index int
func (rangeList *RangeList) checkRange(rangeElement [2]int) (Range, error) {
	if rangeElement[0] > rangeElement[1] {
		return Range{}, fmt.Errorf("invalid range")
	}
	return Range{Start: rangeElement[0], End: rangeElement[1]}, nil
}
