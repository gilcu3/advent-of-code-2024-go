package year2024

import (
	"fmt"

	"aocgen/internal/util"
)

type Day12 struct{}

func (p Day12) Part1(lines []string) string {
	ar := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		ar = append(ar, line)
	}
	n, m := len(ar), len(ar[0])
	ans := 0
	seen := make([][]bool, n)
	for i := range n {
		seen[i] = make([]bool, m)
	}
	dd := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for si := range n {
		for sj := range m {
			if !seen[si][sj] {
				qx, qy := make([]int, 0), make([]int, 0)
				front := 0
				seen[si][sj] = true
				qx, qy = append(qx, si), append(qy, sj)
				per := 0
				for front < len(qx) {
					cx, cy := qx[front], qy[front]
					front += 1
					for d := range 4 {
						nx, ny := cx+dd[d][0], cy+dd[d][1]
						if 0 <= nx && nx < n && 0 <= ny && ny < m && ar[nx][ny] == ar[si][sj] {
							if !seen[nx][ny] {
								qx, qy = append(qx, nx), append(qy, ny)
								seen[nx][ny] = true
							}
						} else {
							per += 1
						}
					}
				}
				ans += front * per
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day12) Part2(lines []string) string {

	ar := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		ar = append(ar, line)
	}
	n, m := len(ar), len(ar[0])
	ans := 0
	seen := make([][]bool, n)
	for i := range n {
		seen[i] = make([]bool, m)
	}
	dd := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for si := range n {
		for sj := range m {
			if !seen[si][sj] {
				qx, qy := make([]int, 0), make([]int, 0)
				front := 0
				seen[si][sj] = true
				qx, qy = append(qx, si), append(qy, sj)
				per := 0
				for front < len(qx) {
					cx, cy := qx[front], qy[front]
					front += 1
					for d := range 4 {
						nx, ny := cx+dd[d][0], cy+dd[d][1]
						if 0 <= nx && nx < n && 0 <= ny && ny < m && ar[nx][ny] == ar[cx][cy] {
							if !seen[nx][ny] {
								qx, qy = append(qx, nx), append(qy, ny)
								seen[nx][ny] = true
							}
						} else if d%2 == 0 { // horizontal
							if ny == 0 || ar[cx][ny-1] != ar[cx][cy] {
								per += 1
							} else if ny > 0 {
								if 0 <= nx && nx < n && 0 <= ny && ny < m {
									if ar[nx][ny-1] == ar[cx][cy] {
										per += 1
									}
								} else if ar[cx][ny-1] != ar[cx][cy] {
									per += 1
								}
							}
						} else { // vertical
							if nx == 0 || ar[nx-1][cy] != ar[cx][cy] {
								per += 1
							} else if nx > 0 {
								if 0 <= nx && nx < n && 0 <= ny && ny < m {
									if ar[nx-1][ny] == ar[cx][cy] {
										per += 1
									}
								} else if ar[nx-1][cy] != ar[cx][cy] {
									per += 1
								}
							}
						}
					}
				}
				ans += front * per
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day12) TestPart1() {
	ansExample1 := []string{"", "140", "772", "1930"}
	for i := 1; i < 4; i++ {
		input := util.ExampleInput(2024, 12, i)
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

func (p Day12) TestPart2() {
	ansExample2 := []string{"", "80", "436", "1206"}
	for i := 1; i < 4; i++ {
		input := util.ExampleInput(2024, 12, i)
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
