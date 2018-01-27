package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

type Coprocessor struct {
	Instructions map[int]*Instruction
	RegMemory    map[string]int // Register memory
	RegDefault   int            // Default value for empty register
	K            int            // Instruction progress
	Counter      int            // Generic
}

func (c *Coprocessor) GetValue(v string) int {
	// accept both letter and
	val, err := strconv.Atoi(v)
	if err != nil {
		val = c.RegMemory[v]
	}
	return val
}

func NewCoprocessor(instruction []string, regDefault int) *Coprocessor {
	c := &Coprocessor{}
	c.Instructions = make(map[int]*Instruction)
	c.RegMemory = make(map[string]int)
	for i, v := range instruction {
		inst := helpers.SplitBySpace(v)
		ti := &Instruction{
			Command:  inst[0],
			Register: inst[1],
		}
		if len(inst) == 3 {
			ti.Value = inst[2]
		}
		c.Instructions[i] = ti
	}
	c.RegDefault = regDefault
	c.K = 0
	c.Counter = 0
	return c
}

func (c *Coprocessor) Run() {
	for c.K < len(c.Instructions) {
		inst := c.Instructions[c.K]
		switch inst.Command {
		case "set":
			val := c.GetValue(inst.Value)
			c.RegMemory[inst.Register] = val
		case "sub":
			val := c.GetValue(inst.Value)
			c.RegMemory[inst.Register] -= val
			//helpers.PrintTwo(inst.Register, c.RegMemory[inst.Register])
		case "mul":
			val := c.GetValue(inst.Value)
			c.RegMemory[inst.Register] *= val
			c.Counter++
		case "jnz":
			val := c.GetValue(inst.Register)
			if val != 0 {
				add, _ := strconv.Atoi(inst.Value)
				c.K += add
				c.K--
			}
		}
		c.K++
	}
}

//Task23 Solution
func Task23() {

	pwd, _ := os.Getwd()
	file := helpers.InputFile(pwd + "/input/input23.txt")
	c1 := NewCoprocessor(file, 0)
	c1.Run()

	fmt.Printf("[Part 1 answer] - %v\n", c1.Counter)
	fmt.Printf("[Part 2 answer] - %v\n", task23PartTwo())
}

func task23PartTwo() int {
	var b, d, h int
	// Count prime number in between with 17 increment.
	for b = 105700; b <= 122700; b += 17 {
		for d = 2; d <= b/2; d++ {
			if b%d == 0 {
				h++
				break
			}
		}
	}
	return h
}
