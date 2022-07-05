package day03

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
		`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	result := firstStar(content)
	want := 198
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
		`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	result := secondStar(content)
	want := 230
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
