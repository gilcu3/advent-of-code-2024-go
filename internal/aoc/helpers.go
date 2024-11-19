package aoc

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func Year() int {
	year, _, _ := time.Now().Date()
	return year
}

func Input(year, day int) []string {
	fileName := fmt.Sprintf("./internal/aoc/year%d/input/day%s.in", year, FormatDay(day))
	return readFile(fileName)
}

func SampleInput(year, day int) []string {
	fileName := fmt.Sprintf("./internal/aoc/year%d/sample/day%s.in", year, FormatDay(day))
	return readFile(fileName)
}

func TestInput(year, day int) []string {
	fileName := fmt.Sprintf("../year%d/input/day%s.in", year, FormatDay(day))
	return readFile(fileName)
}

func readFile(fileName string) []string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		logrus.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	return lines
}
