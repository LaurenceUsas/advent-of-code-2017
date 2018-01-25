package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

//Task18 Solution
func Task18() {
	pwd, _ := os.Getwd()
	instructions := helpers.InputFile(pwd + "/input/input18.txt")
	//fmt.Printf("[Part 1 answer] - %v\n", task18PartOne(instructions))
	fmt.Printf("[Part 2 answer] - %v\n", task18PartTwo(instructions))
}

type Instruction struct {
	Command  string //command
	Register string //register
	Value    string //value
}

type Duet struct {
	Instructions map[int]*Instruction
	RegMemory    map[string]int // Register memory
	RegDefault   int            // Default value for empty register
	K            int            // Instruction progress
	Frequency    int            // Frequency of last sound
	SentValues   int
	IsWaiting    bool
	Finished     bool
}

func (d *Duet) GetValue(v string) int {
	// accept both letter and
	val, err := strconv.Atoi(v)
	if err != nil {
		if value, ok := d.RegMemory[v]; !ok {
			val = d.RegDefault
		} else {
			val = value
		}
	}
	return val
}

func NewDuet(instruction []string, regDefault int) *Duet {
	d := &Duet{}
	d.Instructions = make(map[int]*Instruction)
	d.RegMemory = make(map[string]int)
	for i, v := range instruction {
		inst := helpers.SplitBySpace(v)
		ti := &Instruction{
			Command:  inst[0],
			Register: inst[1],
		}
		if len(inst) == 3 {
			ti.Value = inst[2]
		}
		d.Instructions[i] = ti
	}
	d.RegDefault = regDefault
	d.K = 0
	d.SentValues = 0
	d.IsWaiting = false
	d.Finished = false
	return d
}

func (d *Duet) Run() int {
	for {
		inst := d.Instructions[d.K]
		switch inst.Command {
		case "snd":
			d.Frequency = d.RegMemory[inst.Register]
		case "set":
			val := d.GetValue(inst.Value)
			d.RegMemory[inst.Register] = val
		case "add":
			//if Int add value or String - get int from register.
			add := d.GetValue(inst.Value)
			d.RegMemory[inst.Register] += add
		case "mul":
			val := d.GetValue(inst.Value)
			d.RegMemory[inst.Register] *= val
		case "mod":
			val := d.GetValue(inst.Value)
			d.RegMemory[inst.Register] %= val
		case "rcv":
			val := d.GetValue(inst.Register)
			if val != 0 {
				return d.Frequency
			}
		case "jgz":
			val := d.GetValue(inst.Register)
			if val != 0 {
				add, _ := strconv.Atoi(inst.Value)
				d.K += add
				d.K--
			}
		}
		d.K++
	}
}

func task18PartOne(instructions []string) int {
	d := NewDuet(instructions, 0)
	return d.Run()
}

func task18PartTwo(instructions []string) int {
	d0 := NewDuet(instructions, 0)
	d1 := NewDuet(instructions, 1)

	in0 := make(chan int, 1000)
	in1 := make(chan int, 1000)

	// w1 := false
	// w2 := false

	go d0.RunSync(in1, in0)
	go d1.RunSync(in0, in1)

	// all cases except. d1 and d2 is waiting.
	for {
		if d1.Finished {
			return d1.SentValues
		} else if d0.IsWaiting && d1.IsWaiting {
			fmt.Printf("Both Waiting [%v][%v]\n", len(in0), len(in1))
			if len(in0) == 0 || len(in1) == 0 {
				return d1.SentValues
			}
		}
	}
	// Deadlock when both trying to receive from the empty channel.
	// When both waiting.
}

func (d *Duet) RunSync(out chan<- int, in chan int) {
	for {
		if d.K >= len(d.Instructions) || d.K < 0 {
			fmt.Println("Finished!")
			d.Finished = true
			break
		}
		inst := d.Instructions[d.K]
		switch inst.Command {
		case "snd":
			d.SentValues++
			out <- d.RegMemory[inst.Register]
		case "set":
			val := d.GetValue(inst.Value)
			d.RegMemory[inst.Register] = val
		case "add":
			//if Int add value or String - get int from register.
			add := d.GetValue(inst.Value)
			d.RegMemory[inst.Register] += add
		case "mul":
			val := d.GetValue(inst.Value)
			d.RegMemory[inst.Register] *= val
		case "mod":
			val := d.GetValue(inst.Value)
			d.RegMemory[inst.Register] %= val
		case "rcv":
			if len(in) == 0 {
				d.IsWaiting = true
			}
			val := <-in
			d.IsWaiting = false
			d.RegMemory[inst.Register] = val
		case "jgz":
			val := d.GetValue(inst.Register)
			if val != 0 {
				add, _ := strconv.Atoi(inst.Value)
				d.K += add
				d.K--
			}
		}
		d.K++
	}
}
