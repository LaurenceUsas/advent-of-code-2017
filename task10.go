package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

//Task10 Solution
func Task10() {
	pwd, _ := os.Getwd()
	input := helpers.InputFile(pwd + "/input/input10.txt")[0]
	file := strings.Split(input, ",")

	var instructions []int
	for _, v := range file {
		i, _ := strconv.Atoi(v)
		instructions = append(instructions, i)
	}

	fmt.Printf("[Part 1 answer] - %v\n", task10PartOne(instructions))
	fmt.Printf("[Part 2 answer] - %v\n", task10PartTwo(input))

}

func task10PartOne(input []int) int { // temp

	kh := NewKnotHash()
	for _, cmd := range input {
		kh.reverseSequence(cmd)
	}
	return kh.Sequence[0] * kh.Sequence[1]
}

func task10PartTwo(input string) string { // temp
	kh := NewKnotHash()
	return kh.Hash(input)
}

type KnotHash struct {
	Sequence        [256]int
	CurrentPosition int
	Skip            int
}

func NewKnotHash() *KnotHash {
	kh := &KnotHash{}
	var a [256]int
	for i := 0; i < 256; i++ {
		a[i] = i
	}
	kh.Sequence = a
	kh.CurrentPosition = 0
	kh.Skip = 0

	return kh
}

func (kh *KnotHash) reverseSequence(l int) {
	size := 256
	//Reversing.
	id0 := kh.CurrentPosition   // Start ID
	id1 := (id0 + l - 1) % size // End ID
	for i := 0; i < l/2; i++ {
		t1 := kh.Sequence[id0]
		t2 := kh.Sequence[id1]
		kh.Sequence[id0] = t2
		kh.Sequence[id1] = t1

		id0 = (id0 + 1) % size
		id1 = (id1 + 255) % size // To avoid if when id == 0 and id--
	}

	kh.CurrentPosition = (kh.CurrentPosition + l + kh.Skip) % size
	kh.Skip++
}

func (kh *KnotHash) Hash(input string) string {
	// Clear old junk.
	var a [256]int
	for i := 0; i < 256; i++ {
		a[i] = i
	}
	kh.Sequence = a
	kh.CurrentPosition = 0
	kh.Skip = 0

	// Generate instructions
	var instructions []int
	for _, char := range input {
		instructions = append(instructions, int(char))
	}
	instructions = append(instructions, 17, 31, 73, 47, 23)

	// Do 64 Rounds
	for i := 0; i < 64; i++ {
		for _, cmd := range instructions {
			kh.reverseSequence(cmd)
		}
	}

	sh := kh.Sequence // Sparse Hash
	dh := []byte{}    // Dense Hash

	bs := 0 // bit sum

	for i := 0; i < len(sh); i++ {
		bs ^= sh[i]
		if (i+1)%16 == 0 {
			dh = append(dh, byte(bs))
			bs = 0
		}
	}
	return hex.EncodeToString(dh)
}
