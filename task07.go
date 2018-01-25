package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

type node struct {
	name      string
	weight    int
	sumWeight int
	parent    *node
	children  []*node
}

//Task07 Solution
func Task07() {
	//Read input.
	lines := helpers.InputFile("/input/input07.txt")
	nodes := map[string]*node{}

	for _, line := range lines {
		v := helpers.SplitBySpace(line)

		//if node exists - use
		//else make a new one.
		n := nodes[v[0]]
		if n == nil {
			n = &node{}
		}
		n.name = v[0]
		n.weight, _ = strconv.Atoi(v[1][1 : len(v[1])-1])

		// look for children
		if strings.Contains(line, "->") {
			for _, c := range v[3:] {
				c = strings.Trim(c, ",")
				nc := nodes[c]
				// if exist - reference on this one.
				// if not make + assign name + parent.
				if nc == nil {
					nc = &node{}
					nc.name = c
				}
				nc.parent = n // Why not &n
				n.children = append(n.children, nc)
				nodes[nc.name] = nc
			}
		}
		nodes[n.name] = n
	}

	var root *node
	//Find node without a parent.
	for _, v := range nodes {
		if v.parent == nil {
			root = v
			fmt.Printf("Part 1 answer: %s", root.name)
		}
	}

	calculateNodeWeights(nodes, root)
	findBadNode(root)
}

// Node sumweights.
func calculateNodeWeights(nodes map[string]*node, root *node) int {
	childrenWeight := 0
	for _, c := range root.children {
		childrenWeight += calculateNodeWeights(nodes, c)
	}
	root.sumWeight = root.weight + childrenWeight
	return root.weight + childrenWeight
}

func findBadNode(root *node) {

	cw := map[int]int{}

	for _, c := range root.children {
		cw[c.sumWeight]++
		//Could be optimised here
	}

	for k, v := range cw {
		if v == 1 {
			//Found weight
			for _, vv := range root.children {
				if k == vv.sumWeight {
					findBadNode(vv)
					return
				}
			}
		}
	}

	for _, v := range root.parent.children {
		fmt.Println(v.sumWeight)
	}
}
