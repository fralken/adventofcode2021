package day24

import (
	"aoc2021/utils"
	"strconv"
	"strings"
)

func FirstStar() {
	utils.Star(24, 1, "largest model number accepted by MONAD", firstStar)
}

func SecondStar() {
	utils.Star(24, 2, "smallest model number accepted by MONAD", secondStar)
}

func firstStar(content string) int {
	set := parseInput(content)
	_, max := findMonads(&set, []int{}, 0, 0)
	return max
}

func secondStar(content string) int {
	set := parseInput(content)
	min, _ := findMonads(&set, []int{}, 0, 0)
	return min
}

// This problem cannot be solved by executing the instructions, it takes too long.
// Based on the pattern of instructions, digits can be found by executing the
// following findMonads(). For the input only 3 parameters per block are relevant
func parseInput(content string) [][]int {
	set := make([][]int, 0)
	blocks := 14
	lines := strings.Split(content, "\n")
	blockSize := len(lines) / blocks
	for i:= 0; i < blocks; i++ {
		a, _ := strconv.Atoi(lines[i * blockSize + 4][6:])
		b, _ := strconv.Atoi(lines[i * blockSize + 5][6:])
		c, _ := strconv.Atoi(lines[i * blockSize + 15][6:])
		set = append(set, []int{a, b, c})
	}
	return set
}

func findMonads(set *[][]int, digits []int, i int, z int) (min int, max int) {
	if len(digits) == len(*set) {
		if z == 0 {
			min = arrayToInt(digits)
			max = min
		} else {
			min = 0
			max = 0
		}
		return 
	}

	for w := 1; w <= 9; w++ {
		nmin := 0
		nmax := 0
		if (*set)[i][0] == 26 {
			if w == z % 26 + (*set)[i][1] {
				nmin, nmax = findMonads(set, append(digits, w), i + 1, z / 26)
			}
		} else { // (*set)[i][0] == 1
			nmin, nmax = findMonads(set, append(digits, w), i + 1, z * 26 + w + (*set)[i][2])
		}
		if nmin > 0 && nmin < min || min == 0 { min = nmin }
		if nmax > 0 && nmax > max || max == 0 { max = nmax }
	}
	return
}

func arrayToInt(digits []int) int {
	sum := 0
	for _, d := range digits {
		sum = sum * 10 + d
	}
	return sum
}

// the code below implements the esecution of instructions
// it is only useful for tests
type subprog struct {
	input *func(int)
	instr []*func()
}

type alucomp struct {
	regs map[byte]int
	program []*subprog
}

func compile(content string) *alucomp {
	lines := strings.Split(content, "\n")
	var alu alucomp
	var subp *subprog
	alu.regs = make(map[byte]int)
	alu.program = []*subprog{}
	for _, line := range lines {
		instr := line[0:3]
		reg := line[4]
		if instr == "inp" {
			subp = &subprog{}
			f := func(r byte) func(int) {
				return func(i int) {
					alu.regs[r] = i
				}
			}
			instrf := f(reg)
			subp.input = &instrf
			alu.program = append(alu.program, subp)
		} else {
			op, err := strconv.Atoi(line[6:])
			var inreg byte
			var instrf func()
			if err != nil {	inreg = line[6] }
			switch instr {
			case "add":
				if err == nil {
					f := func(r byte, o int) func() {
						return func() {
							alu.regs[r] += o
						}
					}
					instrf = f(reg, op)
				} else {
					f := func(ro byte, ri byte) func() {
						return func() {
							alu.regs[ro] += alu.regs[ri]
						}
					}
					instrf = f(reg, inreg)
				}
			case "mul":
				if err == nil {
					f := func(r byte, o int) func() {
						return func() {
							alu.regs[r] *= o
						}
					}
					instrf = f(reg, op)
				} else {
					f := func(ro byte, ri byte) func() {
						return func() {
							alu.regs[ro] *= alu.regs[ri]
						}
					}
					instrf = f(reg, inreg)
				}
			case "div":
				if err == nil {
					f := func(r byte, o int) func() {
						return func() { if o != 0 {
							alu.regs[r] /= o }
						}
					}
					instrf = f(reg, op)
				} else {
					f := func(ro byte, ri byte) func(){
						return func() {
							if alu.regs[ri] != 0 { alu.regs[ro] /= alu.regs[ri] }
						}
					}
					instrf = f(reg, inreg)
				}
			case "mod":
				if err == nil {
					f := func(r byte, o int) func() {
						return func() {
							if o > 0 && alu.regs[r] >= 0 { alu.regs[r] %= o }
						}
					}
					instrf = f(reg, op)
				} else {
					f := func(ro byte, ri byte) func() {
						return func() {
							if alu.regs[ri] > 0 && alu.regs[ro] >= 0 { alu.regs[ro] %= alu.regs[ri] }
						}
					}
					instrf = f(reg, inreg)
				}
			case "eql":
				if err == nil {
					f := func(r byte, o int) func() {
						return func() {
							if alu.regs[r] == o { alu.regs[r] = 1 } else { alu.regs[r] = 0 }
						}
					}
					instrf = f(reg, op)
				} else {
					f := func(ro byte, ri byte) func() {
						return func() {
							if alu.regs[ri] == alu.regs[ro] { alu.regs[ro] = 1 } else { alu.regs[ro] = 0 }
						}
					}
					instrf = f(reg, inreg)
				}
			}
			(*subp).instr = append((*subp).instr, &instrf)
		}
	}
	return &alu
}

func execute(alu *alucomp, i int) {
	for _, f := range (*alu.program[i]).instr {
		(*f)()
	}
}
