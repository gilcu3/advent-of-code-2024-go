package year2024

import (
	"fmt"
	"strconv"
	"strings"

	"aocgen/internal/util"
)

type Day07 struct{}

func (p Day07) Part1(lines []string) string {
	ans := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		row := make([]int, 0)
		target := 0
		for i, s := range parts {
			if i == 0 {
				target, _ = strconv.Atoi(s[:len(s)-1])
			} else {
				cur, _ := strconv.Atoi(s)
				row = append(row, cur)
			}
		}
		for m := range 1 << (len(row) - 1) {
			cur := row[0]
			for i := 1; i < len(row); i++ {
				if (1<<(i-1))&m != 0 {
					cur += row[i]
				} else {
					cur *= row[i]
				}
			}
			if cur == target {
				ans += target
				break
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day07) Part2(lines []string) string {
	ans := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		row := make([]int, 0)
		target := 0
		for i, s := range parts {
			if i == 0 {
				target, _ = strconv.Atoi(s[:len(s)-1])
			} else {
				cur, _ := strconv.Atoi(s)
				row = append(row, cur)
			}
		}
		solved := false
		for m := range 1 << (len(row) - 1) {
			cur := row[0]
			for i := 1; i < len(row); i++ {
				if (1<<(i-1))&m != 0 {
					cur += row[i]
				} else {
					cur *= row[i]
				}
			}
			if cur == target {
				solved = true
				ans += target
				break
			}
		}
		if !solved {
			p3 := make([]int, len(row)+1)
			p3[0] = 1
			for i := range len(row) {
				p3[i+1] = p3[i] * 3
			}
			s10 := make([]int, len(row))
			for i := range len(row) {
				s10[i] = len(strconv.Itoa(row[i]))
			}
			p10 := make([]int, 10)
			p10[0] = 1
			for i := range 9 {
				p10[i+1] = p10[i] * 10
			}
			for m := range p3[len(row)-1] {
				cur := row[0]
				for i := 1; i < len(row); i++ {
					t := (m / p3[i-1]) % 3
					if t == 0 {
						cur += row[i]
					} else if t == 1 {
						cur *= row[i]
					} else {
						cur = cur*p10[s10[i]] + row[i]
					}
				}
				if cur == target {
					solved = true
					ans += target
					break
				}
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day07) TestPart1() {
	const ansExample1 = "3749"
	input := util.ExampleInput(2024, 07, 0)
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

func (p Day07) TestPart2() {
	const ansExample2 = "11387"
	input := util.ExampleInput(2024, 07, 0)
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
