package aoc

import (
	"aocgen/internal/util"
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"text/template"

	"github.com/dolmen-go/codegen"
	"github.com/sirupsen/logrus"
)

var validPuzzleFile = regexp.MustCompile(`^day[0-3][0-9]$`)

func puzzlePath(year int) string {
	return fmt.Sprintf("internal/year%d", year)
}

func puzzleFileName(year, day int) string {
	path := puzzlePath(year)
	return fmt.Sprintf("%s/day%s.go", path, util.FormatDay(day))
}

func NewPuzzleFile(year, day int) {
	fileName := puzzleFileName(year, day)
	if _, err := os.Stat(fileName); err == nil || !errors.Is(err, os.ErrNotExist) {
		logrus.Infof("Puzzle file already exists: %s", fileName)
		return
	}

	puzzleTemplate, err := os.ReadFile(util.TemplatePath + "/" + "puzzle.tpl")
	if err != nil {
		logrus.Fatal(err)
	}
	tmpl := codegen.MustParse(string(puzzleTemplate))
	if err := tmpl.CreateFile(fileName, map[string]any{
		"Year": year,
		"Day":  util.FormatDay(day),
	}); err != nil {
		logrus.Fatal(err)
	}

	util.RemoveFirstLine(fileName)

	logrus.Infof("Created file: %s", fileName)
}

func InitializePackage(year int) {
	path := puzzlePath(year)
	if err := util.CreateDirectory(path); err != nil {
		logrus.Fatal(err)
	}
}

func findYears() []int {

	dirs, err := os.ReadDir("internal")
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
	pathYear := fmt.Sprintf("internal/year%d", year)
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
		imports += fmt.Sprintf("\"aocgen/internal/year%d\"\n", year)
		for _, day := range days {
			correctPuzzleName := fmt.Sprintf("Day%s", util.FormatDay(day))
			puzzles += fmt.Sprintf("%d: year%d.%s{},\n", day, year, correctPuzzleName)
			logrus.Debugf("Found puzzle file for %d-%d", year, day)
		}
		inits += fmt.Sprintf("Register(%d, map[int]Puzzle{%s})", year, puzzles)
	}
	initMainTemplate, err := os.ReadFile(util.TemplatePath + "/" + "years.tpl")
	if err != nil {
		logrus.Fatal(err)
	}
	tmpl := codegen.MustParse(string(initMainTemplate))

	err = os.Remove(util.YearsFile)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Deleted file: %s", util.YearsFile)

	if err := tmpl.CreateFile(util.YearsFile, map[string]interface{}{
		"Imports": imports,
		"Inits":   inits,
	}); err != nil {
		logrus.Fatal(err)
	}
}

func UpdateBenchmarks(year int) {
	pathTests := "internal/tests"
	fileName := fmt.Sprintf("%s/year%d_test.go", pathTests, year)

	benchmarks := ""

	util.CreateDirectory(pathTests)
	days := findDays(year)
	if len(days) == 0 {
		os.Remove(fileName)
		return
	}
	singleBenchTemplate, err := os.ReadFile(util.TemplatePath + "/" + "bench_function.tpl")
	if err != nil {
		logrus.Fatal(err)
	}

	t := template.Must(template.New("").Parse(string(singleBenchTemplate)))
	vars := struct {
		Year      int
		Day       int
		FormatDay string
	}{
		Year:      year,
		Day:       0,
		FormatDay: "",
	}
	for _, day := range days {
		var output bytes.Buffer
		vars.Day = day
		vars.FormatDay = util.FormatDay(day)
		if err := t.Execute(&output, vars); err != nil {
			logrus.Fatal(err)
		}
		benchmarks += output.String()
	}
	benchmarkingTemplate, err := os.ReadFile(util.TemplatePath + "/" + "test.tpl")
	if err != nil {
		logrus.Fatal(err)
	}
	tmpl := codegen.MustParse(string(benchmarkingTemplate))
	if err := tmpl.CreateFile(fileName, map[string]interface{}{
		"Year":       year,
		"Benchmarks": benchmarks,
	}); err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Created file: %s", fileName)
}

func RemoveDay(year, day int) {
	sday := fmt.Sprintf("day%s", util.FormatDay(day))
	files := []string{"desc/" + sday + ".md", "input/" + sday + ".in", "example/" + sday + ".in", sday + ".go"}
	path := puzzlePath(year)
	for _, file := range files {
		util.RemoveFile(path + "/" + file)
	}
}

func RemoveYear(year int) {
	for day := range findDays(year) {
		RemoveDay(year, day)
	}
	if _, err := os.Stat(puzzlePath(year)); err == nil {
		util.RemoveFile(puzzlePath(year))
	}

	UpdateBenchmarks(year)
}

func RemoveAll() {
	for _, year := range findYears() {
		RemoveYear(year)
	}
}

// FormatDay zero pads single-digit days
