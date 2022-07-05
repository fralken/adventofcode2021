package day12

import (
	"aoc2021/utils"
	"fmt"
	"strings"
	"unicode"
)

func FirstStar() {
	content := utils.ReadFile("./input/day12.txt")
	value := firstStar(content)
	fmt.Printf("day 12.1 - number of paths: %d\n", value)
}

func SecondStar() {
	content := utils.ReadFile("./input/day12.txt")
	value := secondStar(content)
	fmt.Printf("day 12.2 - number of paths: %d\n", value)
}

func firstStar(content string) int {
	graph := makeGraph(content)
	var path []string
	path = append(path, "start")
	nodes := graph["start"]
	count := 0
	index := 0
	for len(path) > 0 {
		if index >= len(nodes) {
			l := last(path)
			path = path[:len(path)-1]
			if len(path) > 0 {
				nodes = graph[last(path)]
				index = find(nodes, l)
			}
		} else if nodes[index] == "end" {
			count++
		} else if isBig(nodes[index]) || !contains(path, nodes[index]) {
			path = append(path, nodes[index])
			nodes = graph[last(path)]
			index = -1
		}
		index++
	}
	return count
}

func secondStar(content string) int {
	graph := makeGraph(content)
	var path []string
	path = append(path, "start")
	nodes := graph["start"]
	count := 0
	index := 0
	double := ""
	for len(path) > 0 {
		if index >= len(nodes) {
			l := last(path)
			if l == double {
				double = ""
			}
			path = path[:len(path)-1]
			if len(path) > 0 {
				nodes = graph[last(path)]
				index = find(nodes, l)
			}
		} else if nodes[index] == "end" {
			count++
		} else if double == "" || isBig(nodes[index]) || !contains(path, nodes[index]) {
			if isSmall(nodes[index]) && contains(path, nodes[index]) {
				double = nodes[index]
			}
			path = append(path, nodes[index])
			nodes = graph[last(path)]
			index = -1
		}
		index++
	}
	return count
}

type nodes map[string][]string

func makeGraph(content string) nodes {
	lines := strings.Split(content, "\n")
	graph := make(nodes)
	for _, line := range lines {
		nodes := strings.Split(line, "-")
		src := nodes[0]
		dst := nodes[1]
		if src == "start" || dst == "end" {
			graph[src] = append(graph[src], dst)
		} else if dst == "start" || src == "end" {
			graph[dst] = append(graph[dst], src)
		} else {
			graph[src] = append(graph[src], dst)
			graph[dst] = append(graph[dst], src)
		}
	}
	return graph
}

func last(s []string) string {
	return s[len(s)-1]
}

func isBig(s string) bool {
	for _, v := range s {
		if unicode.IsUpper(v) {
			return true
		}
	}
	return false
}

func isSmall(s string) bool {
	for _, v := range s {
		if unicode.IsLower(v) {
			return true
		}
	}
	return false
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func find(slice []string, s string) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}
	panic("node not found")
}
