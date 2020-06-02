package main

type Node struct {
	x, y, height int
}

func newHeap(size int) *Heap {
	return &Heap{
		t: make([]*Node, size+1, size+1),
		p: 0,
	}
}

type Heap struct {
	t []*Node
	p int
}

func (h *Heap) Push(node *Node) {
	h.p++
	h.t[h.p] = node
	p := h.p
	for ; p > 1; {
		if h.t[p].height < h.t[p/2].height {
			h.swap(p, p/2)
			p = p / 2
		} else {
			break
		}
	}
}

func (h *Heap) Pop() *Node {
	if h.p <= 0 {
		return nil
	}
	h.swap(1, h.p)
	h.p--
	p := 1
	for ; p*2 <= h.p; {
		minP := p * 2
		if p*2+1 <= h.p && h.t[p*2+1].height < h.t[minP].height {
			minP = p*2 + 1
		}
		if h.t[minP].height < h.t[p].height {
			h.swap(p, minP)
			p = minP
		} else {
			break
		}
	}
	return h.t[h.p+1]
}

func (h *Heap) swap(x, y int) {
	tmp := h.t[x]
	h.t[x] = h.t[y]
	h.t[y] = tmp
}

func trapRainWaterCustomHeap(heightMap [][]int) int {
	n := len(heightMap)
	if n <= 2 {
		return 0
	}
	m := len(heightMap[0])
	if m <= 2 {
		return 0
	}

	visited := make([][]bool, n, n)
	for i := range visited {
		visited[i] = make([]bool, m, m)
	}

	heap := newHeap(n * m)
	for i := 0; i < m; i++ {
		heap.Push(&Node{
			x:      0,
			y:      i,
			height: heightMap[0][i],
		})
		heap.Push(&Node{
			x:      n - 1,
			y:      i,
			height: heightMap[n-1][i],
		})
		visited[0][i] = true
		visited[n-1][i] = true
	}
	for i := 1; i < n-1; i++ {
		heap.Push(&Node{
			x:      i,
			y:      0,
			height: heightMap[i][0],
		})
		heap.Push(&Node{
			x:      i,
			y:      m - 1,
			height: heightMap[i][m-1],
		})
		visited[i][0] = true
		visited[i][m-1] = true
	}

	directs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	ret := 0
	for ; ; {
		node := heap.Pop()
		if node == nil {
			break
		}
		for _, direct := range directs {
			x := node.x + direct[0]
			y := node.y + direct[1]
			if 0 <= x && x < n && 0 <= y && y < m && !visited[x][y] {
				if heightMap[x][y] < node.height {
					ret += node.height - heightMap[x][y]
					heightMap[x][y] = node.height
				}
				heap.Push(&Node{
					x:      x,
					y:      y,
					height: heightMap[x][y],
				})
				visited[x][y] = true
			}
		}
	}

	return ret
}
