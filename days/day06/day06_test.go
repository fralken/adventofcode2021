package day06

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
		`3,4,3,1,2`
	result := firstStar(content)
	want := 5934
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
		`3,4,3,1,2`
	result := secondStar(content)
	want := 26984457539
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
