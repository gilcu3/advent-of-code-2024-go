package year2024

import (
	"fmt"
	"regexp"
	"strconv"

	"aocgen/internal/util"
)

type Day03 struct{}

func (p Day03) Part1(lines []string) string {
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)
	ans := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) == 3 {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])

				ans += x * y
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day03) Part2(lines []string) string {
	pattern := `mul\((\d+),(\d+)\)|don\'t\(\)|do\(\)`
	re := regexp.MustCompile(pattern)
	ans := 0
	enabled := true
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "don't()" {

				enabled = false
			} else if match[0] == "do()" {
				enabled = true
			} else if len(match) == 3 && enabled {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])

				ans += x * y
			}
		}
	}
	return fmt.Sprint(ans)
}

func (p Day03) TestPart1() {
	const ansExample1 = "161"
	input := util.ExampleInput(2024, 03, 1)
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

func (p Day03) TestPart2() {
	const ansExample2 = "48"
	input := util.ExampleInput(2024, 03, 2)
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
