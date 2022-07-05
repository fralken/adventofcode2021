package day04

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	content := utils.ReadFile("./input/day04.txt")
	value := firstStar(content)
	fmt.Printf("day  4.1 - final score (first board): %d\n", value)
}

func SecondStar() {
	content := utils.ReadFile("./input/day04.txt")
	value := secondStar(content)
	fmt.Printf("day  4.2 - final score (last board): %d\n", value)
}

func firstStar(content string) int {
	lines := strings.Split(content, "\n")
	numbers := utils.StringsToInts(strings.Split(lines[0], ","))
	boards := fillBoards(lines[1:])
	found, number := bingo(boards, numbers)
	sum := 0
	for _, row := range boards[found] {
		for _, cell := range row {
			sum += cell
		}
	}
	return sum * number
}

func secondStar(content string) int {
	lines := strings.Split(content, "\n")
	numbers := utils.StringsToInts(strings.Split(lines[0], ","))
	boards := fillBoards(lines[1:])
	var lastNumber int
	var lastBoard [][]int
	for len(boards) > 0 {
		found, number := bingo(boards, numbers)
		if found >= 0 {
			lastBoard = boards[found]
			lastNumber = number
			boards = remove(boards, found)
		} else {
			break
		}
	}
	sum := 0
	for _, row := range lastBoard {
		for _, cell := range row {
			sum += cell
		}
	}
	return sum * lastNumber
}

func fillBoards(lines []string) (boards [][][]int) {
	var board [][]int
	for _, line := range lines[1:] {
		if line == "" {
			if board != nil {
				boards = append(boards, board)
				board = nil
			}
		} else {
			row := utils.StringsToInts(strings.Fields(line))
			board = append(board, row)
		}
	}
	if board != nil {
		boards = append(boards, board)
		board = nil
	}
	return
}

func bingo(boards [][][]int, numbers []int) (int, int) {
	for _, number := range numbers {
		for t, board := range boards {
			for _, row := range board {
				for c := range row {
					if row[c] == number {
						row[c] = 0
						for i := 0; i < len(board); i++ {
							sum := 0
							for j := 0; j < len(row); j++ {
								sum += board[i][j]
							}
							if sum == 0 {
								return t, number
							}
						}
						for j := 0; j < len(row); j++ {
							sum := 0
							for i := 0; i < len(board); i++ {
								sum += board[i][j]
							}
							if sum == 0 {
								return t, number
							}
						}
					}
				}
			}
		}
	}
	return -1, -1
}

func remove(s [][][]int, i int) [][][]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
