package main

import (
	"fmt"
	"os"
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
	pwd, _ := os.Getwd()
	lines := helpers.InputFile(pwd + "/input/input07.txt")
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
			fmt.Printf("Part 1 answer: %s\n", root.name)
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
	l := len(root.children)
	switch {
	case l == 0:
		return
	case l > 2:
		// Comparer
		var a, b *node
		ag := false // A good
		for _, c := range root.children {
			if a == nil {
				a = c
			} else if b != nil && c.sumWeight == a.sumWeight {
				// B Bad
				findBadNode(b)
				return
			} else if b != nil && c.sumWeight == b.sumWeight {
				// A Bad
				findBadNode(a)
				return
			} else if c.sumWeight != a.sumWeight {
				// New B
				if ag {
					findBadNode(c)
					return
				}
				b = c

			} else if c.sumWeight == a.sumWeight {
				ag = true
			}
		}
		diff := root.sumWeight - (root.parent.sumWeight-root.parent.weight-root.sumWeight)/(len(root.parent.children)-1)
		sw := root.weight - diff

		fmt.Printf("Part 1 answer: %v\n", sw)
		return
	}
}
