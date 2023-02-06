package day18

import (
	"testing"
)

func TestFirstStar1(t *testing.T) {
	content := []string{
		"[[[[[9,8],1],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]",
		"[[6,[5,[4,[3,2]]]],1]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
	}
	want := []string{
		"[[[[0,9],2],3],4]",
		"[7,[6,[5,[7,0]]]]",
		"[[6,[5,[7,0]]],3]",
		"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}
	for i, c := range content {
		n := extractSnailfishes(c)
		n.reduce()
		result := n.print()
		if result != want[i] {
			t.Errorf("Result was incorrect for %s, got: %s, want: %s.", c, result, want[i])
		}
	}
}

func TestFirstStar2(t *testing.T) {
	left := "[[[[4,3],4],4],[7,[[8,4],9]]]"
	right := "[1,1]"
	nl := extractSnailfishes(left)
	nr := extractSnailfishes(right)
	na := nl.add(nr)
	na.reduce()
	result := na.print()
	want := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	if result != want {
		t.Errorf("Result was incorrect got: %s, want: %s.", result, want)
	}
}

func TestFirstStar3(t *testing.T) {
	content := []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
	}
	want := "[[[[1,1],[2,2]],[3,3]],[4,4]]"
	sum := extractSnailfishes(content[0])
	for _, c := range content[1:] {
		n := extractSnailfishes(c)
		sum = sum.add(n)
		sum.reduce()
	}
	result := sum.print()
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}
}

func TestFirstStar4(t *testing.T) {
	content := []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
		"[5,5]",
	}
	want := "[[[[3,0],[5,3]],[4,4]],[5,5]]"
	sum := extractSnailfishes(content[0])
	for _, c := range content[1:] {
		n := extractSnailfishes(c)
		sum = sum.add(n)
		sum.reduce()
	}
	result := sum.print()
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}
}

func TestFirstStar5(t *testing.T) {
	content := []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
		"[5,5]",
		"[6,6]",
	}
	want := "[[[[5,0],[7,4]],[5,5]],[6,6]]"
	sum := extractSnailfishes(content[0])
	for _, c := range content[1:] {
		n := extractSnailfishes(c)
		sum = sum.add(n)
		sum.reduce()
	}
	result := sum.print()
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}
}

func TestFirstStar6(t *testing.T) {
	content := []string{
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
		"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
		"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
		"[7,[5,[[3,8],[1,4]]]]",
		"[[2,[2,2]],[8,[8,1]]]",
		"[2,9]",
		"[1,[[[9,3],9],[[9,0],[0,7]]]]",
		"[[[5,[7,4]],7],1]",
		"[[[[4,2],2],6],[8,7]]",
	}
	want := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"
	sum := extractSnailfishes(content[0])
	for _, c := range content[1:] {
		n := extractSnailfishes(c)
		sum = sum.add(n)
		sum.reduce()
	}
	result := sum.print()
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}
}

func TestFirstStar7(t *testing.T) {
	content := []string{
		"[9,1]",
		"[1,9]",
		"[[9,1],[1,9]]",
		"[[1,2],[[3,4],5]]",
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
	}
	want := []int{
		29,
		21,
		129,
		143,
		1384,
		445,
		791,
		1137,
		3488,
	}
	for i, c := range content {
		n := extractSnailfishes(c)
		n.reduce()
		result := n.magnitude()
		if result != want[i] {
			t.Errorf("Result was incorrect for %s, got: %d, want: %d.", c, result, want[i])
		}
	}
}

func TestFirstStar8(t *testing.T) {
	content := 
	`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`
	result := firstStar(content)
	want := 4140
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar(t *testing.T) {
	content := 
	`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`
	result := secondStar(content)
	want := 3993
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
