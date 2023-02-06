package day21

import (
	"aoc2021/utils"
	"fmt"
)

func FirstStar() {
	utils.Star(21, 1, "score of loser times number of die rolls", firstStar)
}

func SecondStar() {
	utils.Star(21, 2, "number of universes a player wins most", secondStar)
}

func firstStar(content string) int {
	start := playersPositions(content)
	score, rolls := play(start)
	var loser int
	if score[0] < score[1] { loser = score[0] } else { loser = score[1] }
	return loser * rolls
}

func secondStar(content string) int {
	start := playersPositions(content)
	u1, u2 := playQuantum(start[0], start[1], 0, 0, getCombinations())
	if u1 > u2 { return u1 } else { return u2 }
}

func play(position [2]int) ([2]int, int) {
	score := [2]int{0, 0}
	player := 0
	die := 1
	rolls := 0
	for score[0] < 1000 && score[1] < 1000 {
		position[player] = (position[player] + die * 3 + 3) % 10
		if position[player] == 0 { position[player] = 10 }
		score[player] += position[player]
		die = (die + 3) % 100
		rolls += 3
		player = 1 - player
	}
	return score, rolls
}

func playQuantum(position1 int, position2 int, score1 int, score2 int, combinations *map[int]int) (u1 int, u2 int) {
	u1 = 0
	u2 = 0
	for rolls, universes := range *combinations {
		// create a copy of the current game
		// current players positions and scores
		p1 := position1
		p2 := position2
		s1 := score1
		s2 := score2
		// first player moves
		p1 = (p1 + rolls) % 10
		if p1 == 0 { p1 = 10 }
		s1 += p1
		if s1 >= 21 {
			u1 += universes
		} else {
			// play quantum for the other player (swap players)
			w2, w1 := playQuantum(p2, p1, s2, s1, combinations) 
			u1 += (w1 * universes)
			u2 += (w2 * universes)
		}
	}
	return
}

func playersPositions(content string) [2]int {
	var p1 int
	var p2 int
	fmt.Sscanf(content, "Player 1 starting position: %d\nPlayer 2 starting position: %d", &p1, &p2)
	return [2]int{p1, p2}
}

// compute how many universes each sum of dice rolls spawns
func getCombinations() *map[int]int {
	c := make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				c[i+j+k] += 1
			}
		}
	}
	return &c
}
