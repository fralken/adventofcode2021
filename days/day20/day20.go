package day20

import (
	"aoc2021/utils"
	"strings"
)

func FirstStar() {
	utils.Star(20, 1, "lit pixels in the resulting image after 2 steps", firstStar)
}

func SecondStar() {
	utils.Star(20, 2, "lit pixels in the resulting image after 50 steps", secondStar)
}

func firstStar(content string) int {
	image := enhanceImage(content, 2)
	return countLitPixels(image)
}

func secondStar(content string) int {
	image := enhanceImage(content, 50)
	return countLitPixels(image)
}

func enhanceImage(content string, steps int) [][]uint8 {
	algorithm, image := extractAlgorithmAndImage(content)
	for step := 0; step < steps; step++ {
		image = filterImage(image, algorithm, step)
	}
	return image
}

func extractAlgorithmAndImage(content string) (algorithm []uint8, image [][]uint8) {
	parts := strings.Split(content, "\n\n")
	algorithm = extractPattern(parts[0])
	lines := strings.Split(parts[1], "\n")
	image = extractImage(lines)
	return
}

func extractPattern(content string) (pattern []uint8) {
	for _, c := range content {
		if c == '#' {
			pattern = append(pattern, 1)
		} else { // c == '.'
			pattern = append(pattern, 0)
		}
	}
	return
}

func extractImage(lines []string) (image [][]uint8) {
	for _, line := range lines {
		image = append(image, extractPattern(line))
	}
	return
}

func filterImage(source [][]uint8, algorithm []uint8, step int) [][]uint8 {
	rows := len(source) + 2
	cols := len(source[0]) + 2
	dest := make([][]uint8, rows)
	var infinitePixel uint8
	if step == 0 || algorithm[0] == 0 {
		infinitePixel = 0
	} else if step % 2 == 1 {
		infinitePixel = algorithm[0] // 1
	} else {
		infinitePixel = algorithm[len(algorithm)-1]
	} 
	for i := 0; i < rows; i++ {
		dest[i] = make([]uint8, cols)
		for j := 0; j < cols; j++ {
			dest[i][j] = filter(i, j, rows, cols, source, algorithm, infinitePixel)
		}
	}
	return dest
}

func filter(i int, j int, rows int, cols int, source [][]uint8, algorithm []uint8, infinitePixel uint8) uint8 {
	val := uint16(0)
	for di := i - 1; di <= i + 1; di++ {
		for dj := j - 1; dj <= j + 1; dj++ {
			if di <= 0 || dj <= 0 || di >= rows-1 || dj >= cols-1 {
				val = (val << 1) | uint16(infinitePixel)
			} else {
				val = (val << 1) | uint16(source[di - 1][dj - 1])
			}
		}
	}
	return algorithm[val]
}

func countLitPixels(image [][]uint8) int {
	count := 0
	for _, row := range image {
		for _, c := range row {
			count += int(c)
		}
	}
	return count
}