package day07

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
		`16,1,2,0,4,2,7,1,2,14`
	result := firstStar(content)
	want := 37
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
		`16,1,2,0,4,2,7,1,2,14`
	result := secondStar(content)
	want := 168
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
