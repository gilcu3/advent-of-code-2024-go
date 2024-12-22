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
| [Day 14](https://adventofcode.com/2024/day/14) | ⭐ | ⭐ |
| [Day 15](https://adventofcode.com/2024/day/15) | ⭐ | ⭐ |
| [Day 16](https://adventofcode.com/2024/day/16) | ⭐ | ⭐ |
| [Day 17](https://adventofcode.com/2024/day/17) | ⭐ | ⭐ |
| [Day 18](https://adventofcode.com/2024/day/18) | ⭐ | ⭐ |
| [Day 19](https://adventofcode.com/2024/day/19) | ⭐ | ⭐ |
| [Day 20](https://adventofcode.com/2024/day/20) | ⭐ | ⭐ |
| [Day 21](https://adventofcode.com/2024/day/21) | ⭐ | ⭐ |
<!--- advent_readme_stars table --->

<!--- benchmarking table --->
## Benchmarks

| Day | Part 1 | Part 2 |
| :---: | :---: | :---:  |
| [Day 1](./internal/year2024/day01.go) | `192.328µs` | `277.442µs` |
| [Day 2](./internal/year2024/day02.go) | `271.722µs` | `668.116µs` |
| [Day 3](./internal/year2024/day03.go) | `355.281µs` | `700.291µs` |
| [Day 4](./internal/year2024/day04.go) | `1.310ms` | `1.251ms` |
| [Day 5](./internal/year2024/day05.go) | `546.704µs` | `680.442µs` |
| [Day 6](./internal/year2024/day06.go) | `187.472µs` | `24.040ms` |
| [Day 7](./internal/year2024/day07.go) | `2.406ms` | `4.369ms` |
| [Day 8](./internal/year2024/day08.go) | `37.147µs` | `54.927µs` |
| [Day 9](./internal/year2024/day09.go) | `302.286µs` | `55.452ms` |
| [Day 10](./internal/year2024/day10.go) | `1.050ms` | `162.316µs` |
| [Day 11](./internal/year2024/day11.go) | `7.602ms` | `7.479µs` |
| [Day 12](./internal/year2024/day12.go) | `967.359µs` | `1.036ms` |
| [Day 13](./internal/year2024/day13.go) | `359.420µs` | `110.144µs` |
| [Day 14](./internal/year2024/day14.go) | `887.116µs` | `147.532ms` |
| [Day 15](./internal/year2024/day15.go) | `537.784µs` | `1.055ms` |
| [Day 16](./internal/year2024/day16.go) | `7.426ms` | `15.831ms` |
| [Day 17](./internal/year2024/day17.go) | `7.749µs` | `169.068µs` |
| [Day 18](./internal/year2024/day18.go) | `3.490ms` | `6.408ms` |
| [Day 19](./internal/year2024/day19.go) | `10.162ms` | `31.491ms` |
| [Day 20](./internal/year2024/day20.go) | `4.910ms` | `148.379ms` |
| [Day 21](./internal/year2024/day21.go) | `180.354µs` | `1.261ms` |
| [Day 22](./internal/year2024/day22.go) | `24.790ms` | `42.216ms` |

**Total: 551.130ms**
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
