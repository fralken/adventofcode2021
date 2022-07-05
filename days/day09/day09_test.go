package day09

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
		`2199943210
3987894921
9856789892
8767896789
9899965678`
	result := firstStar(content)
	want := 15
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
		`2199943210
3987894921
9856789892
8767896789
9899965678`
	result := secondStar(content)
	want := 1134
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
