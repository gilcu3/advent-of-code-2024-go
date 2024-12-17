package year2024

import (
	"fmt"
	"strconv"
	"strings"

	"aocgen/internal/util"
)

type Day17 struct{}

func pow2(a int) int {
	r := 1
	p2 := 2
	for a > 0 {
		if a&1 != 0 {
			r *= p2
		}
		p2 *= p2
		a >>= 1
	}
	return r
}

func (p Day17) Part1(lines []string) string {
	regs := []int{0, 1, 2, 3, 0, 0, 0, 0}
	ops := []int{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Fields(line)
		if parts[1] == "A:" {
			regs[4], _ = strconv.Atoi(parts[2])
		} else if parts[1] == "B:" {
			regs[5], _ = strconv.Atoi(parts[2])
		} else if parts[1] == "C:" {
			regs[6], _ = strconv.Atoi(parts[2])
		} else {
			opss := strings.Split(parts[1], ",")
			for _, o := range opss {
				ov, _ := strconv.Atoi(o)
				ops = append(ops, ov)
			}
		}
	}
	res := []int{}
	for i := 0; i < len(ops); i += 2 {
		// fmt.Println(regs, i, res)
		if ops[i] == 0 {
			regs[4] = regs[4] / pow2(regs[ops[i+1]])
		} else if ops[i] == 1 {
			regs[5] = regs[5] ^ ops[i+1]
		} else if ops[i] == 2 {
			regs[5] = regs[ops[i+1]] % 8
		} else if ops[i] == 3 {
			if regs[4] != 0 {
				if regs[ops[i+1]] != i {
					i = regs[ops[i+1]] - 2
				}
			}
		} else if ops[i] == 4 {
			regs[5] = regs[5] ^ regs[6]

		} else if ops[i] == 5 {
			res = append(res, regs[ops[i+1]]%8)

		} else if ops[i] == 6 {
			regs[5] = regs[4] / pow2(regs[ops[i+1]])
		} else if ops[i] == 7 {
			regs[6] = regs[4] / pow2(regs[ops[i+1]])
		}
	}

	ans := make([]string, len(res))

	for i, v := range res {
		ans[i] = strconv.Itoa(v)
	}
	rans := strings.Join(ans, ",")
	return fmt.Sprint(rans)
}

// this can be done much easier in reverse
func p17rec(ops, bA []int, cur, t int) int {
	if t < 0 {
		return cur
	}
	// fmt.Println(t, bA, cur)
	ans := -1
	for A := range 8 {
		pos := true
		if t+3 == len(bA) && A == 0 {
			continue
		}
		for k := range 3 {
			bA[t+k] = (A >> k) & 1
		}
		oA := A
		A = A + (1<<3)*cur
		B, C := 0, 0
		for i := 0; i < len(ops); i += 2 {
			// fmt.Println(regs, i, res)
			if ops[i] == 0 {
				if ops[i+1] != 3 {
					panic("not in data")
				}
				A = A >> 3
			} else if ops[i] == 1 {
				B = B ^ ops[i+1]
			} else if ops[i] == 2 {
				if ops[i+1] != 4 {
					panic("not in data")
				}
				B = A & 7
			} else if ops[i] == 3 {
				// enforced elsewhere
			} else if ops[i] == 4 {
				B = B ^ C

			} else if ops[i] == 5 {
				if ops[i+1] == 4 {
					if A&7 != ops[t/3] {
						pos = false
					}
				} else if ops[i+1] == 5 {
					if B&7 != ops[t/3] {
						pos = false
					}
				} else {
					panic("not in data")
				}

			} else if ops[i] == 6 {
				panic("not in data")
			} else if ops[i] == 7 {
				if ops[i+1] != 5 {
					panic("not in data")
				}
				C = A >> B
			}
		}
		if pos {

			c := p17rec(ops, bA, cur<<3+oA, t-3)
			if c != -1 && (ans == -1 || ans > c) {
				ans = c
			}

		}
		for k := range 3 {
			bA[t+k] = -1
		}

	}
	return ans

}

func (p Day17) Part2(lines []string) string {
	ops := []int{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Fields(line)
		if parts[1] == "A:" {
			// regs[4], _ = strconv.Atoi(parts[2])
		} else if parts[1] == "B:" {
			// regs[5], _ = strconv.Atoi(parts[2])
		} else if parts[1] == "C:" {
			// regs[6], _ = strconv.Atoi(parts[2])
		} else {
			opss := strings.Split(parts[1], ",")
			for _, o := range opss {
				ov, _ := strconv.Atoi(o)
				ops = append(ops, ov)
			}
		}
	}
	oa := make([]int, len(ops)*3)
	for i := range len(ops) * 3 {
		oa[i] = -1
	}
	ans := p17rec(ops, oa, 0, len(oa)-3)

	return fmt.Sprint(ans)
}

func (p Day17) TestPart1() {
	const ansExample1 = "4,6,3,5,6,3,5,2,1,0"
	input := util.ExampleInput(2024, 17, 1)
	ans := p.Part1(input)
	if ans == fmt.Sprint(nil) {
	} else if ansExample1 == "" {
		fmt.Println("Correct answer Part1 missing, got", ans)
	} else if ans != ansExample1 {
		fmt.Println("Answer to Part1 incorrect", ans, ansExample1)
	} else {
		fmt.Println("Answer to Part1 correct", ans)
	}
}

func (p Day17) TestPart2() {
	const ansExample2 = "117440"
	input := util.ExampleInput(2024, 17, 2)
	ans := p.Part2(input)
	if ans == fmt.Sprint(nil) {
	} else if ansExample2 == "" {
		fmt.Println("Correct answer Part2 missing, got", ans)
	} else if ans != ansExample2 {
		fmt.Println("Answer to Part2 incorrect", ans, ansExample2)
	} else {
		fmt.Println("Answer to Part2 correct", ans)
	}
}
