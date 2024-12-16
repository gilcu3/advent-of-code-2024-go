package year2024

import (
	"container/heap"
	"fmt"

	"aocgen/internal/util"
)

type Day16 struct{}

type Item struct {
	value    []int // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	index    int   // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// // update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, value []int, priority int) {
// 	item.value = value
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }

func (p Day16) Part1(lines []string) string {
	sig := ".#SE"
	dd := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	ar := make([][]int, 0)
	var sx, sy, ex, ey int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		row := make([]int, len(line))
		for i, c := range line {
			for j, cc := range sig {
				if c == cc {
					row[i] = j
					break
				}
			}
			if row[i] == 2 {
				sx, sy = len(ar), i
				row[i] = 0
			} else if row[i] == 3 {
				ex, ey = len(ar), i
				row[i] = 0
			}
		}
		ar = append(ar, row)
	}
	n, m := len(ar), len(ar[0])
	dist := make([][][]int, n)
	for i := range n {
		dist[i] = make([][]int, m)
		for j := range m {
			dist[i][j] = make([]int, 4)
			for d := range 4 {
				dist[i][j][d] = -1
			}
		}
	}

	var ans int

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: []int{sx, sy, 1}, priority: 0})
	dist[sx][sy][1] = 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		cx, cy, cd, dis := item.value[0], item.value[1], item.value[2], -item.priority
		if dist[cx][cy][cd] < dis {
			continue
		}
		// fmt.Println(cx, cy, cd, dis)
		if cx == ex && cy == ey {
			ans = dis
			break
		}
		nx, ny, nd, ndd := cx, cy, (cd+3)%4, dis+1000
		if dist[nx][ny][nd] == -1 || (dist[nx][ny][nd] > ndd) {
			heap.Push(&pq, &Item{value: []int{nx, ny, nd}, priority: -ndd})
			dist[nx][ny][nd] = ndd
		}
		nx, ny, nd, ndd = cx, cy, (cd+1)%4, dis+1000
		if dist[nx][ny][nd] == -1 || (dist[nx][ny][nd] > ndd) {
			heap.Push(&pq, &Item{value: []int{nx, ny, nd}, priority: -ndd})
			dist[nx][ny][nd] = ndd
		}
		nx, ny, nd, ndd = cx+dd[cd][0], cy+dd[cd][1], cd, dis+1
		if 0 <= nx && nx < n && 0 <= ny && ny < m && ar[nx][ny] == 0 && (dist[nx][ny][nd] == -1 || (dist[nx][ny][nd] > ndd)) {
			heap.Push(&pq, &Item{value: []int{nx, ny, nd}, priority: -ndd})
			dist[nx][ny][nd] = ndd
		}
	}
	return fmt.Sprint(ans)
}

func dijkstra(sx, sy, ex, ey int, ar [][]int, sd int) [][][]int {
	dd := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	n, m := len(ar), len(ar[0])
	dist := make([][][]int, n)
	for i := range n {
		dist[i] = make([][]int, m)
		for j := range m {
			dist[i][j] = make([]int, 4)
			for d := range 4 {
				dist[i][j][d] = -1
			}
		}
	}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	if sd == -1 {
		for d := range 4 {
			heap.Push(&pq, &Item{value: []int{sx, sy, d}, priority: 0})
			dist[sx][sy][d] = 0
		}
	} else {
		heap.Push(&pq, &Item{value: []int{sx, sy, sd}, priority: 0})
		dist[sx][sy][sd] = 0
	}
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		cx, cy, cd, dis := item.value[0], item.value[1], item.value[2], -item.priority
		if dist[cx][cy][cd] < dis {
			continue
		}
		// fmt.Println(cx, cy, cd, dis)
		if cx == ex && cy == ey {
			return dist
		}

		nx, ny, nd, ndd := cx, cy, (cd+3)%4, dis+1000
		if dist[nx][ny][nd] == -1 || (dist[nx][ny][nd] > ndd) {
			heap.Push(&pq, &Item{value: []int{nx, ny, nd}, priority: -ndd})
			dist[nx][ny][nd] = ndd
		}
		nx, ny, nd, ndd = cx, cy, (cd+1)%4, dis+1000
		if dist[nx][ny][nd] == -1 || (dist[nx][ny][nd] > ndd) {
			heap.Push(&pq, &Item{value: []int{nx, ny, nd}, priority: -ndd})
			dist[nx][ny][nd] = ndd
		}
		nx, ny, nd, ndd = cx+dd[cd][0], cy+dd[cd][1], cd, dis+1
		if 0 <= nx && nx < n && 0 <= ny && ny < m && ar[nx][ny] == 0 && (dist[nx][ny][nd] == -1 || (dist[nx][ny][nd] > ndd)) {
			heap.Push(&pq, &Item{value: []int{nx, ny, nd}, priority: -ndd})
			dist[nx][ny][nd] = ndd
		}
	}
	return dist
}

func (p Day16) Part2(lines []string) string {
	sig := ".#SE"
	ar := make([][]int, 0)
	var sx, sy, ex, ey int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		row := make([]int, len(line))
		for i, c := range line {
			for j, cc := range sig {
				if c == cc {
					row[i] = j
					break
				}
			}
			if row[i] == 2 {
				sx, sy = len(ar), i
				row[i] = 0
			} else if row[i] == 3 {
				ex, ey = len(ar), i
				row[i] = 0
			}
		}
		ar = append(ar, row)
	}
	n, m := len(ar), len(ar[0])
	dist := dijkstra(sx, sy, ex, ey, ar, 1)
	dd := -1
	for d := range 4 {
		if dist[ex][ey][d] != -1 && (dd == -1 || dd > dist[ex][ey][d]) {
			dd = dist[ex][ey][d]
		}
	}
	distr := dijkstra(ex, ey, sx, sy, ar, -1)
	ans := 0
	for i := range n {
		for j := range m {
			for d := range 4 {
				if dist[i][j][d] != -1 && distr[i][j][(d+2)%4] != -1 && dist[i][j][d]+distr[i][j][(d+2)%4] == dd {
					ans += 1
					break
				}
			}

		}
	}
	return fmt.Sprint(ans)
}

func (p Day16) TestPart1() {
	const ansExample1 = "11048"
	input := util.ExampleInput(2024, 16, 0)
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

func (p Day16) TestPart2() {
	const ansExample2 = "64"
	input := util.ExampleInput(2024, 16, 0)
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
