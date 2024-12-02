<img src="./.assets/christmas_gopher.svg" width="164">

# 🎄 Advent of Code 2024

Solutions for [Advent of Code](https://adventofcode.com/) in [Go](https://go.dev/).

<!--- advent_readme_stars table --->
## 2024 Results

| Day | Part 1 | Part 2 |
| :---: | :---: | :---: |
<!--- advent_readme_stars table --->

<!--- benchmarking table --->
## Benchmarks

| Day | Part 1 | Part 2 |
| :---: | :---: | :---:  |
| [Day 1](./internal/year2024/day01.go) | `218.654µs` | `276.380µs` |
| [Day 2](./internal/year2024/day02.go) | `262.666µs` | `739.853µs` |

**Total: 1.498ms**
<!--- benchmarking table --->

---

<details>
<summary>Template readme</summary>

# AOCgen

AOCgen is a tool to assist in solving Advent of Code in Go. This is a heavily
modified fork.

## Setup

You need to set up [aoc-cli](https://github.com/scarvalhojr/aoc-cli) in order to download input and making submissions automatically.

Run AOCgen via executable: ```./aocgen```

### Commands

- **bench**: run benchmarks for a given puzzle or year of puzzles
- **build**: run code generation suite, useful for when you've had to remove any code
- **gen**: generate a puzzle
- **input**: display input for a puzzle in the console
- **list**: list all years or puzzles in a year
- **rm**: delete a puzzle and its input
- **run**: run a puzzle

## Generating Code

Use ```aocgen``` via the ```gen``` subcommand to generate code: ```./aocgen gen -y <year> -d <day>```

This will generate two files: the puzzle (```pkg/year<year>/<day>.go```) and its input (```pkg/year<year>/inputs/<day>.txt```)

Open up the puzzle and remove the DO NOT EDIT line to begin working.

Run the puzzle through the ```aocgen``` command as well: ```./aocgen run -y <year> -d <day>```

### Benchmarking

Again, use ```aocgen``` to run benchmarks for a specific day's puzzle or the entire year:

Day: ```./aocgen bench -y <year> -d <day>```

Year: ```./aocgen bench -y <year>```
