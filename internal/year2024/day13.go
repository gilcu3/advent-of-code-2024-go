package year2024

import (
	"fmt"
	"strconv"
	"strings"

	"aocgen/internal/util"
)

type Day13 struct{}

func (p Day13) Part1(lines []string) string {
	ans := 0
	ax, ay, bx, by := 0, 0, 0, 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		cur := -1
		if parts[0] == "Prize:" {
			rx, _ := strconv.Atoi(parts[1][2 : len(parts[1])-1])
			ry, _ := strconv.Atoi(parts[2][2:])
			for ta := range 101 {
				tx, ty := rx-ta*ax, ry-ta*ay
				if tx < 0 || ty < 0 {
					break
				}
				if tx%bx == 0 && ty%by == 0 && tx/bx == ty/by {
					cc := 3*ta + tx/bx
					if cur == -1 || cc < cur {
						cur = cc
					}
				}
			}
			if cur != -1 {
				ans += cur
			}
		} else if parts[1] == "A:" {
			ax, _ = strconv.Atoi(parts[2][2 : len(parts[2])-1])
			ay, _ = strconv.Atoi(parts[3][2:])
		} else {
			bx, _ = strconv.Atoi(parts[2][2 : len(parts[2])-1])
			by, _ = strconv.Atoi(parts[3][2:])
		}
	}

	return fmt.Sprint(ans)
}

func egcd(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := egcd(b, a%b)
	x = y1
	y = x1 - (a/b)*y1
	return gcd, x, y
}

func (p Day13) Part2(lines []string) string {
	ans := 0
	ax, ay, bx, by := 0, 0, 0, 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)
		cur := -1
		if parts[0] == "Prize:" {
			rx, _ := strconv.Atoi(parts[1][2 : len(parts[1])-1])
			ry, _ := strconv.Atoi(parts[2][2:])
			rx += 10000000000000
			ry += 10000000000000
			delta := ax*by - ay*bx
			if delta == 0 {
				// wtf this case never happens!
				g, ta, tb := egcd(ax, bx)
				if rx%g == 0 {
					rx /= g
					ax /= g
					bx /= g
					ta *= rx
					tb *= rx

					if 3*bx >= ax {
						if ta >= bx {
							f := ta / bx
							ta -= f * bx
							tb += f * ax
						} else if ta < 0 {
							f := (-ta + bx - 1) / bx
							ta += f * bx
							tb -= f * ax
						}

					} else {
						if tb >= ax {
							f := tb / ax
							tb -= f * ax
							ta += f * bx
						} else if tb < 0 {
							f := (-tb + ax - 1) / ax
							tb += f * ax
							ta -= f * bx
						}
					}
					if ta >= 0 && tb >= 0 {
						cur = 3*ta + tb
					}
				}
			} else {
				ta, tb := by*rx-bx*ry, ay*rx-ax*ry
				if ta%delta == 0 && tb%delta == 0 {
					cur = (3*ta - tb) / delta
				}
			}
			if cur != -1 {
				ans += cur
			}
		} else if parts[1] == "A:" {
			ax, _ = strconv.Atoi(parts[2][2 : len(parts[2])-1])
			ay, _ = strconv.Atoi(parts[3][2:])
		} else {
			bx, _ = strconv.Atoi(parts[2][2 : len(parts[2])-1])
			by, _ = strconv.Atoi(parts[3][2:])
		}
	}

	return fmt.Sprint(ans)
}

func (p Day13) TestPart1() {
	const ansExample1 = "480"
	input := util.ExampleInput(2024, 13, 0)
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

func (p Day13) TestPart2() {
	const ansExample2 = "875318608908"
	input := util.ExampleInput(2024, 13, 0)
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
