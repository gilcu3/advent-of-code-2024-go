package year2024

import (
	"fmt"
	"strconv"
	"strings"

	"aocgen/internal/util"
)

type Day02 struct{}

func absInt(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func Sign(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}

func (p Day02) Part1(lines []string) string {
	ans := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		ar := make([]int, 0, len(parts))

		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			ar = append(ar, num)
		}
		pos := true
		for i := 0; i < len(ar); i++ {
			if i > 0 {
				d := absInt(ar[i] - ar[i-1])
				if d == 0 || d > 3 {
					pos = false
					break
				}
			}
			if i > 1 {
				if Sign(ar[i]-ar[i-1]) != Sign(ar[i-1]-ar[i-2]) {
					pos = false
					break
				}
			}
		}
		if pos {
			ans += 1
		}

	}
	return fmt.Sprint(ans)
}

func (p Day02) Part2(lines []string) string {
	ans := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		ar := make([]int, 0, len(parts))

		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			ar = append(ar, num)
		}
		pos := true

		for i := 0; i < len(ar); i++ {
			if i > 0 {
				d := absInt(ar[i] - ar[i-1])
				if d == 0 || d > 3 {
					pos = false
					break
				}
			}
			if i > 1 {
				if Sign(ar[i]-ar[i-1]) != Sign(ar[i-1]-ar[i-2]) {
					pos = false
					break
				}
			}
		}
		if pos {
			ans += 1
		} else {
			for i := 0; i < len(ar); i++ {
				ar1 := make([]int, 0, len(parts)-1)
				for j := 0; j < len(ar); j++ {
					if j != i {
						ar1 = append(ar1, ar[j])
					}
				}
				pos = true
				for j := 0; j < len(ar1); j++ {
					if j > 0 {
						d := absInt(ar1[j] - ar1[j-1])
						if d == 0 || d > 3 {
							pos = false
							break
						}
					}
					if j > 1 {
						if Sign(ar1[j]-ar1[j-1]) != Sign(ar1[j-1]-ar1[j-2]) {
							pos = false
							break
						}
					}
				}
				if pos {
					ans += 1
					break
				}
			}

		}
	}
	return fmt.Sprint(ans)
}

func (p Day02) TestPart1() {
	const ansExample1 = "2"
	input := util.ExampleInput(2024, 02, 0)
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

func (p Day02) TestPart2() {
	const ansExample2 = "4"
	input := util.ExampleInput(2024, 02, 0)
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
