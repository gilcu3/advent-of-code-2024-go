package year2024

import (
	"fmt"

	"aocgen/internal/util"
)

type Day20 struct{}

func computeDist(ar [][]int, sx, sy int) [][]int {
	n, m := len(ar), len(ar[0])
	que := [][]int{{sx, sy}}
	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, m)
		for j := range m {
			dist[i][j] = -1
		}
	}
	dist[sx][sy] = 0
	front := 0
	dd := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for front < len(que) {
		c := que[front]
		front += 1
		for d := range 4 {
			nc := []int{c[0] + dd[d][0], c[1] + dd[d][1]}
			if nc[0] < n && nc[0] >= 0 && nc[1] < m && nc[1] >= 0 && ar[nc[0]][nc[1]] == 0 && dist[nc[0]][nc[1]] == -1 {
				que = append(que, nc)
				dist[nc[0]][nc[1]] = dist[c[0]][c[1]] + 1
			}
		}
	}
	return dist
}

func (p Day20) realPart1(lines []string, bound int) string {
	var sx, sy, ex, ey int
	ar := make([][]int, 0)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]int, len(line))
		for j, c := range line {
			if c == '#' {
				row[j] = 1
			} else if c == 'S' {
				sx, sy = i, j
			} else if c == 'E' {
				ex, ey = i, j
			}
		}
		ar = append(ar, row)
	}
	n, m := len(ar), len(ar[0])
	dists := computeDist(ar, sx, sy)
	dis := dists[ex][ey]
	diste := computeDist(ar, ex, ey)
	ans := 0
	for i := range n {
		for j := range m {
			if ar[i][j] == 0 {
				ps := make([][]int, 0)
				for i2 := max(0, i-2); i2 < min(n, i+2+1); i2++ {
					r := 2 - abs(i-i2)
					for j2 := max(0, j-r); j2 < min(m, j+r+1); j2++ {
						if ar[i2][j2] == 0 && abs(i2-i)+abs(j2-j) == 2 {
							ps = append(ps, []int{i2, j2})
						}
					}
				}
				for e := range ps {
					cur := dists[i][j] + 2 + diste[ps[e][0]][ps[e][1]]
					if cur <= dis-bound {
						ans += 1
					}
				}
			}
		}
	}

	return fmt.Sprint(ans)
}

func (p Day20) Part1(lines []string) string {
	return p.realPart1(lines, 100)
}

func (p Day20) realPart2(lines []string, bound int) string {
	cc := 20 - 1
	var sx, sy, ex, ey int
	ar := make([][]int, 0)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]int, len(line))
		for j, c := range line {
			if c == '#' {
				row[j] = 1
			} else if c == 'S' {
				sx, sy = i, j
			} else if c == 'E' {
				ex, ey = i, j
			}
		}
		ar = append(ar, row)
	}
	n, m := len(ar), len(ar[0])
	dists := computeDist(ar, sx, sy)
	dis := dists[ex][ey]
	diste := computeDist(ar, ex, ey)
	ans := 0
	for i := range n {
		for j := range m {
			if ar[i][j] == 0 {
				ts := make([][]int, 0)
				for i2 := max(0, i-cc-1); i2 < min(n, i+cc+1+1); i2++ {
					r := cc + 1 - abs(i-i2)
					for j2 := max(0, j-r); j2 < min(m, j+r+1); j2++ {
						if ar[i2][j2] == 0 {
							ts = append(ts, []int{i2, j2})
						}
					}
				}
				for e := range ts {
					cur := dists[i][j] + abs(i-ts[e][0]) + abs(j-ts[e][1]) + diste[ts[e][0]][ts[e][1]]
					if cur <= dis-bound {
						ans += 1
					}
				}
			}
		}
	}

	return fmt.Sprint(ans)
}

func (p Day20) Part2(lines []string) string {
	return p.realPart2(lines, 100)
}

func (p Day20) TestPart1() {
	const ansExample1 = "44"
	input := util.ExampleInput(2024, 20, 0)
	ans := p.realPart1(input, 1)
	if ans == fmt.Sprint(nil) {
	} else if ansExample1 == "" {
		fmt.Println("Correct answer Part1 missing, got", ans)
	} else if ans != ansExample1 {
		fmt.Println("Answer to Part1 incorrect", ans, ansExample1)
	} else {
		fmt.Println("Answer to Part1 correct", ans)
	}
}

func (p Day20) TestPart2() {
	const ansExample2 = "285"
	input := util.ExampleInput(2024, 20, 0)
	ans := p.realPart2(input, 50)
	if ans == fmt.Sprint(nil) {
	} else if ansExample2 == "" {
		fmt.Println("Correct answer Part2 missing, got", ans)
	} else if ans != ansExample2 {
		fmt.Println("Answer to Part2 incorrect", ans, ansExample2)
	} else {
		fmt.Println("Answer to Part2 correct", ans)
	}
}
