package year2024

import (
	"container/heap"
	"fmt"
	"strconv"

	"aocgen/internal/util"
)

type Day21 struct{}

type path struct {
	x byte
	y byte
}

type Item21 struct {
	value    []int // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	index    int   // The index of the item in the heap.
}

type PriorityQueue21 []*Item21

func (pq PriorityQueue21) Len() int { return len(pq) }

func (pq PriorityQueue21) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue21) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue21) Push(x any) {
	n := len(*pq)
	item := x.(*Item21)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue21) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

/*
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
    | 0 | A |
    +---+---+

    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/

func (p Day21) rec1(from, to byte, i, top int, tab0, tab1 []string, memo []map[path]int) int {
	if i == top {
		return 1
	}
	if v, ok := memo[i][path{from, to}]; ok {
		return v
	}
	var tab []string
	if i == 0 {
		tab = tab0
	} else {
		tab = tab1
	}
	n, m := len(tab), len(tab[0])
	var sx, sy, ex, ey int
	for i := range n {
		for j := range m {
			if tab[i][j] == from {
				sx, sy = i, j
			}
			if tab[i][j] == to {
				ex, ey = i, j
			}
		}
	}
	dist := make([][][]int, n)
	for i := range n {
		dist[i] = make([][]int, m)
		for j := range m {
			dist[i][j] = make([]int, 5)
			for k := range 5 {
				dist[i][j][k] = -1
			}
		}
	}
	pq := make(PriorityQueue21, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item21{value: []int{sx, sy, 4}, priority: 0})
	dist[sx][sy][4] = 0
	dd := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	moves := "><v^A"
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item21)

		cx, cy, last, dis := item.value[0], item.value[1], item.value[2], -item.priority
		if dist[cx][cy][last] < dis {
			continue
		}
		if cx == ex && cy == ey && last == 4 {
			// fmt.Println(i, string(from), string(to), dis)
			if dis == 0 {
				dis = 1
			}
			memo[i][path{from, to}] = dis
			return dis
		}
		for d := range 4 {
			nx, ny := cx+dd[d][0], cy+dd[d][1]
			if 0 <= nx && nx < n && 0 <= ny && ny < m && tab[nx][ny] != ' ' {
				nm := p.rec1(moves[last], moves[d], i+1, top, tab0, tab1, memo)
				if dist[nx][ny][d] == -1 || dist[nx][ny][d] > dis+nm {
					dist[nx][ny][d] = dis + nm
					heap.Push(&pq, &Item21{value: []int{nx, ny, d}, priority: -dist[nx][ny][d]})
				}

			}
		}
		if cx == ex && cy == ey {
			nm := p.rec1(moves[last], moves[4], i+1, top, tab0, tab1, memo)
			if dist[cx][cy][4] == -1 || dist[cx][cy][4] > dis+nm {
				dist[cx][cy][4] = dis + nm
				heap.Push(&pq, &Item21{value: []int{cx, cy, 4}, priority: -dist[cx][cy][4]})
			}
		}
	}
	panic("should not get here")
}

func (p Day21) Part1(lines []string) string {
	tab0 := []string{"789", "456", "123", " 0A"}
	tab1 := []string{" ^A", "<v>"}
	ans := 0
	steps := 2 + 1
	memo := make([]map[path]int, steps)
	for i := range steps {
		memo[i] = make(map[path]int)
	}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		last := byte('A')
		cur := 0
		mul, _ := strconv.Atoi(line[:len(line)-1])
		for i := range line {
			cur += p.rec1(last, line[i], 0, steps, tab0, tab1, memo)
			last = line[i]
		}
		// fmt.Println(cur, mul)
		ans += cur * mul
	}
	return fmt.Sprint(ans)
}

func (p Day21) Part2(lines []string) string {
	tab0 := []string{"789", "456", "123", " 0A"}
	tab1 := []string{" ^A", "<v>"}
	ans := 0
	steps := 25 + 1
	memo := make([]map[path]int, steps)
	for i := range steps {
		memo[i] = make(map[path]int)
	}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		last := byte('A')
		cur := 0
		mul, _ := strconv.Atoi(line[:len(line)-1])
		for i := range line {
			cur += p.rec1(last, line[i], 0, steps, tab0, tab1, memo)
			last = line[i]
		}
		// fmt.Println(cur, mul)
		ans += cur * mul
	}
	return fmt.Sprint(ans)
}

func (p Day21) TestPart1() {
	const ansExample1 = "126384"
	input := util.ExampleInput(2024, 21, 0)
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

func (p Day21) TestPart2() {
	const ansExample2 = "154115708116294"
	input := util.ExampleInput(2024, 21, 0)
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
