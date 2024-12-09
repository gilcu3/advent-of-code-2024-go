package year2024

import (
	"fmt"

	"aocgen/internal/util"
)

type Day09 struct{}

func s2(b, c int) int {
	return (2*b + c - 1) * c / 2
}

func (p Day09) Part1(lines []string) string {
	ar := make([]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		for _, c := range line {
			ar = append(ar, int(byte(c)-'0'))
		}
	}
	b, e := 0, len(ar)-1
	lb, le := ar[b], ar[e]
	ans := 0
	sb := 0
	for b <= e {
		if b == e {
			if b%2 == 0 {
				l := min(lb, le)
				ans += s2(sb, l) * b / 2
				lb, le = 0, 0
				sb += l
				b += 1
			} else {
				b += 1
			}

		} else if b%2 == 0 {
			ans += s2(sb, ar[b]) * b / 2
			sb += ar[b]
			b += 1
			lb = ar[b]
		} else if lb == 0 {
			b += 1
			lb = ar[b]
		} else if e%2 == 1 {
			e -= 1
			le = ar[e]
		} else if le == 0 {
			e -= 1
			le = ar[e]
		} else if le >= lb {
			le -= lb
			ans += s2(sb, lb) * e / 2
			sb += lb
			lb = 0
		} else {
			lb -= le
			ans += s2(sb, le) * e / 2
			sb += le
			le = 0
		}

	}
	return fmt.Sprint(ans)
}

func (p Day09) Part2(lines []string) string {
	ar := make([]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		for _, c := range line {
			ar = append(ar, int(byte(c)-'0'))
		}
	}

	ans := 0
	moved := make([]bool, len(ar))
	left := make([]int, len(ar))
	copy(left, ar)
	for e := len(ar) - 1; e > 0; e-- {
		if e%2 == 0 {
			sb := 0
			for b := range len(ar) {
				if b > e {
					break
				}
				if b%2 == 1 && left[b] >= ar[e] {
					ans += s2(sb+ar[b]-left[b], ar[e]) * e / 2
					moved[e] = true
					left[b] -= ar[e]
					break
				}
				sb += ar[b]
			}
		}
	}
	sb := 0
	for b := range len(ar) {
		if b%2 == 0 && !moved[b] {
			ans += s2(sb, ar[b]) * b / 2
		}
		sb += ar[b]
	}
	return fmt.Sprint(ans)
}

func (p Day09) TestPart1() {
	const ansExample1 = "1928"
	input := util.ExampleInput(2024, 9, 0)
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

func (p Day09) TestPart2() {
	const ansExample2 = "2858"
	input := util.ExampleInput(2024, 9, 0)
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
