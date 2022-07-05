package day01

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
		`199
200
208
210
200
207
240
269
260
263`
	result := firstStar(content)
	want := 7
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
		`199
200
208
210
200
207
240
269
260
263`
	result := secondStar(content)
	want := 5
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
