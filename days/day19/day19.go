package day19

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

func FirstStar() {
	utils.Star(19, 1, "number of beacons", firstStar)
}

func SecondStar() {
	utils.Star(19, 2, "largest Manhattan distance", secondStar)
}

func firstStar(content string) int {
	scanners := parseScanners(content)
	pairs := overlappingScanners(scanners)
	beacons, _ := translateAndRotateScanners(scanners, pairs)
	return beacons
}

func secondStar(content string) int {
	scanners := parseScanners(content)
	pairs := overlappingScanners(scanners)
	_, origins := translateAndRotateScanners(scanners, pairs)
	distance := 0
	for i, a := range origins[:len(origins)-1] {
		for _, b := range origins[i+1:] {
			d := manhattan(a, b)
			if d > distance {
				distance = d
			}
		}
	}
	return distance
}

type point struct {
	x, y, z int
}

// array of beacons per scanner and set of
// distances between any two pair of beacons
type scanner struct {
	beacons []*point
	dist    map[int][2]*point
}

// pair of scanner indices that share at least 12 beacons
type pair struct {
	a, b int
}

// store a rotation index and a location of a scanner
type rotloc struct {
	r int
	t point
}

func parseScanners(content string) []*scanner {
	// compute distance of every pair of beacons per scanner
	// for simplicity we use manhattan distance
	computeDistances := func(s *scanner) {
		for i, a := range s.beacons[:len(s.beacons)-1] {
			for _, b := range s.beacons[i+1:] {
				d := manhattan(a, b)
				s.dist[d] = [2]*point{a, b}
			}
		}
	}
	blocks := strings.Split(content, "\n\n")
	var scanners []*scanner
	for _, block := range blocks {
		positions := strings.Split(block, "\n")[1:] // ignore header
		s := scanner{[]*point{}, map[int][2]*point{}}
		for _, pos := range positions {
			var p point
			fmt.Sscanf(pos, "%d,%d,%d", &p.x, &p.y, &p.z)
			s.beacons = append(s.beacons, &p)
		}
		computeDistances((&s))
		scanners = append(scanners, &s)
	}
	return scanners
}

func overlappingScanners(scanners []*scanner) map[pair][]int {
	// we need to find scanners that overlap with at least 12 beacons
	// which means that (12 * 11 / 2) distances must match
	// (the sum of first n integers is (n + 1) * n / 2)
	// return for each pair an array of common distances
	count := 12 * 11 / 2
	pairs := make(map[pair][]int)
	for i, s1 := range scanners[:len(scanners)-1] {
		for j, s2 := range scanners[i+1:] {
			list := []int{}
			for k1 := range s1.dist {
				for k2 := range s2.dist {
					if k1 == k2 {
						list = append(list, k1)
					}
				}
			}
			if len(list) >= count {
				pairs[pair{i, i + 1 + j}] = list
			}
		}
	}
	return pairs
}

func translateAndRotateScanners(scanners []*scanner, pairs map[pair][]int) (int, []*point) {
	// scanner[0] is the reference, so its origin is {0,0,0} and it is not rotated
	// keep a set of unique beacons and we add beacons of scanner[0]
	unique := make(map[point]bool)
	for _, beacon := range scanners[0].beacons {
		unique[*beacon] = true
	}
	// keep an array of positions of each scanner, relative to scanner[0]
	origin := make([]*point, len(scanners))
	origin[0] = &point{0, 0, 0}
	// keep an array of indices of rotations of each scanner, relative to scanner[0]
	rotation := make([]int, len(scanners))
	rotation[0] = 0
	move := len(scanners)
	// rotate and translate scanners relative to scanner[0]
	// if origin is defined it means we have already moved a scanner
	for move > 1 {
		for p, dists := range pairs {
			a := p.a
			b := p.b
			if origin[a] == nil && origin[b] != nil {
				a, b = b, a
			}
			if origin[a] != nil && origin[b] == nil {
				rt := make(map[rotloc]bool)
				for _, dist := range dists {
					pa := scanners[a].dist[dist]
					pb := scanners[b].dist[dist]
					// find relative location between scanners and relative rotation
					loc, rb := findLocationAndRotation(pa, rotation[a], pb)
					if loc != nil {
						// make sure that for at least two common distances we find
						// the same relative location and translation. For this we use
						// a set of already found locations and rotations
						if _, ok := rt[rotloc{rb, *loc}]; ok {
							origin[b] = add(origin[a], loc)
							rotation[b] = rb
							// add rotated and translated beacons to the set of unique beacons
							for _, beacon := range scanners[b].beacons {
								unique[*add(origin[b], beacon.rotate(rb))] = true
							}
							move--
							break
						} else {
							rt[rotloc{rb, *loc}] = true
						}
					}
				}
			}
		}
	}
	return len(unique), origin
}

func findLocationAndRotation(a [2]*point, ra int, b [2]*point) (*point, int) {
	// "a" is a pair of beacons at the same distance of beacons of pair "b".
	// If the two tranlsation between beacons "a" and "b" are equal we have found
	// the relative location between the two scanners.
	// We must rotate scanner "a" with the already known rotation, and we must find
	// the rotation of scanner "b" such that the two translations between beacons match
	a0 := a[0].rotate(ra)
	a1 := a[1].rotate(ra)
	for i := 0; i < 24; i++ {
		b0 := b[0].rotate(i)
		b1 := b[1].rotate(i)
		l0 := sub(a0, b0)
		l1 := sub(a1, b1)
		if *l0 == *l1 {
			return l0, i
		}
		l0 = sub(a0, b1)
		l1 = sub(a1, b0)
		if *l0 == *l1 {
			return l0, i
		}
	}
	return nil, -1
}

// these are the 24 rotations of 90 degrees around axis x, y, z
// https://www.euclideanspace.com/maths/algebra/matrix/transforms/examples/index.htm
func (c *point) rotate(i int) *point {
	x, y, z := c.x, c.y, c.z
	return &[]point{
		{x, y, z},
		{y, z, x},
		{z, x, y},
		{x, -y, -z},
		{-y, -z, x},
		{-z, x, -y},
		{-x, y, -z},
		{y, -z, -x},
		{-z, -x, y},
		{-x, -y, z},
		{-y, z, -x},
		{z, -x, -y},
		{x, -z, y},
		{-z, y, x},
		{y, x, -z},
		{x, z, -y},
		{z, -y, x},
		{-y, x, z},
		{-x, z, y},
		{z, y, -x},
		{y, -x, z},
		{-x, -z, -y},
		{-z, -y, -x},
		{-y, -x, -z},
	}[i]
}

func manhattan(a, b *point) int {
	return abs(a.x-b.x) + abs(a.y-b.y) + abs(a.z-b.z)
}

func add(a, b *point) *point {
	return &point{a.x + b.x, a.y + b.y, a.z + b.z}
}

func sub(a, b *point) *point {
	return &point{a.x - b.x, a.y - b.y, a.z - b.z}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
