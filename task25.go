package main

import (
	"fmt"
)

//Task25 Solution
func Task25() {
	fmt.Printf("[Part 1 answer] - %v\n", task25PartOne("A", 12399302))
	fmt.Printf("[Part 2 answer] - %v\n", 0)
}

type TurnerNode struct {
	v int
	l *TurnerNode
	r *TurnerNode
}

func (tn *TurnerNode) Right() *TurnerNode {
	if tn.r == nil {
		tn.r = &TurnerNode{l: tn}
	}
	return tn.r
}

func (tn *TurnerNode) Left() *TurnerNode {
	if tn.l == nil {
		tn.l = &TurnerNode{r: tn}
	}
	return tn.l
}

type TurnerInstructions struct {
	Instruction map[string]*TurnerCommand
}

type TurnerCommand struct {
	WriteValue [2]int
	NextMove   [2]int
	NextState  [2]string
}

func task25PartOne(state string, steps int) int {
	n := &TurnerNode{
		v: 0,
		l: nil,
		r: nil,
	}

	ti := NewInstructions()

	for i := 0; i < steps; i++ {
		inst := ti.Instruction[state] //Get instruction.
		cv := n.v                     //Value

		n.v = inst.WriteValue[cv]  // Assign Value
		switch inst.NextMove[cv] { // Move
		case -1:
			n = n.Left()
		case 1:
			n = n.Right()
		}
		state = inst.NextState[cv] // New state
	}

	// Count all. Fun with channels :)
	cl := make(chan int)
	cr := make(chan int)

	go func() {
		count := 0
		nn := n
		for nn.l != nil {
			nn = nn.l
			count += nn.v
		}
		cl <- count
	}()

	go func() {
		count := 0
		nn := n
		for nn.r != nil {
			nn = nn.r
			count += nn.v
		}
		cr <- count
	}()

	return n.v + <-cl + <-cr
}

/*
Begin in state A.
Perform a diagnostic checksum after 12399302 steps.
*/
func NewInstructions() *TurnerInstructions {
	ti := &TurnerInstructions{}
	c := make(map[string]*TurnerCommand)
	c["A"] = &TurnerCommand{
		WriteValue: [2]int{1, 0},
		NextMove:   [2]int{1, 1},
		NextState:  [2]string{"B", "C"},
	}
	c["B"] = &TurnerCommand{
		WriteValue: [2]int{0, 0},
		NextMove:   [2]int{-1, 1},
		NextState:  [2]string{"A", "D"},
	}
	c["C"] = &TurnerCommand{
		WriteValue: [2]int{1, 1},
		NextMove:   [2]int{1, 1},
		NextState:  [2]string{"D", "A"},
	}
	c["D"] = &TurnerCommand{
		WriteValue: [2]int{1, 0},
		NextMove:   [2]int{-1, -1},
		NextState:  [2]string{"E", "D"},
	}
	c["E"] = &TurnerCommand{
		WriteValue: [2]int{1, 1},
		NextMove:   [2]int{1, -1},
		NextState:  [2]string{"F", "B"},
	}
	c["F"] = &TurnerCommand{
		WriteValue: [2]int{1, 1},
		NextMove:   [2]int{1, 1},
		NextState:  [2]string{"A", "E"},
	}
	ti.Instruction = c
	return ti
}
