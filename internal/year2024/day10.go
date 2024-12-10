package year2024

import (
	"fmt"

	"aocgen/internal/util"
)

type Day10 struct{}

func (p Day10) Part1(lines []string) string {
	sx, sy := make([]int, 0), make([]int, 0)
	ar := make([][]int, 0)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]int, len(line))
		for j, c := range line {
			row[j] = int(byte(c) - byte('0'))
			if row[j] == 0 {
				sx = append(sx, i)
				sy = append(sy, j)
			}
		}
		ar = append(ar, row)
	}
	n, m := len(ar), len(ar[0])
	dd := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	ans := 0
	for i := range len(sx) {
		seen := make([][]bool, n)
		for i := range len(seen) {
			seen[i] = make([]bool, m)
		}
		qx, qy := []int{sx[i]}, []int{sy[i]}
		seen[sx[i]][sy[i]] = true
		front := 0
		for front < len(qx) {
			cx, cy := qx[front], qy[front]
			front += 1
			for d := range len(dd) {
				nx, ny := cx+dd[d][0], cy+dd[d][1]
				if 0 <= nx && nx < n && 0 <= ny && ny < m && !seen[nx][ny] && ar[nx][ny] == ar[cx][cy]+1 {
					qx = append(qx, nx)
					qy = append(qy, ny)
					seen[nx][ny] = true
					if ar[nx][ny] == 9 {
						ans += 1
					}
				}
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day10) Part2(lines []string) string {
	sx, sy := make([]int, 0), make([]int, 0)
	ar := make([][]int, 0)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]int, len(line))
		for j, c := range line {
			row[j] = int(byte(c) - byte('0'))
			if row[j] == 0 {
				sx = append(sx, i)
				sy = append(sy, j)
			}
		}
		ar = append(ar, row)
	}
	n, m := len(ar), len(ar[0])
	dd := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	ans := 0
	seen, cc := make([][]bool, n), make([][]int, n)
	for i := range len(seen) {
		cc[i] = make([]int, m)
		seen[i] = make([]bool, m)
	}
	for i := range len(sx) {
		seen[sx[i]][sy[i]] = true
		cc[sx[i]][sy[i]] = 1
	}

	front := 0
	for front < len(sx) {
		cx, cy := sx[front], sy[front]
		front += 1
		if ar[cx][cy] == 9 {
			ans += cc[cx][cy]
		}
		for d := range len(dd) {
			nx, ny := cx+dd[d][0], cy+dd[d][1]
			if 0 <= nx && nx < n && 0 <= ny && ny < m && ar[nx][ny] == ar[cx][cy]+1 {
				if !seen[nx][ny] {
					sx = append(sx, nx)
					sy = append(sy, ny)
					seen[nx][ny] = true
				}

				cc[nx][ny] += cc[cx][cy]

			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day10) TestPart1() {
	ansExample1 := []string{"", "1", "36"}
	for i := 1; i < 3; i++ {
		input := util.ExampleInput(2024, 10, i)
		ans := p.Part1(input)
		if ans == fmt.Sprint(nil) {
		} else if ansExample1[i] == "" {
			fmt.Println("Correct answer Part1 missing, got", ans)
		} else if ans != ansExample1[i] {
			fmt.Println("Answer to Part1 incorrect", ans, ansExample1[i])
		} else {
			fmt.Println("Answer to Part1 correct", ans)
		}
	}

}

func (p Day10) TestPart2() {
	ansExample2 := []string{"", "16", "81"}
	for i := 1; i < 3; i++ {
		input := util.ExampleInput(2024, 10, i)
		ans := p.Part2(input)
		if ans == fmt.Sprint(nil) {
		} else if ansExample2[i] == "" {
			fmt.Println("Correct answer Part2 missing, got", ans)
		} else if ans != ansExample2[i] {
			fmt.Println("Answer to Part2 incorrect", ans, ansExample2[i])
		} else {
			fmt.Println("Answer to Part2 correct", ans)
		}

	}
}
