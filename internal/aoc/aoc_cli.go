package aoc

import (
	"aocgen/internal/util"
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
	return fmt.Sprintf(util.PuzzlePath_tpl+"/input/day%s.in", year, util.FormatDay(day))
}

func getExamplePath(year, day int) string {
	return fmt.Sprintf(util.PuzzlePath_tpl+"/example/day%s.in", year, util.FormatDay(day))
}

func getDescPath(year, day int) string {
	return fmt.Sprintf(util.PuzzlePath_tpl+"/desc/day%s.md", year, util.FormatDay(day))
}

func createPuzzleDirs(year int) {
	path0 := fmt.Sprintf(util.PuzzlePath_tpl+"/input/", year)
	path1 := fmt.Sprintf(util.PuzzlePath_tpl+"/example/", year)
	path2 := fmt.Sprintf(util.PuzzlePath_tpl+"/desc/", year)
	for _, path := range []string{path0, path1, path2} {
		util.CreateDirectory(path)
	}
}

func createExampleFile(year, day int) {
	exampleFileName := getExamplePath(year, day)
	if _, err := os.Stat(exampleFileName); err == nil {
		logrus.Infof("Example file already exists: %s", exampleFileName)
		return
	}
	fExample, errExample := os.Create(exampleFileName)
	if errExample != nil {
		logrus.Fatal(errExample)
	}
	defer fExample.Close()

	logrus.Infof("Generated empty example file: %s", exampleFileName)
}

func Download(year, day int) {
	descPath := getDescPath(year, day)
	inputPath := getInputPath(year, day)
	createPuzzleDirs(year)
	createExampleFile(year, day)
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
