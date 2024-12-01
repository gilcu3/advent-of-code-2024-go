package year2024

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aocgen/internal/util"
)

type Day01 struct{}

func (p Day01) Part1(lines []string) string {

	var list1 []int
	var list2 []int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	var ans int = 0
	for i := 0; i < len(list1); i++ {
		if list1[i] < list2[i] {
			ans += list2[i] - list1[i]
		} else {
			ans += list1[i] - list2[i]
		}
	}

	return fmt.Sprint(ans)
}

func (p Day01) Part2(lines []string) string {
	var list1 []int
	var list2 []int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	var ans int = 0
	cc := make(map[int]int)
	for i := 0; i < len(list2); i++ {
		cc[list2[i]]++
	}
	for i := 0; i < len(list1); i++ {
		ans += list1[i] * cc[list1[i]]
	}
	return fmt.Sprint(ans)
}

func (p Day01) TestPart1() {
	const ansExample1 = "11"
	input := util.ExampleInput(2024, 01, 0)
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

func (p Day01) TestPart2() {
	const ansExample2 = "31"
	input := util.ExampleInput(2024, 01, 0)
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
