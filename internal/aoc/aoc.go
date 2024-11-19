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

func Run(year, day, part int, p Puzzle, input []string, submitRun bool) {
	if p == nil {
		logrus.Fatal("Failed to run empty puzzle")
		return
	}
	var ans string
	if part == 0 {
		ans = p.PartA(input)
		logrus.Infof("%d Day %d, Part A Result: %v", year, day, ans)
	} else if part == 1 {
		ans = p.PartB(input)
		logrus.Infof("%d Day %d, Part B Result: %v", year, day, ans)
	}
	if submitRun {
		Submit(year, day, part, ans)
	}
}

func TestRun(year, day, part int, p Puzzle, input []string) {
	if p == nil {
		logrus.Fatal("Failed to run empty puzzle")
		return
	}
	if part == 0 {
		p.SampleA(input)
	} else if part == 1 {
		p.SampleB(input)
	}

}

func RunDay(year, day, part int, testRun, submitRun bool) {
	if testRun {
		TestRun(year, day, part, NewPuzzle(year, day), SampleInput(year, day))
	} else {
		Run(year, day, part, NewPuzzle(year, day), Input(year, day), submitRun)
	}

}
