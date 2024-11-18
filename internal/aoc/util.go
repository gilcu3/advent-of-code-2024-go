package aoc

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

	if seconds > 0 {
		return fmt.Sprintf("%.3fm", seconds)
	} else if milliseconds > 0 {
		return fmt.Sprintf("%.3fms", milliseconds)
	} else if microseconds > 0 {
		return fmt.Sprintf("%.3fµs", microseconds)
	} else {
		return fmt.Sprintf("%.3fns", ns)
	}
}

func ParseBenchMark(output string) []BenchmarkResult {
	var benchmarks []BenchmarkResult

	lines := strings.Split(output, "\n")

	re := regexp.MustCompile(`^Benchmark(\d\d\d\d)(\d\d)/Part(A|B)-\d+\s+\d+\s+(\d+|[\d.]+)\s+ns/op`)

	for _, line := range lines {
		if matches := re.FindStringSubmatch(line); matches != nil {
			year, _ := strconv.Atoi(matches[1])
			day, _ := strconv.Atoi(matches[2])
			part := 0
			if matches[3] == "B" {
				part = 1
			}
			// Parse the time as an int64
			var time float64
			fmt.Sscanf(matches[4], "%f", &time)

			benchmarks = append(benchmarks, BenchmarkResult{year, day, part, time})
		}
	}
	return benchmarks
}

func parseResults(results []BenchmarkResult) map[int][]*float64 {
	table := make(map[int][]*float64)
	for i := 1; i <= 25; i++ {
		table[i] = []*float64{nil, nil}
	}
	for _, r := range results {
		table[r.Day][r.Part] = &r.Time
	}
	return table
}

const headerTable = `
## Benchmarks

| Day | Part 1 | Part 2 |
| :---: | :---: | :---:  |
`

const marker = "<!--- benchmarking table --->"

func printTable(table map[int][]*float64) string {
	var result string
	result = ""
	result += marker
	result += headerTable
	total := 0.0
	for i := 1; i <= 25; i++ {
		partA := table[i][0]
		resPartA := "-"
		if partA != nil {
			total += *partA
			resPartA = humanTime(*partA)
		}
		partB := table[i][1]
		resPartB := "-"
		if partB != nil {
			total += *partB
			resPartB = humanTime(*partB)
		}
		if partA != nil || partB != nil {
			result += fmt.Sprintf("| [Day %d](./internal/aoc/day%.2d.go) | `%s` | `%s` |\n", i, i, resPartA, resPartB)
		}

	}
	result += fmt.Sprintf("\n**Total: %s**\n", humanTime(total))
	result += marker
	return result
}

// "| [Day 1](./src/bin/01.rs) | `10ms` | `20ms` |"
func UpdateBenchmarkResults(results []BenchmarkResult) {
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
	table := parseResults(results)
	tableString := printTable(table)
	modReadme := strings.Join([]string{readme[:start], tableString, readme[end:]}, "")
	os.WriteFile(path, []byte(modReadme), 0644)
}
