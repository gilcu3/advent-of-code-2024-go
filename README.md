<img src="./.assets/christmas_gopher.svg" width="164">

# 🎄 Advent of Code 2024

Solutions for [Advent of Code](https://adventofcode.com/) in [Go](https://go.dev/).

<!--- advent_readme_stars table --->
## 2024 Results

| Day | Part 1 | Part 2 |
| :---: | :---: | :---: |
| [Day 1](https://adventofcode.com/2024/day/1) | ⭐ | ⭐ |
| [Day 2](https://adventofcode.com/2024/day/2) | ⭐ | ⭐ |
| [Day 3](https://adventofcode.com/2024/day/3) | ⭐ | ⭐ |
| [Day 4](https://adventofcode.com/2024/day/4) | ⭐ | ⭐ |
| [Day 5](https://adventofcode.com/2024/day/5) | ⭐ | ⭐ |
| [Day 6](https://adventofcode.com/2024/day/6) | ⭐ | ⭐ |
| [Day 7](https://adventofcode.com/2024/day/7) | ⭐ | ⭐ |
| [Day 8](https://adventofcode.com/2024/day/8) | ⭐ | ⭐ |
| [Day 9](https://adventofcode.com/2024/day/9) | ⭐ | ⭐ |
| [Day 10](https://adventofcode.com/2024/day/10) | ⭐ | ⭐ |
| [Day 11](https://adventofcode.com/2024/day/11) | ⭐ | ⭐ |
| [Day 12](https://adventofcode.com/2024/day/12) | ⭐ | ⭐ |
| [Day 13](https://adventofcode.com/2024/day/13) | ⭐ | ⭐ |
<!--- advent_readme_stars table --->

<!--- benchmarking table --->
## Benchmarks

| Day | Part 1 | Part 2 |
| :---: | :---: | :---:  |
| [Day 1](./internal/year2024/day01.go) | `182.817µs` | `257.806µs` |
| [Day 2](./internal/year2024/day02.go) | `239.129µs` | `572.252µs` |
| [Day 3](./internal/year2024/day03.go) | `362.334µs` | `728.405µs` |
| [Day 4](./internal/year2024/day04.go) | `1.390ms` | `1.313ms` |
| [Day 5](./internal/year2024/day05.go) | `536.276µs` | `629.213µs` |
| [Day 6](./internal/year2024/day06.go) | `265.911µs` | `693.900ms` |
| [Day 7](./internal/year2024/day07.go) | `9.614ms` | `760.792ms` |
| [Day 8](./internal/year2024/day08.go) | `46.777µs` | `69.634µs` |
| [Day 9](./internal/year2024/day09.go) | `421.512µs` | `59.309ms` |
| [Day 10](./internal/year2024/day10.go) | `1.532ms` | `223.650µs` |
| [Day 11](./internal/year2024/day11.go) | `8.649ms` | `8.897µs` |
| [Day 12](./internal/year2024/day12.go) | `1.194ms` | `1.120ms` |
| [Day 13](./internal/year2024/day13.go) | `396.389µs` | `130.837µs` |

**Total: 1.544s**
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
