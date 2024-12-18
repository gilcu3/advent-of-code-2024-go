package year2024

import (
	"fmt"
	"strconv"
	"strings"

	"aocgen/internal/util"
)

type Day18 struct{}

func (p Day18) realPart1(lines []string, n, k int) string {
	ar := make([][]int, n)
	bad := make([][]int, 0)
	for i := range ar {
		ar[i] = make([]int, n)
	}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		bad = append(bad, []int{x, y})
	}
	for i := range k {
		ar[bad[i][0]][bad[i][1]] = 1
	}
	dd := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	que := [][]int{{0, 0}}
	front := 0
	ar[0][0] = -1
	for front < len(que) {
		c := que[front]
		front += 1
		for d := range 4 {
			nc := []int{c[0] + dd[d][0], c[1] + dd[d][1]}
			if 0 <= nc[0] && nc[0] < n && 0 <= nc[1] && nc[1] < n && ar[nc[0]][nc[1]] == 0 {
				ar[nc[0]][nc[1]] = ar[c[0]][c[1]] - 1
				que = append(que, nc)
			}
		}
	}
	ans := -ar[n-1][n-1] - 1

	return fmt.Sprint(ans)
}
func (p Day18) Part1(lines []string) string {
	return p.realPart1(lines, 71, 1024)
}

func (p Day18) realPart2(lines []string, n int) string {

	bad := make([][]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		bad = append(bad, []int{x, y})
	}

	dd := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	b, e := 0, len(bad)
	ar := make([][]int, n)
	for i := range n {
		ar[i] = make([]int, n)
	}
	for b < e {
		k := (b + e + 1) / 2
		for i := range n {
			for j := range n {
				ar[i][j] = 0
			}
		}
		que := [][]int{{0, 0}}
		front := 0
		for i := range k {
			ar[bad[i][0]][bad[i][1]] = 1
		}
		ar[0][0] = -1
		for front < len(que) {
			c := que[front]
			front += 1
			for d := range 4 {
				nc := []int{c[0] + dd[d][0], c[1] + dd[d][1]}
				if 0 <= nc[0] && nc[0] < n && 0 <= nc[1] && nc[1] < n && ar[nc[0]][nc[1]] == 0 {
					ar[nc[0]][nc[1]] = ar[c[0]][c[1]] - 1
					que = append(que, nc)
				}
			}
		}
		if ar[n-1][n-1] == 0 {
			e = k - 1
		} else {
			b = k
		}
	}

	ans := fmt.Sprintf("%d,%d", bad[b][0], bad[b][1])

	return fmt.Sprint(ans)
}

func (p Day18) Part2(lines []string) string {
	return p.realPart2(lines, 71)
}

func (p Day18) TestPart1() {
	const ansExample1 = "22"
	input := util.ExampleInput(2024, 18, 0)
	ans := p.realPart1(input, 7, 12)
	if ans == fmt.Sprint(nil) {
	} else if ansExample1 == "" {
		fmt.Println("Correct answer Part1 missing, got", ans)
	} else if ans != ansExample1 {
		fmt.Println("Answer to Part1 incorrect", ans, ansExample1)
	} else {
		fmt.Println("Answer to Part1 correct", ans)
	}
}

func (p Day18) TestPart2() {
	const ansExample2 = "6,1"
	input := util.ExampleInput(2024, 18, 0)
	ans := p.realPart2(input, 7)
	if ans == fmt.Sprint(nil) {
	} else if ansExample2 == "" {
		fmt.Println("Correct answer Part2 missing, got", ans)
	} else if ans != ansExample2 {
		fmt.Println("Answer to Part2 incorrect", ans, ansExample2)
	} else {
		fmt.Println("Answer to Part2 correct", ans)
	}
}
