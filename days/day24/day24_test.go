package day24

import "testing"

func TestFirstStar1(t *testing.T) {
	content :=
	`inp x
mul x -1`
	alu := compile(content)
	(*alu.program[0].input)(2)
	execute(alu, 0)
	result := alu.regs['x']
	want := -2
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestFirstStar2(t *testing.T) {
	content :=
	`inp z
inp x
mul z 3
eql z x`
	alu := compile(content)
	(*alu.program[0].input)(2)
	execute(alu, 0)
	(*alu.program[1].input)(6)
	execute(alu, 1)
	result := alu.regs['z']
	want := 1
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestFirstStar3(t *testing.T) {
	content :=
	`inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2`
	alu := compile(content)
	want := 13
	(*alu.program[0].input)(want)
	execute(alu, 0)
	result := alu.regs['z'] + 2*alu.regs['y'] + 4*alu.regs['x'] + 8*alu.regs['w']
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
