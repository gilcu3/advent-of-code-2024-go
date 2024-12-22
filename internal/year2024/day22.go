package year2024

import (
	"fmt"
	"strconv"

	"aocgen/internal/util"
)

type Day22 struct{}

func (p Day22) Part1(lines []string) string {
	k, ans, mod := 2000, 0, 16777216
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		s, _ := strconv.Atoi(line)
		for range k {
			s = ((s * 64) ^ s) % mod
			s = ((s / 32) ^ s) % mod
			s = ((s * 2048) ^ s) % mod
		}
		ans += s

	}
	return fmt.Sprint(ans)
}

func (p Day22) Part2(lines []string) string {
	k, ans, mod := 2000, 0, 16777216
	mx := 19 * 19 * 19 * 19
	cdp := make([]int, mx)
	dp := make([]int, mx)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		s, _ := strconv.Atoi(line)
		c := 0
		for j := range k {
			s0 := s
			s = ((s * 64) ^ s) % mod
			s = ((s / 32) ^ s) % mod
			s = ((s * 2048) ^ s) % mod
			c = (c*19 + s0%10 - s%10 + 9) % mx
			if j >= 3 {
				if cdp[c] != i+1 {
					cdp[c] = i + 1
					dp[c] += s % 10
				}
			}
		}
	}
	for _, v := range dp {
		if ans < v {
			ans = v
		}
	}
	return fmt.Sprint(ans)
}

func (p Day22) TestPart1() {
	const ansExample1 = "37327623"
	input := util.ExampleInput(2024, 22, 0)
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

func (p Day22) TestPart2() {
	const ansExample2 = "24"
	input := util.ExampleInput(2024, 22, 0)
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
