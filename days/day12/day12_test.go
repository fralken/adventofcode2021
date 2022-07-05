package day12

import "testing"

func TestFirstStar1(t *testing.T) {
	content :=
		`start-A
start-b
A-c
A-b
b-d
A-end
b-end`
	result := firstStar(content)
	want := 10
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestFirstStar2(t *testing.T) {
	content :=
		`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`
	result := firstStar(content)
	want := 19
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestFirstStar3(t *testing.T) {
	content :=
		`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`
	result := firstStar(content)
	want := 226
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar1(t *testing.T) {
	content :=
		`start-A
start-b
A-c
A-b
b-d
A-end
b-end`
	result := secondStar(content)
	want := 36
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar2(t *testing.T) {
	content :=
		`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`
	result := secondStar(content)
	want := 103
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar3(t *testing.T) {
	content :=
		`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`
	result := secondStar(content)
	want := 3509
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
