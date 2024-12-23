package year2024

import (
	"fmt"
	"strings"

	"aocgen/internal/util"
)

type Day23 struct{}

func (p Day23) isGood(node, t int) bool {
	return node/26%26 == t
}

func (p Day23) Part1(lines []string) string {
	mx := 26 * 26
	gg := make([][]int, mx)
	t := int('t' - 'a')

	for i := range mx {
		gg[i] = make([]int, mx)
	}
	edges := make([][]int, mx)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		nodes := strings.Split(line, "-")
		nn := []int{0, 0}
		for i := range 2 {
			for j := range 2 {
				c := int(nodes[i][j] - byte('a'))
				nn[i] = nn[i]*26 + c
			}
		}
		gg[nn[1]][nn[0]], gg[nn[0]][nn[1]] = 1, 1
		edges[nn[0]] = append(edges[nn[0]], nn[1])
		edges[nn[1]] = append(edges[nn[1]], nn[0])
	}
	ans := []int{0, 0, 0}
	for v := range mx {
		cc := 0
		if p.isGood(v, t) {
			for i, nv1 := range edges[v] {
				if p.isGood(nv1, t) {
					cc += 1
				}
				for _, nv2 := range edges[v][i+1:] {
					if p.isGood(nv2, t) {
						cc += 1
					}
					if gg[nv1][nv2] == 1 {
						ans[cc] += 1
					}
					if p.isGood(nv2, t) {
						cc -= 1
					}
				}
				if p.isGood(nv1, t) {
					cc -= 1
				}
			}
		}

	}
	return fmt.Sprint(ans[0] + ans[1]/2 + ans[2]/3)
}

func (p Day23) conv2str(v int) string {
	b := []byte{byte(v/26 + 'a'), byte(v%26 + 'a')}
	return string(b)
}

// useless exploration
// func (p Day23) BiconnectedComponents(adjList [][]int) [][][]int {
// 	n := len(adjList)
// 	discovery := make([]int, n)
// 	low := make([]int, n)
// 	parent := make([]int, n)
// 	bccStack := [][][]int{}
// 	time := 0
// 	stack := [][]int{}
// 	for i := range n {
// 		discovery[i] = -1
// 		parent[i] = -1
// 	}

// 	var dfs func(int)

// 	dfs = func(u int) {
// 		discovery[u] = time
// 		low[u] = time
// 		time++

// 		children := 0

// 		for _, v := range adjList[u] {
// 			if discovery[v] == -1 {
// 				parent[v] = u
// 				stack = append(stack, []int{u, v})
// 				children += 1

// 				dfs(v)

// 				low[u] = min(low[u], low[v])

// 				if (parent[u] == -1 && children > 1) || (parent[u] != -1 && low[v] >= discovery[u]) {
// 					bcc := [][]int{}
// 					for {
// 						edge := stack[len(stack)-1]
// 						stack = stack[:len(stack)-1]
// 						bcc = append(bcc, edge)
// 						if edge[0] == u && edge[1] == v {
// 							break
// 						}
// 					}
// 					bccStack = append(bccStack, bcc)
// 				}
// 			} else if v != parent[u] && low[u] > discovery[v] {
// 				low[u] = min(low[u], discovery[v])
// 				stack = append(stack, []int{u, v})
// 			}
// 		}

// 	}

// 	for i := 0; i < n; i++ {
// 		if discovery[i] == -1 {
// 			dfs(i)
// 			if len(stack) > 0 {
// 				bcc := [][]int{}
// 				for len(stack) > 0 {
// 					edge := stack[len(stack)-1]
// 					stack = stack[:len(stack)-1]
// 					bcc = append(bcc, edge)
// 				}
// 				bccStack = append(bccStack, bcc)
// 			}
// 		}
// 	}

// 	return bccStack
// }

func (p Day23) rec(bc []int, gg [][]int, i int, cur []int) []int {
	ans := make([]int, len(cur))
	copy(ans, cur)
	if i == len(bc) {
		return ans
	}
	for j := i; j < len(bc); j++ {
		c := true
		for _, v := range cur {
			if gg[bc[j]][v] == 0 {
				c = false
				break
			}
		}
		if c {
			cur = append(cur, bc[j])
			cans := p.rec(bc, gg, j+1, cur)
			if len(cans) > len(ans) {
				ans = cans
			}
			cur = cur[:len(cur)-1]
		}
	}
	return ans
}

func (p Day23) Part2(lines []string) string {
	mx := 26 * 26
	gg := make([][]int, mx)
	for i := range mx {
		gg[i] = make([]int, mx)
	}
	edges := make([][]int, mx)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		nodes := strings.Split(line, "-")
		nn := []int{0, 0}
		for i := range 2 {
			for j := range 2 {
				c := int(nodes[i][j] - byte('a'))
				nn[i] = nn[i]*26 + c
			}
		}
		edges[nn[0]] = append(edges[nn[0]], nn[1])
		edges[nn[1]] = append(edges[nn[1]], nn[0])
		gg[nn[0]][nn[1]] = 1
		gg[nn[1]][nn[0]] = 1
	}
	bc := []int{}
	for i := range mx {
		if len(edges[i]) > 0 {
			bc = append(bc, i)
		}
	}
	al := p.rec(bc, gg, 0, []int{})

	ans := ""
	for i, a := range al {
		if i == 0 {
			ans = ans + p.conv2str(a)
		} else {
			ans = ans + "," + p.conv2str(a)
		}
	}
	return ans
}

func (p Day23) TestPart1() {
	const ansExample1 = "7"
	input := util.ExampleInput(2024, 23, 0)
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

func (p Day23) TestPart2() {
	const ansExample2 = "co,de,ka,ta"
	input := util.ExampleInput(2024, 23, 0)
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
