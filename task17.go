package main

import (
	"fmt"
)

//Task17 Solution
func Task17() {
	input := 356
	fmt.Printf("[Part 1 answer] - %v\n", task17PartOne(input))
	//fmt.Printf("[Part 2 answer] - %v\n", task17PartTwo(input)) //9min
}

type Spinlock struct {
	Value int
	Next  *Spinlock
}

func NewSpinlock() *Spinlock {
	sl := &Spinlock{}
	sl.Next = sl
	sl.Value = 0
	return sl
}

func (sl *Spinlock) Insert(n, skip int) {
	for i := 1; i <= n; i++ {
		for s := 0; s < skip; s++ {
			sl = sl.Next
		}
		// Create new with next id.
		nsl := NewSpinlock()
		nsl.Value = i
		nsl.Next = sl.Next
		sl.Next = nsl
		// Step into current
		sl = nsl
	}
}

func (sl *Spinlock) ValueAfterN(n int) int {
	v := sl.Value
	for v != n {
		v = sl.Value
		sl = sl.Next
	}
	return sl.Value
}

func task17PartOne(skip int) int {
	sl := NewSpinlock()
	sl.Insert(2017, skip)
	result := sl.ValueAfterN(2017)
	return result
}

func task17PartTwo(skip int) int {
	sl := NewSpinlock()
	sl.Insert(50000000, skip)
	result := sl.ValueAfterN(0)
	return result
}
