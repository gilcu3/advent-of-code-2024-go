package year2024

import (
	"fmt"
	"regexp"
	"strconv"

	"aocgen/internal/util"
)

type Day14 struct{}

func (p Day14) realPart1(lines []string, n, m int) string {
	re := regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)

	px, py, vx, vy := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if matches := re.FindStringSubmatch(line); matches != nil {
			x, _ := strconv.Atoi(matches[1])
			px = append(px, x)
			y, _ := strconv.Atoi(matches[2])
			py = append(py, y)
			x, _ = strconv.Atoi(matches[3])
			vx = append(vx, x)
			y, _ = strconv.Atoi(matches[4])
			vy = append(vy, y)
		}
	}
	steps := 100
	quad := [][]int{{0, 0}, {0, 0}}
	for i := range len(px) {
		x, y, sx, sy := px[i], py[i], vx[i], vy[i]
		for range steps {
			x, y = (x+sx+n)%n, (y+sy+m)%m
		}
		if x != n/2 && y != m/2 {
			quad[2*x/n][2*y/m] += 1
		}
	}
	ans := 1
	for i := range 2 {
		for j := range 2 {
			ans *= quad[i][j]
		}
	}

	return fmt.Sprint(ans)
}

func (p Day14) Part1(lines []string) string {
	return p.realPart1(lines, 101, 103)
}

// func render(ar [][]int, n, m int) {
// 	fmt.Println()

// 	red := "\033[31m"
// 	reset := "\033[0m"
// 	for i := range n {
// 		for j := range m {
// 			if ar[i][j] > 0 {
// 				fmt.Printf("%s%d%s", red, ar[i][j], reset)
// 			} else {
// 				fmt.Print(0)
// 			}

// 		}
// 		fmt.Println()
// 	}

// 	fmt.Println()
// }

func visit(ar [][]int, seen [][]int, mark int, dd [][]int, px, py int) int {

	n, m := len(ar), len(ar[0])

	qx, qy := []int{px}, []int{py}
	seen[px][py] = mark
	front := 0
	ans := 0
	for front < len(qx) {
		cx, cy := qx[front], qy[front]
		front += 1
		ans += ar[cx][cy]
		for d := range 4 {
			nx, ny := cx+dd[d][0], cy+dd[d][1]
			if 0 <= nx && nx < n && 0 <= ny && ny < m && ar[nx][ny] > 0 && seen[nx][ny] != mark {
				seen[nx][ny] = mark
				qx, qy = append(qx, nx), append(qy, ny)
			}
		}
	}
	return ans

}

func (p Day14) realPart2(lines []string, n, m int) string {
	re := regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)

	px, py, vx, vy := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if matches := re.FindStringSubmatch(line); matches != nil {
			x, _ := strconv.Atoi(matches[1])
			px = append(px, x)
			y, _ := strconv.Atoi(matches[2])
			py = append(py, y)
			x, _ = strconv.Atoi(matches[3])
			vx = append(vx, x)
			y, _ = strconv.Atoi(matches[4])
			vy = append(vy, y)
		}
	}
	ar := make([][]int, n)
	for i := range n {
		ar[i] = make([]int, m)
	}
	steps := 0
	mx := 0
	seen := make([][]int, n)
	for i := range n {
		seen[i] = make([]int, m)
	}
	mark := 0
	dd := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for {

		for i := range len(px) {
			ar[px[i]][py[i]] += 1
		}
		steps += 1

		cur := 0
		mark += 1
		for i := range len(px) {
			if seen[px[i]][py[i]] != mark {
				cc := visit(ar, seen, mark, dd, px[i], py[i])
				cur = max(cur, cc)
			}
		}
		mx = max(mx, cur)
		if cur >= len(px)/3 { // interpretation of majority, a posteriori
			// render(ar, n, m)
			break
		}
		for i := range len(px) {
			ar[px[i]][py[i]] -= 1
		}

		for i := range len(px) {
			px[i], py[i] = (px[i]+vx[i]+n)%n, (py[i]+vy[i]+m)%m
		}

	}

	return fmt.Sprint(steps - 1)
}

func (p Day14) Part2(lines []string) string {
	return p.realPart2(lines, 101, 103)
}

func (p Day14) TestPart1() {
	const ansExample1 = "12"
	input := util.ExampleInput(2024, 14, 0)
	ans := p.realPart1(input, 11, 7)
	if ans == fmt.Sprint(nil) {
	} else if ansExample1 == "" {
		fmt.Println("Correct answer Part1 missing, got", ans)
	} else if ans != ansExample1 {
		fmt.Println("Answer to Part1 incorrect", ans, ansExample1)
	} else {
		fmt.Println("Answer to Part1 correct", ans)
	}
}

func (p Day14) TestPart2() {
	const ansExample2 = "3"
	input := util.ExampleInput(2024, 14, 0)
	ans := p.realPart2(input, 11, 7)
	if ans == fmt.Sprint(nil) {
	} else if ansExample2 == "" {
		fmt.Println("Correct answer Part2 missing, got", ans)
	} else if ans != ansExample2 {
		fmt.Println("Answer to Part2 incorrect", ans, ansExample2)
	} else {
		fmt.Println("Answer to Part2 correct", ans)
	}
}
