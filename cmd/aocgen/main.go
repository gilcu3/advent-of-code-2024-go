package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"aocgen/internal/aoc"
	"aocgen/internal/util"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var year, day, part int
var updateReadme bool
var testRun bool
var submitRun bool

var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Run benchmarks for a given puzzle or whole year",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if year == 0 {
			year = time.Now().Year()
		}

		benchArgRegex := fmt.Sprintf("^Benchmark%d", year)
		if day > 0 {
			benchArgRegex += util.FormatDay(day)
		}

		cmdArgs := fmt.Sprintf("go test -bench %s aocgen/%s", benchArgRegex, util.TestsPath)
		c := exec.Command("bash", "-c", cmdArgs)
		out, err := c.Output()
		if err != nil {
			logrus.Error(err)
		}
		results := util.ParseBenchMark(string(out))
		table := util.ParseResults(results)

		if updateReadme {
			util.UpdateBenchmarkResults(results, table, year)
		} else {
			logrus.Info("\n", util.PrintTable(table, year))
		}
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build generated code",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		aoc.RegisterYears()

		aoc.UpdateYearsFile()

		years := aoc.Years()
		for _, year := range years {
			aoc.InitializePackage(year)
			aoc.UpdateBenchmarks(year)
		}
	},
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a puzzle",
	Long:  "Generate a puzzle from year and day inputs",
	Args: func(cmd *cobra.Command, args []string) error {
		if year <= 0 {
			year = time.Now().Year()
		}
		if day <= 0 {
			if time.Now().Month() == 12 {
				day = time.Now().Day()
			} else if time.Now().Month() == 11 && time.Now().Day() == 30 {
				day = 1
			} else {
				return errors.New("invalid day")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		aoc.InitializePackage(year)
		aoc.Download(year, day)
		aoc.NewPuzzleFile(year, day)
		aoc.InitializePackage(year)
		aoc.UpdateYearsFile()
		aoc.UpdateBenchmarks(year)
	},
}

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Display puzzle input for a given puzzle",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		aoc.Check()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all years or list all puzzles in a year",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		aoc.RegisterYears()

		if year != 0 {
			puzzles := aoc.Puzzles(year)
			keys := make([]int, 0)
			keysStrings := make([]string, 0)

			for k := range puzzles {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			for k := range keys {
				keysStrings = append(keysStrings, strconv.Itoa(keys[k]))
			}

			fmt.Printf("%d puzzles completed or in progress:\n", len(keys))
			fmt.Println(strings.Join(keysStrings, ", "))
			return
		}

		years := aoc.Years()
		var yearsStrings []string
		for y := range years {
			yearsStrings = append(yearsStrings, strconv.Itoa(years[y]))
		}

		fmt.Printf("%d years completed or in progress:\n", len(years))
		fmt.Println(strings.Join(yearsStrings, ", "))
	},
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a puzzle and its input",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if year <= 0 {
			aoc.RemoveAll()
		} else if day <= 0 {
			aoc.RemoveYear(year)
		} else {
			aoc.RemoveDay(year, day)
			aoc.UpdateBenchmarks(year)
		}
		aoc.UpdateYearsFile()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a puzzle",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if testRun && submitRun {
			return errors.New("cannot test and submit at the same time")
		}
		if year <= 0 {
			year = time.Now().Year()
		}
		if day <= 0 {
			day = time.Now().Day()
		}
		if part < 0 || part > 2 {
			return errors.New("invalid part")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		aoc.RegisterYears()

		aoc.RunDay(year, day, part, testRun, submitRun)
	},
}

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "AOC is a tool to support completing Advent of Code puzzles",
	Long:  "AOC supports generating puzzle data, including inputs directly from the website, and benchmarking answers",
}

func Execute() {
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", 0, "year input")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "day input")
	benchCmd.Flags().BoolVar(&updateReadme, "update", false, "Update the Readme file")
	runCmd.PersistentFlags().IntVarP(&part, "part", "p", 0, "part input")
	runCmd.Flags().BoolVar(&testRun, "test", false, "Run tests")
	runCmd.Flags().BoolVar(&submitRun, "submit", false, "Run tests")
	rootCmd.AddCommand(benchCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(inputCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
