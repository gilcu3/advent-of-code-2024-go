package year2024

import (
	"fmt"

	"aocgen/internal/util"
)

type Day15 struct{}

func (p Day15) Part1(lines []string) string {
	sig := "^>v<.#O@"
	dd := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	ar := make([][]int, 0)
	var sx, sy int
	sec := 0
	moves := make([]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			sec += 1
			continue
		}
		if sec == 0 {
			row := make([]int, len(line))
			for i, c := range line {
				for j, cc := range sig {
					if c == cc {
						row[i] = j
						break
					}
				}
				if row[i] == 7 {
					sx, sy = len(ar), i
					row[sy] = 4
				}
			}
			ar = append(ar, row)
		} else {
			for _, c := range line {
				var m int
				for j, cc := range sig {
					if c == cc {
						m = j
						break
					}
				}
				moves = append(moves, m)
			}

		}

	}
	n, m := len(ar), len(ar[0])

	for _, d := range moves {
		nx, ny := sx+dd[d][0], sy+dd[d][1]
		if ar[nx][ny] == 4 {
			sx, sy = nx, ny
		} else if ar[nx][ny] == 6 {
			for ar[nx][ny] == 6 {
				nx, ny = nx+dd[d][0], ny+dd[d][1]
			}
			if ar[nx][ny] == 4 {
				ar[nx][ny] = 6
				sx, sy = sx+dd[d][0], sy+dd[d][1]
				ar[sx][sy] = 4
			}
		}
	}
	ans := 0
	for i := range n {
		for j := range m {
			if ar[i][j] == 6 {
				ans += 100*i + j
			}
		}
	}
	return fmt.Sprint(ans)
}

// func p15render(sig string, ar [][]int, n, m, sx, sy int) {
// 	for i := range n {
// 		for j := range m {
// 			if ar[i][j] == 6 {
// 				fmt.Printf("[")
// 			} else if j > 0 && ar[i][j-1] == 6 {
// 				fmt.Printf("]")
// 			} else if i == sx && j == sy {
// 				fmt.Printf("@")
// 			} else {
// 				fmt.Printf("%c", sig[ar[i][j]])
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

func (p Day15) Part2(lines []string) string {
	sig := "^>v<.#O@"
	dd := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	ar := make([][]int, 0)
	var sx, sy int
	sec := 0
	moves := make([]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			sec += 1
			continue
		}
		if sec == 0 {
			row := make([]int, 2*len(line))
			for i, c := range line {
				for j, cc := range sig {
					if c == cc {
						row[2*i] = j
						if j == 5 {
							row[2*i+1] = 5
						} else {
							row[2*i+1] = 4
						}

						break
					}
				}
				if row[2*i] == 7 {
					sx, sy = len(ar), 2*i
					row[sy] = 4
				}
			}
			ar = append(ar, row)
		} else {
			for _, c := range line {
				var m int
				for j, cc := range sig {
					if c == cc {
						m = j
						break
					}
				}
				moves = append(moves, m)
			}

		}

	}
	n, m := len(ar), len(ar[0])

	for _, d := range moves {
		// p15render(sig, ar, n, m, sx, sy)
		qx, qy := []int{sx}, []int{sy}
		front := 0
		willmove := true
		for front < len(qx) && willmove {
			cx, cy := qx[front], qy[front]
			front += 1
			if d%2 == 0 {
				nx := cx + dd[d][0]
				if nx >= 0 {
					if ar[nx][cy] == 5 || (cy < m-1 && ar[nx][cy+1] == 5 && front != 1) {
						willmove = false
					} else {
						if cy > 0 && ar[nx][cy-1] == 6 && (qx[len(qx)-1] != nx || qy[len(qy)-1] != cy-1) {
							qx, qy = append(qx, nx), append(qy, cy-1)
						}
						if ar[nx][cy] == 6 {
							qx, qy = append(qx, nx), append(qy, cy)
						}
						if cy < m-1 && ar[nx][cy+1] == 6 && (cx != sx || cy != sy) {
							qx, qy = append(qx, nx), append(qy, cy+1)
						}

					}
				} else {
					willmove = false
				}
			} else if d == 3 {
				ny := cy - 2
				if ny >= 0 {
					if ar[cx][ny+1] == 5 {
						willmove = false
					} else if ar[cx][ny] == 6 {
						qx, qy = append(qx, cx), append(qy, ny)
					}
				} else {
					willmove = false
				}
			} else {
				ny := cy + 2
				if front == 1 {
					ny = cy + 1
				}
				if ny < m {
					if ar[cx][ny] == 5 {
						willmove = false
					} else if ar[cx][ny] == 6 {
						qx, qy = append(qx, cx), append(qy, ny)
					}
				} else {
					willmove = false
				}
			}
		}
		// fmt.Printf("%c %d \n", sig[d], len(qx))
		if willmove {
			for i := len(qx) - 1; i > 0; i-- {
				cx, cy := qx[i], qy[i]
				nx, ny := cx+dd[d][0], cy+dd[d][1]
				ar[cx][cy] = 4
				ar[nx][ny] = 6
			}
			sx, sy = sx+dd[d][0], sy+dd[d][1]
		}

	}
	// p15render(sig, ar, n, m, sx, sy)
	ans := 0
	for i := range n {
		for j := range m {
			if ar[i][j] == 6 {
				ans += 100*i + j
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day15) TestPart1() {
	const ansExample1 = "2028"
	input := util.ExampleInput(2024, 15, 1)
	ans := p.Part1(input)
	if ans == fmt.Sprint(nil) {
	} else if ansExample1 == "" {
		fmt.Println("Correct answer Part1 missing, got", ans)
	} else if ans != ansExample1 {
		fmt.Println("Answer to Part1 incorrect", ans, ansExample1)
	} else {
		fmt.Println("Answer to Part1 correct", ans)
	}
}

func (p Day15) TestPart2() {
	ansExample2 := []string{"", "1751", "618", "9021"}
	for i := 1; i < 4; i++ {
		input := util.ExampleInput(2024, 15, i)
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
