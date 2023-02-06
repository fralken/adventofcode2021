package day17

import "testing"

func TestFirstStar(t *testing.T) {
	content := "target area: x=20..30, y=-10..-5"
	result := firstStar(content)
	want := 45
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content := "target area: x=20..30, y=-10..-5"
	result := secondStar(content)
	want := 112
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
