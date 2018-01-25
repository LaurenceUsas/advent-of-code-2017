package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

//Task16 Solution
func Task16() {
	//read input.
	pwd, _ := os.Getwd()
	file := helpers.InputFile(pwd + "/input/input16.txt")
	instructions := strings.Split(file[0], ",")

	fmt.Printf("[Part 1 answer] - %v\n", task16PartOne(instructions))
	fmt.Printf("[Part 2 answer] - %v\n", task16PartTwo(instructions))
}

type PermutationNode struct {
	Next  *PermutationNode
	Value byte
}

type Permutator struct {
	Head          *PermutationNode
	Tail          *PermutationNode
	HashsetValues map[byte]*PermutationNode
}

func NewPermutator(input string) *Permutator {
	p := new(Permutator)

	p.HashsetValues = make(map[byte]*PermutationNode)

	var previous *PermutationNode
	for _, char := range input {
		v := byte(char)
		node := new(PermutationNode)
		node.Value = v
		p.HashsetValues[v] = node

		if previous != nil {
			previous.Next = node
		}
		p.HashsetValues[v] = node

		if p.Head == nil {
			p.Head = node
		} else {
			p.Tail.Next = node
		}
		p.Tail = node
		previous = node
	}
	return p
}

// 0(n)
func (p *Permutator) Offset(input int) {
	a := p.Head
	b := p.Head
	for i := 1; i < input; i++ {
		b = b.Next
	}

	p.Head = b.Next

	last := p.Tail
	last.Next = a
	p.Tail = b
	b.Next = nil
}

// 0(1)
// Because data is small its easier to swap data.
func (p *Permutator) SwapByID(a, b int) {
	var an *PermutationNode
	var bn *PermutationNode

	n := p.Head
	for i := 0; i <= maxInt(a, b); i++ {
		if i == a {
			an = n
		} else if i == b {
			bn = n
		}
		n = n.Next
	}

	swapNodeValues(an, bn, p)
}

func maxInt(a, b int) int {
	max := a
	if a < b {
		max = b
	}
	return max
}

// O(1)
func (p *Permutator) SwapByValue(a, b byte) {
	an := p.HashsetValues[a]
	bn := p.HashsetValues[b]

	swapNodeValues(an, bn, p)
}

func swapNodeValues(an, bn *PermutationNode, p *Permutator) {
	if an == nil || bn == nil {
		return
	}
	var t byte
	t = an.Value
	an.Value = bn.Value
	bn.Value = t

	p.HashsetValues[an.Value] = an
	p.HashsetValues[bn.Value] = bn
}

func (p *Permutator) Value() string {
	out := ""
	n := p.Head
	var ba []byte
	for n != nil {
		ba = append(ba, n.Value)
		n = n.Next
	}
	out = string(ba)
	return out
}

/*
s1, a spin of size 1: eabcd.
x3/4, swapping the last two programs: eabdc.
pe/b, swapping programs e and b: baedc.
*/

func task16PartOne(instructions []string) string {

	p := NewPermutator("abcdefghijklmnop")
	for _, command := range instructions {
		fmt.Println(p.Value())
		fmt.Println(command)

		switch {
		case strings.Contains(command, "s"):
			v, _ := strconv.Atoi(command[1:])
			p.Offset(v)
		case strings.Contains(command, "x"):
			v := strings.Split(command[1:], "/")
			a, _ := strconv.Atoi(v[0])
			b, _ := strconv.Atoi(v[1])
			p.SwapByID(a, b)
		case strings.Contains(command, "p"):
			v := strings.Split(command[1:], "/")
			a, _ := utf8.DecodeRuneInString(v[0])
			b, _ := utf8.DecodeRuneInString(v[1])
			p.SwapByValue(byte(a), byte(b))
		}
	}

	result := p.Value()
	return result
}

func task16PartTwo(instructions []string) string {
	result := ""
	return result
}
