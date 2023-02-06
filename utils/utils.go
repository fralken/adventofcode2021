package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func StringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func Star[T int|string](day int, star int, msg string, run func(string) T) {
	content := ReadFile(fmt.Sprintf("./input/day%02d.txt", day))
	fmt.Printf("day %2d.%d - %s: %v\n", day, star, msg, run(content))
}