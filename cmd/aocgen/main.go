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

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var year, day int
var updateReadme bool

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
			benchArgRegex += fmt.Sprintf("Day%s", aoc.FormatDay(day))
		}

		cmdArgs := fmt.Sprintf("go test -bench %s aocgen/internal/aoc/year%d", benchArgRegex, year)
		c := exec.Command("bash", "-c", cmdArgs)
		out, err := c.Output()
		if err != nil {
			logrus.Error(err)
		}
		println(string(out))

		if updateReadme {
			results := aoc.ParseBenchMark(string(out))
			aoc.UpdateBenchmarkResults(results)
		}
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build generated code",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		aoc.RegisterYears()

		aoc.InitializeYearsPackages()

		years := aoc.Years()
		for _, y := range years {
			aoc.InitializePackage(y)
			aoc.NewBenchmarks(y)
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
		aoc.InitializeYearsPackages()
		aoc.NewBenchmarks(year)
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
		if year <= 0 {
			return errors.New("invalid year")
		}
		if day <= 0 {
			return errors.New("invalid day")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		aoc.RemovePuzzle(year, day)
		aoc.RemovePuzzleInput(year, day)
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a puzzle",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if year <= 0 {
			logrus.Fatal("invalid year")
		}

		aoc.RegisterYears()

		if day > 0 {
			runDay(year, day)
			return
		}

		runYear(year)
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

func runYear(year int) {
	puzzles := aoc.Puzzles(year)
	for i := 1; i <= len(puzzles); i++ {
		runDay(year, i)
	}
}

func runDay(year, day int) {
	aoc.Run(year, day, aoc.NewPuzzle(year, day), aoc.Input(year, day))
}

func main() {
	Execute()
}
