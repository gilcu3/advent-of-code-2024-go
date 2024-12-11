package year2024

import (
	"fmt"
	"strconv"
	"strings"

	"aocgen/internal/util"
)

type Day11 struct{}

func rec(n, reps int) int {
	if reps == 0 {
		return 1
	} else if n == 0 {
		return rec(1, reps-1)
	} else {
		ns := strconv.Itoa(n)
		if len(ns)%2 == 0 {
			n1, _ := strconv.Atoi(ns[:len(ns)/2])
			n2, _ := strconv.Atoi(ns[len(ns)/2:])
			return rec(n1, reps-1) + rec(n2, reps-1)
		} else {
			return rec(n*2024, reps-1)
		}

	}
}

func (p Day11) Part1(lines []string) string {
	reps := 25
	ar := make([]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		for _, p := range parts {
			n, _ := strconv.Atoi(p)
			ar = append(ar, n)
		}
	}
	ans := 0
	for _, n := range ar {
		ans += rec(n, reps)
	}
	return fmt.Sprint(ans)
}

var results = make([]map[int]int, 0)

func rec2(n, reps int) int {

	if reps == -1 {
		return 1
	} else if n == 0 {
		return rec2(1, reps-1)
	} else {
		if v, exists := results[reps][n]; exists {
			return v
		}
		ans := 0
		ns := strconv.Itoa(n)
		if len(ns)%2 == 0 {
			n1, _ := strconv.Atoi(ns[:len(ns)/2])
			n2, _ := strconv.Atoi(ns[len(ns)/2:])
			ans = rec2(n1, reps-1) + rec2(n2, reps-1)
		} else {
			ans = rec2(n*2024, reps-1)
		}
		results[reps][n] = ans
		return ans
	}
}

func (p Day11) Part2(lines []string) string {
	reps := 75
	for range reps {
		results = append(results, make(map[int]int))
	}
	ar := make([]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		for _, p := range parts {
			n, _ := strconv.Atoi(p)
			ar = append(ar, n)
		}
	}
	ans := 0
	for _, n := range ar {
		ans += rec2(n, reps-1)
	}
	return fmt.Sprint(ans)
}

func (p Day11) TestPart1() {
	const ansExample1 = "55312"
	input := util.ExampleInput(2024, 11, 0)
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

func (p Day11) TestPart2() {
	const ansExample2 = "65601038650482"
	input := util.ExampleInput(2024, 11, 0)
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
