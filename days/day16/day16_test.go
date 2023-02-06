package day16

import "testing"

func TestFirstStar1(t *testing.T) {
	content := "8A004A801A8002F478"
	result := firstStar(content)
	want := 16
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestFirstStar2(t *testing.T) {
	content := "620080001611562C8802118E34"
	result := firstStar(content)
	want := 12
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestFirstStar3(t *testing.T) {
	content := "C0015000016115A2E0802F182340"
	result := firstStar(content)
	want := 23
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestFirstStar4(t *testing.T) {
	content := "A0016C880162017C3686B18A3D4780"
	result := firstStar(content)
	want := 31
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar1(t *testing.T) {
	content := "C200B40A82"
	result := secondStar(content)
	want := 3
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar2(t *testing.T) {
	content := "04005AC33890"
	result := secondStar(content)
	want := 54
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar3(t *testing.T) {
	content := "880086C3E88112"
	result := secondStar(content)
	want := 7
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar4(t *testing.T) {
	content := "CE00C43D881120"
	result := secondStar(content)
	want := 9
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar5(t *testing.T) {
	content := "D8005AC2A8F0"
	result := secondStar(content)
	want := 1
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar6(t *testing.T) {
	content := "F600BC2D8F"
	result := secondStar(content)
	want := 0
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar7(t *testing.T) {
	content := "9C005AC2F8F0"
	result := secondStar(content)
	want := 0
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}

func TestSecondStar8(t *testing.T) {
	content := "9C0141080250320F1802104A08"
	result := secondStar(content)
	want := 1
	if result != want {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, want)
	}
}
