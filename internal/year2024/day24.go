package year2024

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aocgen/internal/util"
)

type Day24 struct{}

func (p Day24) conv(c byte, mm map[string]int, vals []int) int {
	x := 0
	for k, v := range mm {
		if vals[v] == -1 {
			return -1
		}
		if k[0] == c && vals[v] == 1 {
			o, _ := strconv.Atoi(k[1:])
			x ^= 1 << o
		}
	}
	return x
}

func (p Day24) Part1(lines []string) string {
	sec := 0
	mm := make(map[string]int)
	edges := [][]int{}
	vals := []int{}
	ind := 0
	gg := [][]int{}
	for _, line := range lines {
		if len(line) == 0 {
			sec += 1
			continue
		}
		if sec == 0 {
			parts := strings.Split(line, ": ")
			mm[parts[0]] = ind
			ind += 1
			v, _ := strconv.Atoi(parts[1])
			vals = append(vals, v)
			gg = append(gg, []int{})
		} else {
			parts := strings.Split(line, " ")
			for _, p := range []int{0, 2, 4} {
				if _, ok := mm[parts[p]]; !ok {
					mm[parts[p]] = ind
					vals = append(vals, -1)
					gg = append(gg, []int{})
					ind += 1
				}
			}
			op := []int{0, 0, 0, 0}
			for i, p := range []int{0, 2, 4} {
				op[i+1] = mm[parts[p]]
			}
			if parts[1] == "AND" {
				op[0] = 0
			} else if parts[1] == "OR" {
				op[0] = 1
			} else {
				op[0] = 2
			}
			gg[mm[parts[0]]] = append(gg[mm[parts[0]]], len(edges))
			gg[mm[parts[2]]] = append(gg[mm[parts[2]]], len(edges))
			edges = append(edges, op)
		}
	}
	que := []int{}
	front := 0
	for i := range ind {
		if vals[i] != -1 {
			que = append(que, i)
		}
	}
	for front < len(que) {
		i := que[front]
		front += 1
		for _, e := range gg[i] {
			o, a, b, c := edges[e][0], edges[e][1], edges[e][2], edges[e][3]
			if vals[a] != -1 && vals[b] != -1 {
				if o == 0 {
					vals[c] = vals[a] & vals[b]
				} else if o == 1 {
					vals[c] = vals[a] | vals[b]
				} else {
					vals[c] = vals[a] ^ vals[b]
				}
				que = append(que, c)
			}
		}
	}
	ans := p.conv('z', mm, vals)
	return fmt.Sprint(ans)
}

func (p Day24) Part2(lines []string) string {
	sec := 0
	mm := make(map[string]int)
	edges := [][]int{}
	ind := 0
	gg := [][]int{}
	par := []int{}
	for _, line := range lines {
		if len(line) == 0 {
			sec += 1
			continue
		}
		if sec == 0 {
			parts := strings.Split(line, ": ")
			mm[parts[0]] = ind
			ind += 1
			gg = append(gg, []int{})
			par = append(par, -1)
		} else {
			parts := strings.Split(line, " ")
			for _, p := range []int{0, 2, 4} {
				if _, ok := mm[parts[p]]; !ok {
					mm[parts[p]] = ind
					gg = append(gg, []int{})
					par = append(par, -1)
					ind += 1
				}
			}
			op := []int{0, 0, 0, 0}
			for i, p := range []int{0, 2, 4} {
				op[i+1] = mm[parts[p]]
			}
			if parts[1] == "AND" {
				op[0] = 0
			} else if parts[1] == "OR" {
				op[0] = 1
			} else {
				op[0] = 2
			}
			par[mm[parts[4]]] = len(edges)
			gg[mm[parts[0]]] = append(gg[mm[parts[0]]], len(edges))
			gg[mm[parts[2]]] = append(gg[mm[parts[2]]], len(edges))
			edges = append(edges, op)
		}
	}
	strs := make([]string, ind)
	for k, v := range mm {
		strs[v] = k
	}
	mx, my, mz := 0, 0, 0
	for k := range mm {
		if k[0] == 'x' {
			mx += 1
		} else if k[0] == 'y' {
			my += 1
		} else if k[0] == 'z' {
			mz += 1
		}
	}

	// structure of addition
	// z0 = x0 ^ y0
	// of0 = x0 & y0
	// zi = (xi ^ yi) ^ ofi-1
	// ofi = (xi & yi) | ((xi ^ yi) & ofi-1)
	xor, and, x, y, of := make([]int, mx), make([]int, mx), make([]int, mx), make([]int, mx), make([]int, mx)

	for i := range mx {
		x[i] = mm[fmt.Sprintf("x%.2d", i)]
		y[i] = mm[fmt.Sprintf("y%.2d", i)]
		xor[i] = -1
		and[i] = -1
		of[i] = -1
	}
	z := make([]int, mz)
	for i := range mz {
		z[i] = mm[fmt.Sprintf("z%.2d", i)]
	}
	for e := range len(edges) {
		o, a, b, _ := edges[e][0], edges[e][1], edges[e][2], edges[e][3]
		if o == 0 {
			for i := range mx {
				if (x[i] == a && y[i] == b) || (x[i] == b && y[i] == a) {
					and[i] = e
					break
				}
			}
		}
		if o == 2 {
			for i := range mx {
				if (x[i] == a && y[i] == b) || (x[i] == b && y[i] == a) {
					xor[i] = e
					break
				}
			}
		}
	}
	changed := []int{}
	for i := range mx {
		if i == 0 {
			e := xor[0]
			c := edges[e][3]
			if c != z[0] {
				changed = append(changed, e)
			}
			e = and[0]
			c = edges[e][3]
			if len(gg[c]) != 2 {
				changed = append(changed, e)
			}
			continue
		}
		e := xor[i]
		c := edges[e][3]
		if len(gg[c]) != 2 {
			changed = append(changed, e)
		} else {
			p := -1
			cxor, cand := false, false
			pos := true
			for _, ee := range gg[c] {
				o, a, b, _ := edges[ee][0], edges[ee][1], edges[ee][2], edges[ee][3]
				if o == 2 {
					cxor = true
				} else if o == 0 {
					cand = true
				}
				if a == c {
					if p != -1 && p != b {
						pos = false
					}
					p = b
				}
				if b == c {
					if p != -1 && p != a {
						pos = false
					}
					p = a
				}
			}
			if !cxor || !pos || !cand {
				changed = append(changed, e)
			} else {
				for _, ee := range gg[c] {
					o, _, _, cc := edges[ee][0], edges[ee][1], edges[ee][2], edges[ee][3]
					if o == 2 {
						if cc != z[i] {
							changed = append(changed, ee)
						}
					} else if o == 0 {
						if len(gg[cc]) != 1 {
							changed = append(changed, ee)
						}
					}
				}
			}
		}

		e = and[i]
		c = edges[e][3]
		if len(gg[c]) != 1 {
			changed = append(changed, e)
		} else {
			ee := gg[c][0]
			o, a, b, _ := edges[ee][0], edges[ee][1], edges[ee][2], edges[ee][3]
			if o != 1 {
				changed = append(changed, e)
			} else if a == c {
				if len(gg[b]) != 1 || edges[gg[b][0]][0] != 1 {
					changed = append(changed, e)
				}
			} else if a == b {
				if len(gg[c]) != 1 || edges[gg[c][0]][0] != 1 {
					changed = append(changed, e)
				}
			}
		}
		e = par[z[i]]
		if edges[e][0] != 2 {
			changed = append(changed, e)
		}
	}

	changeds := []string{}
	for _, v := range changed {
		changeds = append(changeds, strs[edges[v][3]])
	}
	sort.Slice(changeds, func(i, j int) bool { return changeds[i] < changeds[j] })
	uni := []string{}
	for _, s := range changeds {
		if len(uni) == 0 || s != uni[len(uni)-1] {
			uni = append(uni, s)
		}
	}
	return strings.Join(uni, ",")
}

func (p Day24) TestPart1() {
	ansExample1 := []string{"", "4", "2024"}
	for i := 1; i < 3; i++ {
		input := util.ExampleInput(2024, 24, i)
		ans := p.Part1(input)
		if ans == fmt.Sprint(nil) {
		} else if ansExample1[i] == "" {
			fmt.Println("Correct answer Part1 missing, got", ans)
		} else if ans != ansExample1[i] {
			fmt.Println("Answer to Part1 incorrect", ans, ansExample1[i])
		} else {
			fmt.Println("Answer to Part1 correct", ans)
		}
	}

}

func (p Day24) TestPart2() {
	const ansExample2 = "t01,z01"
	input := util.ExampleInput(2024, 24, 3)
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
