package year2024

import (
	"fmt"
	"strings"

	"aocgen/internal/util"
)

type Day19 struct{}

func (p Day19) Part1(lines []string) string {
	sec := 0
	ans := 0
	words := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			sec += 1
			continue
		}
		if sec == 0 {
			words = strings.Split(line, ", ")
		} else {
			n := len(line)
			mem := make([]bool, len(line)+1)
			mem[n] = true
			for i := n - 1; i >= 0; i-- {
				for j := range words {
					if len(words[j])+i <= n && mem[i+len(words[j])] {
						pos := true
						for k := range words[j] {
							if words[j][k] != line[i+k] {
								pos = false
								break
							}
						}
						if pos {
							mem[i] = true
							break
						}
					}
				}
			}
			if mem[0] {
				ans += 1
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day19) Part2(lines []string) string {
	sec := 0
	ans := 0
	words := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			sec += 1
			continue
		}
		if sec == 0 {
			words = strings.Split(line, ", ")
		} else {
			n := len(line)
			mem := make([]int, len(line)+1)
			mem[n] = 1
			for i := n - 1; i >= 0; i-- {
				for j := range words {
					if len(words[j])+i <= n && mem[i+len(words[j])] > 0 {
						pos := true
						for k := range words[j] {
							if words[j][k] != line[i+k] {
								pos = false
								break
							}
						}
						if pos {
							mem[i] += mem[i+len(words[j])]
						}
					}
				}
			}
			ans += mem[0]
		}
	}
	return fmt.Sprint(ans)
}

func (p Day19) TestPart1() {
	const ansExample1 = "6"
	input := util.ExampleInput(2024, 19, 0)
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

func (p Day19) TestPart2() {
	const ansExample2 = "16"
	input := util.ExampleInput(2024, 19, 0)
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
