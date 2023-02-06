package day23

import (
	"aoc2021/utils"
	"container/heap"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(23, 1, "least energy for 8 amphipods", firstStar)
}

func SecondStar() {
	utils.Star(23, 2, "least energy for 16 amphipods", secondStar)
}

func firstStar(content string) int {
	b := makeBurrow(content, -1, nil)
	return b.move()
}

func secondStar(content string) int {
	b := makeBurrow(content, 3, []string{"  #D#C#B#A#", "  #D#B#A#C#"})
	return b.move()
}

func makeBurrow(content string, pos int, rooms []string) burrow {
	lines := strings.Split(content, "\n")
	b := ""
	size := burrow(b).width()
	for i, line := range lines {
		if i == pos {
			for _, room := range rooms {
				b = fmt.Sprintf("%s%-*s", b, size, room)
			}
		}
		b = fmt.Sprintf("%s%-*s", b, size, line)
	}
	return burrow(b)
}

func (b burrow) move() int {
	roomOf := map[rune]int{ 'A': 3, 'B': 5, 'C': 7, 'D': 9 }
	energyOf := map[rune]int{ 'A': 1, 'B': 10, 'C': 100, 'D': 1000 }
	limit := b.roomLimit()
	totalEnergy := 0
	totalAmphipods := len(b.findAmphipods())
	// map containing visited configurations with minimum required energy
	states := make(map[burrow]int)
	states[b] = 0
	// priority queue based on minimum energy of visited configurations
	q := queue{}
	heap.Init(&q)
	heap.Push(&q, &state{0, b})
	// update the priority queue if configuration is new or requires less energy
	addState := func(burrow burrow, energy int) {
		if e, ok := states[burrow]; !ok || e > energy {
			states[burrow] = energy
			heap.Push(&q, &state{ energy, burrow })
		}
	}
	// check if amphipod can stay in this position in the hallway, that is not in front of rooms
	canStay := func(p coord) bool {	return p.r == 1 && p.c != 3 && p.c != 5 && p.c != 7 && p.c != 9	}
	for len(q) > 0 {
		s := heap.Pop(&q).(*state)
		b := s.burrow
		e := s.energy
		if totalEnergy == 0 || totalEnergy >= e {
			finish := 0
			amphipods := b.findAmphipods()
			for _, a := range amphipods {
				if b.isAtHome(a.name, a.pos, roomOf[a.name]) {
					finish++
				} else if a.pos.r > 1 {
					// it is in wrong room, try to move out
					steps := 0
					r := a.pos.r - 1
					for r > 1 && b.isFree(coord{r, a.pos.c}) {
						r--
						steps++
					}
					steps++
					if r == 1 {
						// it is in the hallway, move left and right
						// steps -> number of steps to go from wrong room to the hallway
						// i -> number of steps in the hallway
						// m -> direction: -1 left, 1 right
						// k - 1 -> number of steps in the room
						for i, hitleft, hitright := 1, false, false; !hitleft || !hitright; i++ {
							for _, m := range []int{ -1, 1 } {
								j := a.pos.c + i * m
								if ((j >= 1 && m < 0 && !hitleft) || (j <= 11 && m > 0 && !hitright)) && b.isFree(coord{1, j}) {
									if canStay(coord{1, j}) {
										addState(b.update(a.name, a.pos, coord{1, j}),
											e + (steps + i) * energyOf[a.name])
									} else if j == roomOf[a.name] {
										// try entering room
										for k := 2; k < limit; k++ {
											if b.isFree(coord{k, j}) && b.isAtHome(a.name, coord{k, j}, roomOf[a.name]) {
												addState(b.update(a.name, a.pos, coord{k, j}),
													e + (steps + i + k - 1) * energyOf[a.name])
											}
										}
									}
								} else if m > 0 {
									hitright = true
								} else { // m < 0 
									hitleft = true
								}
							}
						}
					}
				} else {
					// it is in the hallway, try entering room
					// default: room is on the right
					start := a.pos.c + 1
					room := roomOf[a.name]
					step := 1
					if a.pos.c > roomOf[a.name] {
						// room is on the left, change direction
						start = a.pos.c - 1
						step = -1
					}
					j := start
					steps := 0
					for ; j != room && b.isFree(coord{a.pos.r, j}); j += step {
						steps++
					}
					if j == room {
						// try entering the room
						// steps -> number of steps in the hallway
						// k - 1 -> number of steps in the room
						steps++
						for k := 2; k < limit; k++ {
							if b.isFree(coord{k, j}) && b.isAtHome(a.name, coord{k, j}, roomOf[a.name]) {
								addState(b.update(a.name, a.pos, coord{k, j}),
									e + (steps + k - 1) * energyOf[a.name])
							}
						}
					}
				}
			}
			if finish == totalAmphipods {
				if totalEnergy == 0 || totalEnergy > e {
					totalEnergy = e
				}
			}
		}
	}
	return totalEnergy
}

func (b burrow) findAmphipods() []amphipod {
	a := []amphipod{}
	size := b.width()
	for i, cell := range b {
		if cell != '#' && cell != '.' && cell != ' ' {
			r := i / size
			c := i - (r * size)
			a = append(a, amphipod{ cell, coord{ r, c } })
		}
	}
	return a
}

func (b burrow) isAtHome(name rune, pos coord, room int) bool {
	size := b.width()
	limit := b.roomLimit()
	if pos.c != room || pos.r == 1 { return false }
	for r := pos.r + 1; r < limit; r++ {
		i := r * size + pos.c
		if b[i] != byte(name) { return false }
	}
	return true
}

func (b burrow) isFree(p coord) bool {
	size := b.width()
	return b[p.r * size + p.c] == '.'
}

func (b burrow) update(name rune, old, new coord) burrow {
	newBurrow := []byte(b)
	size := b.width()
	newBurrow[old.r * size + old.c] = '.'
	newBurrow[new.r * size + new.c] = byte(name)
	return burrow(newBurrow)
}

func (b burrow) width() int { return 13 }
 
func (b burrow) roomLimit() int { return len(b) / 13 - 1 }

type coord struct {
	r, c int
}

type amphipod struct {
	name rune
	pos coord
}

type burrow string

type state struct {
	energy int
	burrow burrow
}

type queue []*state

func (q queue) Len() int { return len(q) }

func (q queue) Less(i, j int) bool {
	return q[i].energy < q[j].energy
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(x any) {
	item := x.(*state)
	*q = append(*q, item)
}

func (q *queue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*q = old[:n-1]
	return item
}
