package year2024

import (
	"fmt"

	"aocgen/internal/util"
)

type Day25 struct{}

func (p Day25) Part1(lines []string) string {
	locks, keys := [][]int{}, [][]int{}
	cur := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		cur = append(cur, line)
		if len(cur) == 7 {
			islock := cur[0][0] == '#'
			cols := make([]int, 5)
			for i := range 7 {
				for j := range 5 {
					if cur[i][j] == '#' {
						cols[j] += 1
					}
				}
			}
			// fmt.Println(cols)
			if islock {
				locks = append(locks, cols)
			} else {
				keys = append(keys, cols)
			}
			cur = []string{}
		}
	}
	ans := 0
	n, m := len(locks), len(keys)
	for i := range n {
		for j := range m {
			pos := true
			for k := range 5 {
				if locks[i][k]+keys[j][k] > 7 {
					pos = false
					break
				}
			}
			if pos {
				// fmt.Println(i, j, locks[i], keys[j])
				ans += 1
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day25) Part2(lines []string) string {
	return fmt.Sprint(nil)
}

func (p Day25) TestPart1() {
	const ansExample1 = "3"
	input := util.ExampleInput(2024, 25, 0)
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

func (p Day25) TestPart2() {
	const ansExample2 = ""
	input := util.ExampleInput(2024, 25, 0)
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
