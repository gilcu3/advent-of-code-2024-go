package year2024

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aocgen/internal/util"
)

type Day05 struct{}

func (p Day05) Part1(lines []string) string {
	sec := 0
	M := make([][]int, 0)
	ans := 0
	for _, line := range lines {
		if len(line) == 0 {
			sec += 1
			continue
		}
		if sec == 0 {
			rule := strings.Split(line, "|")
			m1, _ := strconv.Atoi(rule[0])
			m2, _ := strconv.Atoi(rule[1])
			M = append(M, []int{m1, m2})
		} else if sec == 1 {
			pp := strings.Split(line, ",")
			ppn := make([]int, 0)
			for _, p := range pp {
				pn, _ := strconv.Atoi(p)
				ppn = append(ppn, pn)
			}
			pos := true
			for i, p1 := range ppn {
				for _, p2 := range ppn[i+1:] {
					for _, r := range M {
						if r[0] == p2 && r[1] == p1 {
							pos = false
							break
						}
					}
					if !pos {
						break
					}
				}
				if !pos {
					break
				}
			}
			if pos {
				ans += ppn[len(ppn)/2]
			}
		}
	}
	return fmt.Sprint(ans)
}

var M2 = [][]int{}

type CustomInt []int

func (w CustomInt) Len() int {
	return len(w)
}

func (w CustomInt) Less(i, j int) bool {
	for _, r := range M2 {
		if r[0] == w[i] && r[1] == w[j] {
			return true
		}
	}
	return false
}

func (w CustomInt) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (p Day05) Part2(lines []string) string {
	sec := 0
	ans := 0
	for _, line := range lines {
		if len(line) == 0 {
			sec += 1
			continue
		}
		if sec == 0 {
			rule := strings.Split(line, "|")
			m1, _ := strconv.Atoi(rule[0])
			m2, _ := strconv.Atoi(rule[1])
			M2 = append(M2, []int{m1, m2})
		} else if sec == 1 {
			pp := strings.Split(line, ",")
			ppn := make([]int, 0)
			for _, p := range pp {
				pn, _ := strconv.Atoi(p)
				ppn = append(ppn, pn)
			}
			pos := true
			for i, p1 := range ppn {
				for _, p2 := range ppn[i+1:] {
					for _, r := range M2 {
						if r[0] == p2 && r[1] == p1 {
							pos = false
							break
						}
					}
					if !pos {
						break
					}
				}
				if !pos {
					break
				}
			}
			if !pos {
				sort.Sort(CustomInt(ppn))
				ans += ppn[len(ppn)/2]
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day05) TestPart1() {
	const ansExample1 = "143"
	input := util.ExampleInput(2024, 05, 0)
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

func (p Day05) TestPart2() {
	const ansExample2 = "123"
	input := util.ExampleInput(2024, 05, 0)
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
