package aoc

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/sirupsen/logrus"
)

func Check() bool {
	c := exec.Command("aoc", "-V")
	_, err := c.Output()
	if err != nil {
		logrus.Error(err)
		return false
	}
	return true
}

func call_aoc_cli(args []string) string {
	c := exec.Command("aoc", args...)
	output, err := c.Output()
	if err != nil {
		logrus.Error(err)
	}
	return string(output)
}

func getInputPath(year, day int) string {
	return fmt.Sprintf("internal/aoc/year%d/input/day%s.in", year, FormatDay(day))
}

func getSamplePath(year, day int) string {
	return fmt.Sprintf("internal/aoc/year%d/sample/day%s.in", year, FormatDay(day))
}

func getDescPath(year, day int) string {
	return fmt.Sprintf("internal/aoc/year%d/desc/day%s.md", year, FormatDay(day))
}

func createPuzzleDirs(year int) {
	path0 := fmt.Sprintf("internal/aoc/year%d/input/", year)
	path1 := fmt.Sprintf("internal/aoc/year%d/sample/", year)
	path2 := fmt.Sprintf("internal/aoc/year%d/desc/", year)
	for _, path := range []string{path0, path1, path2} {
		CreateDirectory(path)
	}
}

func createSampleFile(year, day int) {
	sampleFileName := getSamplePath(year, day)
	if _, err := os.Stat(sampleFileName); err != nil && err == os.ErrNotExist {
		logrus.Infof("Sample file already exists: %s", sampleFileName)
		return
	}
	fSample, errSample := os.Create(sampleFileName)
	if errSample != nil {
		logrus.Fatal(errSample)
	}
	defer fSample.Close()

	logrus.Infof("Generated empty sample file: %s", sampleFileName)
}

func Download(year, day int) {
	descPath := getDescPath(year, day)
	inputPath := getInputPath(year, day)
	createPuzzleDirs(year)
	createSampleFile(year, day)
	args := []string{"download", "--overwrite", "--input-file", inputPath, "--puzzle-file", descPath, "--year", strconv.Itoa(year), "--day", strconv.Itoa(day)}
	call_aoc_cli(args)
	logrus.Infof("🎄 Successfully wrote input to %s.", inputPath)
	logrus.Infof("🎄 Successfully wrote description to %s.", descPath)
}

func Submit(year, day, part int, ans string) {
	args := []string{"submit", "--year", strconv.Itoa(year), "--day", strconv.Itoa(day), strconv.Itoa(part), ans}
	output := call_aoc_cli(args)
	logrus.Infof("🎄 Successfully submitted. Result: %s", output)
}
