package day18

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(18, 1, "magnitude of the final sum", firstStar)
}

func SecondStar() {
	utils.Star(18, 2, "largest magnitude of any sum of two different numbers", secondStar)
}

func firstStar(content string) int {
	lines := strings.Split(content, "\n")
	sum := extractSnailfishes(lines[0])
	for _, l := range lines[1:] {
		sum = sum.add(extractSnailfishes(l))
		sum.reduce()
	} 
	return sum.magnitude()
}

func secondStar(content string) int {
	lines := strings.Split(content, "\n")
	max := 0
	for i, li := range lines {
		for j, lj := range lines {
			if i != j {
				si := extractSnailfishes(li)
				sj := extractSnailfishes(lj)
				sum := si.add(sj)
				sum.reduce()
				mag := sum.magnitude()
				if max < mag { max = mag }
			}
		}
	}
	return max
}

type node struct {
	parent *node
	value  int
	left   *node
	right  *node
}

func extractSnailfishes(content string) *node {
	root := node { nil, -1, nil, nil }
	root.navigate(content, 0)
	return &root
}

func (n *node) navigate(content string, i int) int {
	for i < len(content) {
		if content[i] == '[' {
			c := node { n, -1, nil, nil }
			n.left = &c
			i = c.navigate(content, i+1)
		} else if content[i] >= '0' && content[i] <= '9' {
			j := i+1
			for content[j] >= '0' && content[j] <= '9' { j++ }
			n.value = utils.StringToInt(content[i:j])
			i = j
			break
		} else if content[i] == ',' {
			c := node { n, -1, nil, nil }
			n.right = &c
			i = c.navigate(content, i+1)
		} else if content[i] == ']' {
			i++
			break 
		}
	}
	return i
}

func (n *node) reduce() {
	for n.explode(0) || n.split() {}
}

func (n *node) explode(level int) bool {
	if level == 4 && n.value == -1 {
		var p *node
		// explode left
		for p = n; p.parent != nil && p.parent.left == p; p = p.parent {}
		if p.parent != nil {
			p = p.parent.left
			for p.value == -1 { p = p.right }
			p.value += n.left.value
		}
		// explode right
		for p = n; p.parent != nil && p.parent.right == p; p = p.parent {}
		if p.parent != nil {
			p = p.parent.right
			for p.value == -1 { p = p.left }
			p.value += n.right.value
		}
		// replace node with 0
		n.value = 0
		n.left = nil
		n.right = nil
		return true
	} else if n.value == -1 {
		return n.left.explode(level + 1) || n.right.explode(level + 1)
	} else {
		return false
	}
}

func (n *node) split() bool {
	if n.value > 9 {
		vleft := n.value / 2
		vright := n.value - vleft
		n.value = -1
		n.left = &node { n, vleft, nil, nil }
		n.right = &node { n, vright, nil, nil }
		return true
	} else if n.value == -1 {
		return n.left.split() || n.right.split()
	} else {
		return false
	}
}

func (left *node) add(right *node) *node {
	root := node{ nil, -1, left, right }
	left.parent = &root
	right.parent = &root
	return &root
}

func (n *node) magnitude() int {
	if n == nil {
		return 0
	} else if n.value == -1 {
		return 3 * n.left.magnitude() + 2 * n.right.magnitude()
	} else {
		return n.value
	}
}

func (n *node) print() string {
	if n.value >= 0 {
		return fmt.Sprintf("%d", n.value)
	} else {
		return fmt.Sprintf("[%s,%s]", n.left.print(), n.right.print())
	}
}