package main

import (
	"fmt"

	Graph "github.com/custergo/study_algo/02.Play-with-Graph-Algorithms/utils"
)

// DFS
type GraphDFS struct {
	G       *Graph.AdjList
	visited map[int]bool // 记录每个节点是否被遍历过
	pre     []int        // 存放先序遍历结果
	post    []int        // 存放后序遍历结果
}

func NewGraphDFS(G *Graph.AdjList) *GraphDFS {
	var pre []int                    // 存放先序遍历结果
	var post []int                   // 存放后序遍历结果
	visited := make(map[int]bool, 0) // 记录每个节点是否被遍历过
	for v := 0; v < G.V; v++ {       // 遍历每个顶点，避免忽略联通分量
		if !visited[v] {
			dfs(v, G, &pre, &post, visited)
		}
	}

	return &GraphDFS{G: G, visited: visited, pre: pre, post: post}
}

func dfs(v int, G *Graph.AdjList, pre *[]int, post *[]int, visited map[int]bool) {
	visited[v] = true
	*pre = append(*pre, v) // 图的深度优先遍历的先序遍历

	for w := G.Adjacency(v).Front(); w != nil; w = w.Next() {
		ww := w.Value.(int) // 遍历顶点v的所有相邻点
		if !visited[ww] {   // 如果没有遍历过，就递归调用
			dfs(ww, G, pre, post, visited)
		}
	}

	*post = append(*post, v) // 图的深度优先遍历的后序遍历
}

func (g *GraphDFS) Pre() []int  { return g.pre }
func (g *GraphDFS) Post() []int { return g.post }

func main() {
	if g, err := Graph.NewAdjList("g.txt"); err == nil {
		dfs := NewGraphDFS(g)
		fmt.Println("先序遍历结果: ", dfs.Pre())  // 先序遍历结果: [0 1 3 2 6 4 5]
		fmt.Println("后序遍历结果: ", dfs.Post()) // 后序遍历结果: [6 2 3 4 1 0 5]
	}
}
