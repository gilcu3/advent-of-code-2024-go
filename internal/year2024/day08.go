package year2024

import (
	"fmt"

	"aocgen/internal/util"
)

type Day08 struct{}

func (p Day08) Part1(lines []string) string {
	ar := make([]string, 0)
	mm := make(map[byte]int)
	xs := make([][]int, 0)
	ys := make([][]int, 0)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		for j, c := range line {
			if c != '.' {
				if v, exists := mm[byte(c)]; exists {
					xs[v] = append(xs[v], i)
					ys[v] = append(ys[v], j)
				} else {
					mm[byte(c)] = len(xs)
					xs = append(xs, []int{i})
					ys = append(ys, []int{j})
				}
			}
		}
		ar = append(ar, line)

	}
	n, m := len(ar), len(ar[0])
	ans := 0
	marked := make([][]bool, n)
	for i := range n {
		marked[i] = make([]bool, m)
	}
	for v := range len(xs) {
		vx, vy := xs[v], ys[v]
		k := len(vx)
		for i := range k {
			for j := range k {
				if i != j {
					nx, ny := vx[i]+2*(vx[j]-vx[i]), vy[i]+2*(vy[j]-vy[i])
					if 0 <= nx && nx < n && 0 <= ny && ny < m && !marked[nx][ny] {
						marked[nx][ny] = true
						ans += 1
					}
					if (vx[i]-vx[j])%3 == 0 && (vy[i]-vy[j])%3 == 0 {
						nx, ny = (vx[i]+2*vx[j])/3, (vy[i]+2*vy[j])/3
						if 0 <= nx && nx < n && 0 <= ny && ny < m && !marked[nx][ny] {
							marked[nx][ny] = true
							ans += 1
						}
					}
				}
			}
		}
	}

	return fmt.Sprint(ans)
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func (p Day08) Part2(lines []string) string {
	ar := make([]string, 0)
	mm := make(map[byte]int)
	xs := make([][]int, 0)
	ys := make([][]int, 0)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		for j, c := range line {
			if c != '.' {
				if v, exists := mm[byte(c)]; exists {
					xs[v] = append(xs[v], i)
					ys[v] = append(ys[v], j)
				} else {
					mm[byte(c)] = len(xs)
					xs = append(xs, []int{i})
					ys = append(ys, []int{j})
				}
			}
		}
		ar = append(ar, line)

	}
	n, m := len(ar), len(ar[0])
	ans := 0
	marked := make([][]bool, n)
	for i := range n {
		marked[i] = make([]bool, m)
	}
	for v := range len(xs) {
		vx, vy := xs[v], ys[v]
		k := len(vx)
		for i := range k {
			for j := i + 1; j < k; j++ {
				if !marked[vx[i]][vy[i]] {
					marked[vx[i]][vy[i]] = true
					ans += 1
				}
				g := gcd(abs(vx[j]-vx[i]), abs(vy[j]-vy[i]))
				dx, dy := (vx[j]-vx[i])/g, (vy[j]-vy[i])/g
				nx, ny := vx[i]+dx, vy[i]+dy
				for 0 <= nx && nx < n && 0 <= ny && ny < m {
					if !marked[nx][ny] {
						marked[nx][ny] = true
						ans += 1
					}
					nx, ny = nx+dx, ny+dy
				}
				nx, ny = vx[i]-dx, vy[i]-dy
				for 0 <= nx && nx < n && 0 <= ny && ny < m {
					if !marked[nx][ny] {
						marked[nx][ny] = true
						ans += 1
					}
					nx, ny = nx-dx, ny-dy
				}
			}
		}
	}

	return fmt.Sprint(ans)
}

func (p Day08) TestPart1() {
	const ansExample1 = "14"
	input := util.ExampleInput(2024, 8, 0)
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

func (p Day08) TestPart2() {
	const ansExample2 = "34"
	input := util.ExampleInput(2024, 8, 0)
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
