package year2024

import (
	"fmt"

	"aocgen/internal/util"
)

type Day06 struct{}

func (p Day06) Part1(lines []string) string {
	ar := make([][]int, 0)
	sig := "^>v<.#"
	var sx, sy int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]int, len(line))
		for i, c := range line {
			for j, cc := range sig {
				if c == cc {
					row[i] = j
					break
				}
			}
			if row[i] < 4 {
				sx, sy = len(ar), i
			}
		}
		ar = append(ar, row)
	}
	n, m := len(ar), len(ar[0])
	ans := 0
	seen := make([][]bool, n)
	for i := 0; i < n; i++ {
		seen[i] = make([]bool, m)
	}
	dd := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for {
		if !seen[sx][sy] {
			ans += 1
			seen[sx][sy] = true
		}
		nsx, nsy := sx+dd[ar[sx][sy]][0], sy+dd[ar[sx][sy]][1]
		if nsx < n && nsx >= 0 && nsy < m && nsy >= 0 {
			if ar[nsx][nsy] == 4 {
				ar[nsx][nsy] = ar[sx][sy]
				ar[sx][sy] = 4
				sx, sy = nsx, nsy
			} else if ar[nsx][nsy] == 5 {
				ar[sx][sy] = (ar[sx][sy] + 1) % 4
			}
		} else {
			break
		}
	}
	return fmt.Sprint(ans)
}

func ccopy(ori [][]int) [][]int {
	cop := make([][]int, len(ori))
	for i := range ori {
		cop[i] = make([]int, len(ori[i]))
		copy(cop[i], ori[i])
	}
	return cop
}

func isLoop(ar [][]int, sx, sy int) bool {
	n, m := len(ar), len(ar[0])
	seen := make([][]int, n)
	for i := range seen {
		seen[i] = make([]int, m)
	}
	dd := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for {
		if seen[sx][sy]&(1<<ar[sx][sy]) != 0 {
			return true
		}
		seen[sx][sy] ^= (1 << ar[sx][sy])
		nsx, nsy := sx+dd[ar[sx][sy]][0], sy+dd[ar[sx][sy]][1]
		if nsx < n && nsx >= 0 && nsy < m && nsy >= 0 {
			if ar[nsx][nsy] == 4 {
				ar[nsx][nsy] = ar[sx][sy]
				ar[sx][sy] = 4
				sx, sy = nsx, nsy
			} else if ar[nsx][nsy] == 5 {
				ar[sx][sy] = (ar[sx][sy] + 1) % 4
			} else {
				fmt.Println("impossible")
			}

		} else {
			break
		}
	}
	return false
}

func (p Day06) Part2(lines []string) string {
	ar := make([][]int, 0)
	sig := "^>v<.#"
	var sx, sy int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]int, len(line))
		for i, c := range line {
			for j, cc := range sig {
				if c == cc {
					row[i] = j
					break
				}
			}
			if row[i] < 4 {
				sx, sy = len(ar), i
			}
		}
		ar = append(ar, row)
	}
	n, m := len(ar), len(ar[0])
	ans := 0
	dd := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	seen := make([][]bool, n)
	for i := 0; i < n; i++ {
		seen[i] = make([]bool, m)
	}
	for {
		seen[sx][sy] = true
		nsx, nsy := sx+dd[ar[sx][sy]][0], sy+dd[ar[sx][sy]][1]
		if nsx < n && nsx >= 0 && nsy < m && nsy >= 0 {
			if ar[nsx][nsy] == 4 {
				ar[nsx][nsy] = 5
				if !seen[nsx][nsy] && isLoop(ccopy(ar), sx, sy) {
					ans += 1
				}
				ar[nsx][nsy] = ar[sx][sy]
				ar[sx][sy] = 4
				sx, sy = nsx, nsy
			} else if ar[nsx][nsy] == 5 {
				ar[sx][sy] = (ar[sx][sy] + 1) % 4
			}
		} else {
			break
		}
	}
	return fmt.Sprint(ans)
}

func (p Day06) TestPart1() {
	const ansExample1 = "41"
	input := util.ExampleInput(2024, 06, 0)
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

func (p Day06) TestPart2() {
	const ansExample2 = "6"
	input := util.ExampleInput(2024, 06, 0)
	ans := p.Part2(input)
	if ans == fmt.Sprint(nil) {
	} else if ansExample2 == "" {
		fmt.Println("Correct answer Part2 missing, got", ans)
	} else if ans != ansExample2 {
		fmt.Println("Answer to Part2 incorrect", ans, ansExample2)
	} else {
		fmt.Println("Answer to Part2 correct", ans)
	}
}
