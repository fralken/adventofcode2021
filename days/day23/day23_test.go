package day23

import "testing"

func TestFirstStar(t *testing.T) {
	content :=
`#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`
	result := firstStar(content)
	want := 12521
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content :=
`#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`
	result := secondStar(content)
	want := 44169
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
