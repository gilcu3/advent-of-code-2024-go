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
<!--- advent_readme_stars table --->

<!--- benchmarking table --->
## Benchmarks

| Day | Part 1 | Part 2 |
| :---: | :---: | :---:  |
| [Day 1](./internal/year2024/day01.go) | `182.190µs` | `243.919µs` |
| [Day 2](./internal/year2024/day02.go) | `224.742µs` | `502.042µs` |
| [Day 3](./internal/year2024/day03.go) | `382.239µs` | `760.404µs` |
| [Day 4](./internal/year2024/day04.go) | `1.427ms` | `1.393ms` |
| [Day 5](./internal/year2024/day05.go) | `25.657ms` | `362.694ms` |
| [Day 6](./internal/year2024/day06.go) | `257.014µs` | `733.423ms` |
| [Day 7](./internal/year2024/day07.go) | `10.095ms` | `805.884ms` |
| [Day 8](./internal/year2024/day08.go) | `47.977µs` | `69.748µs` |

**Total: 1.943s**
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
