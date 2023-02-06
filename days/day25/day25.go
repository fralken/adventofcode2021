package day25

import (
	"aoc2021/utils"
	"strings"
)

func FirstStar() {
	utils.Star(25, 1, "first step on which no sea cucumbers move", firstStar)
}

func SecondStar() {
	utils.Star(25, 2, "THE END ...", secondStar)
}

func firstStar(content string) int {
	sea := scanSeaCucumbers(content)
	count := 0
	for equal := false; !equal; count++ {
		sea, equal = step(sea)
	}
	return count
}

func secondStar(content string) string {
	return "Thanks for watching"
}

func scanSeaCucumbers(content string) *[][]byte {
	var sea [][]byte
	rows := strings.Split(content, "\n")
	for r, row := range rows {
		sea = append(sea, make([]byte,len(row)))
		for c, col := range row {
			sea[r][c] = byte(col)
		}
	}
	return &sea
}

func step(sea *[][]byte) (*[][]byte, bool) {
	equal := true
	var next1 [][]byte
	for r, row := range *sea {
		next1 = append(next1, make([]byte, len(row)))
		lastc := len(row)-1
		for c := lastc; c >= 0; c-- {
			if (*sea)[r][c] == '.' {
				if c == 0 && (*sea)[r][lastc] == '>' || c > 0 && (*sea)[r][c-1] == '>' {
					next1[r][c] = '>'
					if c == 0 { next1[r][lastc] = '.' } else { next1[r][c-1] = '.' }
					c--
					equal = false
					continue
				}
			}
			next1[r][c] = (*sea)[r][c]
		}
	}
	var next2 [][]byte
	for _, row := range next1 {
		next2 = append(next2, make([]byte, len(row)))
	}
	lastr := len(*sea)-1
	for c := range next1[0] {
		for r := lastr; r >= 0; r-- {
			if next1[r][c] == '.' {
				if r == 0 && next1[lastr][c] == 'v' || r > 0 && next1[r-1][c] == 'v' {
					next2[r][c] = 'v'
					if r == 0 { next2[lastr][c] = '.' } else { next2[r-1][c] = '.' }
					r--
					equal = false
					continue
				}
			}
			next2[r][c] = next1[r][c]
		}
	}
	return &next2, equal
}