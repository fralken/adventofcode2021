package day02

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
		`forward 5
down 5
forward 8
up 3
down 8
forward 2`
	result := firstStar(content)
	want := 150
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
		`forward 5
down 5
forward 8
up 3
down 8
forward 2`
	result := secondStar(content)
	want := 900
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
