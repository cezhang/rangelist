package redblack

import (
	. "cezhang/rangelist/rangelist/core"
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

/*
	RedBlack provides underlying implementation of RangeList based on red-black tree.
	Since inputs are only int, all the type conversion or assertion assumes to be correct without further check.
*/
type RedBlack struct {
	rbt *rbt.Tree
}

// New returns a RedBlack pointer
func New() *RedBlack {
	return &RedBlack{
		rbt: rbt.NewWithIntComparator(),
	}
}

// Put inserts range into reb-black tree that start of range as key and end of end as value
func (rb *RedBlack) Put(r Range) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("something wrong with Put method : %s", r)
		}
	}()
	rb.rbt.Put(r.Start, r.End)
	return nil
}

// Union returns the union result between given range and  existing range(s)
func (rb *RedBlack) Union(r Range) (o Range, err error) {
	defer func() {
		if r := recover(); r != nil {
			o = Range{}
			err = fmt.Errorf("something wrong with Union method : %s", r)
		}
	}()

	unionStart := r.Start
	if floorStartNode, found := rb.rbt.Floor(r.Start); found {
		e := floorStartNode.Value.(int)
		//     |_____________|  <--- e
		//  r.Start ---> |_______________|
		if e >= r.Start {
			unionStart = floorStartNode.Key.(int)
		} else {
			//     |_____________|  <--- e
			//  		r.Start ---> |_______________|
			// do nothing here
		}
	}

	unionEnd := r.End
	if floorEndNode, found := rb.rbt.Floor(r.End); found {
		if floorEnd, ok := floorEndNode.Value.(int); ok {
			//     						|___________|  <--- floorEnd
			//  r.Start ---> |_______________|
			if floorEnd >= r.End {
				unionEnd = floorEnd
			} else {
				//     					 |____|  <--- floorEnd
				//  r.Start ---> |_______________|
				// do nothing here
			}
		}
	}

	return Range{Start: unionStart, End: unionEnd}, nil
}

// Bound returns lists of start of existing range(s) within the given range
//    |_______|                yes
//      |_______|              yes
//                 |_______|   yes
// 				    |_______|  no
//    |___________|     <----  input
func (rb *RedBlack) Bound(r Range) ([]int, error) {
	out := make([]int, 0)
	if floorStartNode, found := rb.rbt.Ceiling(r.Start); found {
		it := rb.rbt.IteratorAt(floorStartNode)
		for {
			key := it.Key().(int)
			if key > r.End {
				return out, nil
			}
			out = append(out, key)
			if !it.Next() {
				break
			}
		}
	}

	return out, nil
}

// Delete deletes nodes with a given key from reb-black tree
func (rb *RedBlack) Delete(key int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("something wrong with Remove method : %s", r)
		}
	}()
	rb.rbt.Remove(key)
	return nil
}

// List returns lists of ranges sorting by start  in ascending order
func (rb *RedBlack) List() (rs []Range, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("something wrong with All method : %s", r)
		}
	}()

	out := make([]Range, 0)
	it := rb.rbt.Iterator()
	for it.Next() {
		out = append(out, Range{Start: it.Key().(int), End: it.Value().(int)})
	}
	return out, nil
}

// Diff returns the difference(or subtraction) result between given range and  existing range(s)
func (rb *RedBlack) Diff(r Range) ([]Range, error) {

	diff := make([]Range, 0)
	if floorStartNode, found := rb.rbt.Floor(r.Start - 1); found {
		e := floorStartNode.Value.(int)
		//   |_______|  <--- e						|_____|      <-- need to insert
		//                          --- diff ---> 		  |__|   <-- to remove
		//        |___________|							 |___________|
		if e >= r.Start {
			diff = append(diff, Range{Start: floorStartNode.Key.(int), End: r.Start})
		}

	}

	if floorEndNode, found := rb.rbt.Floor(r.End - 1); found {
		floorEnd := floorEndNode.Value.(int)
		//   				|_______| <--- floorEnd								      |_____|  <-- need to insert
		//                                				 --- diff ---> 		  	  |__|         <-- to remove
		//        |___________|							   				 |___________|
		if floorEnd >= r.End {
			diff = append(diff, Range{Start: r.End, End: floorEndNode.Value.(int)})
		}
	}

	return diff, nil
}
