package year2024

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"aocgen/internal/util"
)

type Day07 struct{}

var dd1 []map[int]bool

func rec1(ar []int, target, i int) bool {
	if i == -1 {
		return target == 0
	}
	if target < 0 {
		return false
	}
	if ans, exists := dd1[i][target]; exists {
		return ans
	}
	ans := false
	ans = ans || rec1(ar, target-ar[i], i-1)
	if !ans && target%ar[i] == 0 {
		ans = ans || rec1(ar, target/ar[i], i-1)
	}
	dd1[i][target] = ans
	return ans
}

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
		dd1 = make([]map[int]bool, len(row))
		for i := range len(row) {
			dd1[i] = make(map[int]bool)
		}
		if rec1(row, target, len(row)-1) {
			ans += target
		}
	}
	return fmt.Sprint(ans)
}

var dd2 []map[int]bool

func p10(a int) int {
	return int(math.Pow10(len(strconv.Itoa(a))))
}

func p7rec2(ar []int, target, i int) bool {
	if i == -1 {
		return target == 0
	}
	if target < 0 {
		return false
	}
	if ans, exists := dd2[i][target]; exists {
		return ans
	}
	ans := false
	ans = ans || p7rec2(ar, target-ar[i], i-1)
	if !ans && target%ar[i] == 0 {
		ans = ans || p7rec2(ar, target/ar[i], i-1)
	}
	if !ans {
		e10 := p10(ar[i])
		if (target-ar[i])%e10 == 0 {
			ans = ans || p7rec2(ar, (target-ar[i])/e10, i-1)
		}
	}
	dd2[i][target] = ans
	return ans
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
		dd2 = make([]map[int]bool, len(row))
		for i := range len(row) {
			dd2[i] = make(map[int]bool)
		}
		if p7rec2(row, target, len(row)-1) {
			ans += target
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
