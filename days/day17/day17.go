package day17

import (
	"aoc2021/utils"
	"fmt"
	"math"
)

func FirstStar() {
	utils.Star(17, 1, "highest y position in trajectory", firstStar)
}

func SecondStar() {
	utils.Star(17, 2, "distinct initial velocity values", secondStar)
}

// y velocity slows by 1 at each step, the highest position is when y velocity reaches 0
// sum of first n integers is n*(n+1)/2
// if initial y velocity is n=(-ymin-1) then at position y=0 velocity will be -n=ymin+1 so 
// final y position will be ymin that is within the target
// n*(n+1)/2 = (-ymin-1)*(-ymin-1+1)/2 = ymin*(ymin+1)/2
func firstStar(content string) int {
	_, _, ymin, _ := parseTarget(content)
	return ymin * (ymin + 1) / 2
}

func secondStar(content string) int {
	xmin, xmax, ymin, ymax := parseTarget(content)
	count := 0
	// sum of first n integers is (n^2 + n)/2 so
	// vx starts from n such that (n^2 + n)/2=xmin that is
	// n^2 + n - 2*xmin = 0, hence n = -1 + math.Sqrt(1+8*xmin)/2
	// that we can approximate with:
	vxStart := int(math.Sqrt(float64(xmin * 2)) - 1)

	vxstop := make(map[int]bool)
	vxs := make(map[int][]int)
	for vxInit := vxStart; vxInit <= xmax; vxInit++ {
		x, vx, step := 0, vxInit, 0
		for x <= xmax && vx != 0 {
			x += vx
			if vx > 0 { vx -= 1}
			step += 1
			if xmin <= x && x <= xmax {
				vxs[vxInit] = append(vxs[vxInit], step)
				if vx == 0 { vxstop[vxInit] = true }
			}
		}

	}

	vys := make(map[int][]int)
	for vyInit := ymin; vyInit <= -ymin; vyInit++ {
		y, vy, step := 0, vyInit, 0
		for y >= ymin {
			y += vy
			vy -= 1
			step += 1
			if ymin <= y && y <= ymax {
				vys[vyInit] = append(vys[vyInit], step)
			}
		}

	}
	
	for vxInit, xsteps := range vxs {
		for _, ysteps := range vys {
			for _, ys := range ysteps {
				found := false
				for _, xs := range xsteps {
					if xs == ys || xs < ys && vxstop[vxInit] {
						count++
						found = true
						break;
					}
				}
				if found { break }
			}
		}
	}

	return count
}

func parseTarget(content string) (xmin int, xmax int, ymin int, ymax int) {
	fmt.Sscanf(content, "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)
	return
}
