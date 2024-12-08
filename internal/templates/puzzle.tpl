// Code generated by aocgen; DO NOT EDIT.
package year{{.Year}}

import (
    "fmt"

    "aocgen/internal/util"
)

type Day{{.Day}} struct{}


func (p Day{{.Day}}) Part1(lines []string) string {
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		
	}
    return fmt.Sprint(nil)
}



func (p Day{{.Day}}) Part2(lines []string) string {
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		
	}
    return fmt.Sprint(nil)
}


func (p Day{{.Day}}) TestPart1() {
    const ansExample1 = ""
    input := util.ExampleInput({{.Year}}, {{.UDay}}, 0)
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

func (p Day{{.Day}}) TestPart2() {
    const ansExample2 = ""
    input := util.ExampleInput({{.Year}}, {{.UDay}}, 0)
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