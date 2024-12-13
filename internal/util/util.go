package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const PuzzlePath_tpl = "internal/year%d"
const TemplatePath = "internal/templates"
const YearsFile = "internal/aoc/years.go"
const TestsPath = "internal/tests"

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); err == nil || !errors.Is(err, os.ErrNotExist) {
		logrus.Infof("Directory already exists: %s", path)
		return nil
	}

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	logrus.Infof("Created directory: %s", path)
	return nil
}

func RemoveFile(fileName string) {
	err := os.RemoveAll(fileName)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(fmt.Sprintf("File/Dir deleted: %s", fileName))

}

func FormatDay(day int) string {
	return fmt.Sprintf("%.2d", day)
}

func RemoveFirstLine(filePath string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if len(lines) > 0 {
		lines = lines[1:]
	}

	file, err = os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write line: %w", err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush buffer: %w", err)
	}

	return nil
}

type BenchmarkResult struct {
	Year int
	Day  int
	Part int
	Time float64
}

func humanTime(ns float64) string {
	microseconds := ns / 1000
	milliseconds := microseconds / 1000
	seconds := milliseconds / 1000

	if seconds >= 1 {
		return fmt.Sprintf("%.3fs", seconds)
	} else if milliseconds >= 1 {
		return fmt.Sprintf("%.3fms", milliseconds)
	} else if microseconds >= 1 {
		return fmt.Sprintf("%.3fµs", microseconds)
	} else {
		return fmt.Sprintf("%.3fns", ns)
	}
}

func fromHumanTime(tt string) float64 {
	ans := 0.0
	if strings.HasSuffix(tt, "ns") {
		ans, _ = strconv.ParseFloat(tt[:len(tt)-2], 64)
	} else if strings.HasSuffix(tt, "µs") {
		ss, _ := strings.CutSuffix(tt, "µs")
		ans, _ = strconv.ParseFloat(ss, 64)
		ans *= 1000
	} else if strings.HasSuffix(tt, "ms") {
		ans, _ = strconv.ParseFloat(tt[:len(tt)-2], 64)
		ans *= 1000000
	} else if strings.HasSuffix(tt, "s") {
		ans, _ = strconv.ParseFloat(tt[:len(tt)-1], 64)
		ans *= 1000000000
	}

	return ans
}

func ParseBenchMark(output string) []BenchmarkResult {
	var benchmarks []BenchmarkResult

	lines := strings.Split(output, "\n")

	re := regexp.MustCompile(`^Benchmark(\d\d\d\d)(\d\d)/Part(1|2)-\d+\s+\d+\s+(\d+|[\d.]+)\s+ns/op`)

	for _, line := range lines {
		if matches := re.FindStringSubmatch(line); matches != nil {
			year, _ := strconv.Atoi(matches[1])
			day, _ := strconv.Atoi(matches[2])
			part := 1
			if matches[3] == "2" {
				part = 2
			}
			time, _ := strconv.ParseFloat(matches[4], 64)

			benchmarks = append(benchmarks, BenchmarkResult{year, day, part, time})
		}
	}
	return benchmarks
}

func ParseResults(results []BenchmarkResult) map[int][]*float64 {
	table := make(map[int][]*float64)
	for i := 1; i <= 25; i++ {
		table[i] = []*float64{nil, nil}
	}
	for _, r := range results {
		table[r.Day][r.Part-1] = &r.Time
	}
	return table
}

const headerTable = `
## Benchmarks

| Day | Part 1 | Part 2 |
| :---: | :---: | :---:  |
`

const marker = "<!--- benchmarking table --->"

func PrintTable(table map[int][]*float64, year int) string {
	var result string
	result = ""
	result += marker
	result += headerTable
	total := 0.0
	for i := 1; i <= 25; i++ {
		part1 := table[i][0]
		resPart1 := "-"
		if part1 != nil {
			total += *part1
			resPart1 = humanTime(*part1)
		}
		part2 := table[i][1]
		resPart2 := "-"
		if part2 != nil {
			total += *part2
			resPart2 = humanTime(*part2)
		}
		if part1 != nil || part2 != nil {
			result += fmt.Sprintf("| [Day %d](./internal/year%d/day%s.go) | `%s` | `%s` |\n", i, year, FormatDay(i), resPart1, resPart2)
		}

	}
	result += fmt.Sprintf("\n**Total: %s**\n", humanTime(total))
	result += marker
	return result
}

func ExtractTable(tableString string) map[int][]*float64 {
	table := make(map[int][]*float64)
	for i := 1; i <= 25; i++ {
		table[i] = []*float64{nil, nil}
	}
	lines := strings.Split(tableString, "\n")

	re := regexp.MustCompile(`^\| \[Day (\d+)\][^\|]*\| .(\d+(?:\.\d+)?.?s). \| .(\d+(?:\.\d+)?.?s). \|$`)

	for _, line := range lines {
		if matches := re.FindStringSubmatch(line); matches != nil {
			day, _ := strconv.Atoi(matches[1])
			p1 := fromHumanTime(matches[2])
			p2 := fromHumanTime(matches[3])
			table[day] = []*float64{&p1, &p2}
		}
	}
	return table
}

func UpdateBenchmarkResults(results []BenchmarkResult, table map[int][]*float64, year int) {

	path := "README.md"
	readmeBytes, _ := os.ReadFile(path)
	readme := string(readmeBytes)
	start := strings.Index(readme, marker)
	if start == -1 {
		logrus.Errorf("Could not find start marker in README.md")
		return
	}
	var end int
	end = strings.Index(readme[start+len(marker):], marker)
	if end == -1 {
		logrus.Errorf("Could not find end marker in README.md")
		return
	}
	end += 2*len(marker) + start
	table0 := ExtractTable(readme[start:end])
	for i := range 25 {
		if table[i+1][0] != nil {
			table0[i+1] = table[i+1]
		}
	}
	tableString := PrintTable(table0, year)
	logrus.Infof("\n" + tableString)
	modReadme := strings.Join([]string{readme[:start], tableString, readme[end:]}, "")
	os.WriteFile(path, []byte(modReadme), 0644)
}

func Input(year, day int) []string {
	fileName := fmt.Sprintf(PuzzlePath_tpl+"/input/day%s.in", year, FormatDay(day))
	return readFile(fileName)
}

func ExampleInput(year, day, part int) []string {
	var fileName string
	if part == 0 {
		fileName = fmt.Sprintf(PuzzlePath_tpl+"/example/day%s.in", year, FormatDay(day))
	} else {
		fileName = fmt.Sprintf(PuzzlePath_tpl+"/example/day%s-%d.in", year, FormatDay(day), part)
	}

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
