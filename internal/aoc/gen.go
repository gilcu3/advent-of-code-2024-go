package aoc

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dolmen-go/codegen"
	"github.com/sirupsen/logrus"
)

const aoc_path = "internal/aoc"
const yearsFile = aoc_path + "/" + "years.go"

var validPuzzleFile = regexp.MustCompile(`^day[0-3][0-9]$`)

func puzzlePath(year int) string {
	return fmt.Sprintf("%s/year%d", aoc_path, year)
}

func puzzleFileName(year, day int) string {
	path := puzzlePath(year)
	return fmt.Sprintf("%s/day%s.go", path, FormatDay(day))
}

const puzzleTemplate = `// Code generated by aocgen; DO NOT EDIT.
	package year{{.Year}}

	type Day{{.Day}} struct{}
	
	func (p Day{{.Day}}) PartA(lines []string) any {
		return nil
	}
	
	func (p Day{{.Day}}) PartB(lines []string) any {
		return nil
	}
	
`

func NewPuzzleFile(year, day int) {
	fileName := puzzleFileName(year, day)
	if _, err := os.Stat(fileName); err == nil || !errors.Is(err, os.ErrNotExist) {
		logrus.Infof("Puzzle file already exists: %s", fileName)
		return
	}

	tmpl := codegen.MustParse(puzzleTemplate)
	if err := tmpl.CreateFile(fileName, map[string]any{
		"Year": year,
		"Day":  FormatDay(day),
	}); err != nil {
		logrus.Fatal(err)
	}

	RemoveFirstLine(fileName)

	logrus.Infof("Created file: %s", fileName)
}

func InitializePackage(year int) {
	path := puzzlePath(year)
	if err := CreateDirectory(path); err != nil {
		logrus.Fatal(err)
	}
}

const initMainTemplate = `// Code generated by aocgen; DO NOT EDIT.
	package aoc

	import (
		{{.Imports}}
	)

	func RegisterYears() {
		{{.Inits}}
	}
`

func findYears() []int {

	dirs, err := os.ReadDir(aoc_path)
	if err != nil {
		logrus.Fatal(err)
		return []int{}
	}
	years := []int{}

	for _, dir := range dirs {
		if !dir.IsDir() || len(dir.Name()) < 8 || strings.HasSuffix(dir.Name(), "tests") || dir.Name()[:4] != "year" {
			continue
		}
		year, _ := strconv.Atoi(dir.Name()[4:])
		logrus.Infof("Found directory to include in codegen: %d", year)
		years = append(years, year)
	}
	return years
}

func findDays(year int) []int {
	pathYear := fmt.Sprintf("%s/year%d", aoc_path, year)
	files, err := os.ReadDir(pathYear)
	if err != nil {
		return []int{}
	}
	days := []int{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		puzzleName := strings.Split(file.Name(), ".")[0]
		if !validPuzzleFile.Match([]byte(puzzleName)) {
			continue
		}
		day, _ := strconv.Atoi(puzzleName[3:])
		days = append(days, day)
	}
	return days
}

func UpdateYearsFile() {

	var imports, inits string
	for _, year := range findYears() {

		puzzles := ""
		days := findDays(year)
		if len(days) == 0 {
			continue
		}
		imports += fmt.Sprintf("\"aocgen/%s/year%d\"\n", aoc_path, year)
		for _, day := range days {
			correctPuzzleName := fmt.Sprintf("Day%s", FormatDay(day))
			puzzles += fmt.Sprintf("%d: year%d.%s{},\n", day, year, correctPuzzleName)
			logrus.Debugf("Found puzzle file for %d-%d", year, day)
		}
		inits += fmt.Sprintf("Register(%d, map[int]Puzzle{%s})", year, puzzles)
	}

	tmpl := codegen.MustParse(initMainTemplate)

	err := os.Remove(yearsFile)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Deleted file: %s", yearsFile)

	if err := tmpl.CreateFile(yearsFile, map[string]interface{}{
		"Imports": imports,
		"Inits":   inits,
	}); err != nil {
		logrus.Fatal(err)
	}
}

const benchmarkingTemplate = `// Code generated by aocgen; DO NOT EDIT.
	package tests

	import (
		"testing"
		"aocgen/internal/aoc"
	)

	{{.Benchmarks}}
`

func UpdateBenchmarks(year int) {
	pathTests := fmt.Sprintf("%s/tests", aoc_path)
	fileName := fmt.Sprintf("%s/year%d_test.go", pathTests, year)

	benchmarks := ""

	CreateDirectory(pathTests)
	days := findDays(year)
	if len(days) == 0 {
		os.Remove(fileName)
		return
	}

	for _, day := range days {
		benchmarks += fmt.Sprintf(`func Benchmark%d%s(b *testing.B) {
				aoc.RegisterYears()
				input := aoc.TestInput(%d, %d)
				p := aoc.NewPuzzle(%d, %d)
				if p.PartA(input) != nil {
					b.Run("PartA", func(b *testing.B) {
						b.ResetTimer()
						for i := 0; i < b.N; i++ {
							p.PartA(input)
						}
					})
				}
				if p.PartB(input) != nil {
					b.Run("PartB", func(b *testing.B) {
						b.ResetTimer()
						for i := 0; i < b.N; i++ {
							p.PartB(input)
						}
					})
				}
				
		}
		`, year, FormatDay(day), year, day, year, day)
	}

	tmpl := codegen.MustParse(benchmarkingTemplate)
	if err := tmpl.CreateFile(fileName, map[string]interface{}{
		"Year":       year,
		"Benchmarks": benchmarks,
	}); err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Created file: %s", fileName)
}

func RemoveDay(year, day int) {
	sday := fmt.Sprintf("day%s", FormatDay(day))
	files := []string{"desc/" + sday + ".md", "input/" + sday + ".in", "sample/" + sday + ".in", sday + ".go"}
	path := puzzlePath(year)
	for _, file := range files {
		RemoveFile(path + "/" + file)
	}
}

func RemoveYear(year int) {
	for day := range findDays(year) {
		RemoveDay(year, day)
	}
	if _, err := os.Stat(puzzlePath(year)); err == nil {
		RemoveFile(puzzlePath(year))
	}

	UpdateBenchmarks(year)
}

func RemoveAll() {
	for _, year := range findYears() {
		RemoveYear(year)
	}
}

// FormatDay zero pads single-digit days
