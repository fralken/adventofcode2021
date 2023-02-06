package day21

import "testing"

func TestFirstStar(t *testing.T) {
	content := `Player 1 starting position: 4
Player 2 starting position: 8`
	result := firstStar(content)
	want := 739785
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content := `Player 1 starting position: 4
Player 2 starting position: 8`
	result := secondStar(content)
	want := 444356092776315
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
