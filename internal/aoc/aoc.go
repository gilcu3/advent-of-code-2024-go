package aoc

import (
	"sort"

	"github.com/sirupsen/logrus"
)

type Puzzle interface {
	PartA([]string) string
	PartB([]string) string
	SampleA([]string)
	SampleB([]string)
}

var puzzles = map[int]map[int]Puzzle{}

func Register(year int, p map[int]Puzzle) {
	puzzles[year] = p
	for day := range puzzles[year] {
		logrus.Debugf("Registered %d: Day %d", year, day)
	}
}

func Years() []int {
	years := make([]int, 0)
	for y := range puzzles {
		if y > 0 {
			years = append(years, y)
		}
	}
	sort.Ints(years)
	return years
}

func Puzzles(year int) map[int]Puzzle {
	p, ok := puzzles[year]
	if !ok {
		logrus.Fatalf("Year not found: %d", year)
	}
	return p
}

func NewPuzzle(year, day int) Puzzle {
	puzzle, ok := puzzles[year][day]
	if !ok {
		logrus.Fatalf("Puzzle not found: %d-%d", year, day)
	}
	return puzzle
}

func Run(year, day int, p Puzzle, input []string) {
	if p == nil {
		logrus.Fatal("Failed to run empty puzzle")
		return
	}

	logrus.Infof("%d Day %d, Part A Result: %v", year, day, p.PartA(input))
	logrus.Infof("%d Day %d, Part B Result: %v", year, day, p.PartB(input))
}

func TestRun(year, day int, p Puzzle, input []string) {
	if p == nil {
		logrus.Fatal("Failed to run empty puzzle")
		return
	}
	p.SampleA(input)
	p.SampleB(input)
}

func RunYear(year int, testRun bool) {
	puzzles := Puzzles(year)
	for i := 1; i <= len(puzzles); i++ {
		RunDay(year, i, testRun)
	}
}

func RunDay(year, day int, testRun bool) {
	if testRun {
		TestRun(year, day, NewPuzzle(year, day), SampleInput(year, day))
	} else {
		Run(year, day, NewPuzzle(year, day), Input(year, day))
	}

}
