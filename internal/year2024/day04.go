package year2024

import (
	"fmt"

	"aocgen/internal/util"
)

type Day04 struct{}

func (p Day04) Part1(lines []string) string {
	n, m := len(lines), len(lines[0])
	ar := make([][]int, n)

	M := make(map[byte]int)
	s := "XMAS"
	for i := range len(s) {
		M[s[i]] = i
	}

	for i, line := range lines {
		if len(line) == 0 {
			n -= 1
			continue
		}
		ar[i] = make([]int, m)
		for j := range m {
			ar[i][j] = M[line[j]]
		}
	}
	ans := 0
	for i := range n {
		for j := range m {
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if di != 0 || dj != 0 {
						pos := true
						ii, jj := i, j
						for t := range 4 {
							if !(ii < n && ii >= 0 && jj < m && jj >= 0 && ar[ii][jj] == t) {
								pos = false
								break
							}
							ii += di
							jj += dj
						}
						if pos {
							ans += 1
						}
					}
				}
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day04) Part2(lines []string) string {
	n, m := len(lines), len(lines[0])
	ar := make([][]int, n)

	M := make(map[byte]int)
	s := "XMAS"
	for i := range len(s) {
		M[s[i]] = i
	}

	for i, line := range lines {
		if len(line) == 0 {
			n -= 1
			continue
		}
		ar[i] = make([]int, m)
		for j := range m {
			ar[i][j] = M[line[j]]
		}
	}
	var pat [3][3]int
	pat[0] = [3]int{1, 0, 3}
	pat[1] = [3]int{0, 2, 0}
	pat[2] = [3]int{1, 0, 3}
	ans := 0
	for i := range n {
		for j := range m {

			for dj := -1; dj <= 1; dj += 2 {
				di := 1
				pos := true
				ii := i
				for x := range 3 {
					jj := j
					for y := range 3 {
						if !(ii < n && ii >= 0 && jj < m && jj >= 0 && (ar[ii][jj] == pat[x][y] || pat[x][y] == 0)) {
							pos = false
							break
						}
						jj += dj
					}
					if !pos {
						break
					}
					ii += di
				}

				if pos {
					ans += 1
				}
			}
			for di := -1; di <= 1; di += 2 {
				dj := 1
				pos := true
				ii := i
				for x := range 3 {
					jj := j
					for y := range 3 {
						if !(ii < n && ii >= 0 && jj < m && jj >= 0 && (ar[ii][jj] == pat[y][x] || pat[y][x] == 0)) {
							pos = false
							break
						}
						jj += dj
					}
					if !pos {
						break
					}
					ii += di
				}

				if pos {
					ans += 1
				}
			}

		}
	}
	return fmt.Sprint(ans)
}

func (p Day04) TestPart1() {
	const ansExample1 = "18"
	input := util.ExampleInput(2024, 04, 0)
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

func (p Day04) TestPart2() {
	const ansExample2 = "9"
	input := util.ExampleInput(2024, 04, 0)
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
