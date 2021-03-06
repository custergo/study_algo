package _4_Query_in_Segment_Tree

import (
	"bytes"
	"fmt"
)

type SegmentTree struct {
	data   []interface{}
	tree   []interface{}
	merger func(interface{}, interface{}) interface{}
}

func NewSegmentTree(arr []interface{}, merger func(interface{}, interface{}) interface{}) *SegmentTree {
	segmentTree := &SegmentTree{
		data:   make([]interface{}, len(arr)),
		tree:   make([]interface{}, len(arr)*4),
		merger: merger,
	}
	for i := 0; i < len(arr); i++ {
		segmentTree.data[i] = arr[i]
	}
	segmentTree.buildSegmentTree(0, 0, len(arr)-1)
	return segmentTree
}

// 在treeIndex的位置创建表示区间[l...r]的线段树
// treeIndex: 创建的线段树根节点索引, l,r: 该节点表示的区间
func (st *SegmentTree) buildSegmentTree(treeIndex, l, r int) {
	if l == r { // 区间长度为1，只有一个元素
		st.tree[treeIndex] = st.data[l] // 该节点存储的就是元素本身
		return
	}

	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildSegmentTree(leftTreeIndex, l, mid)
	st.buildSegmentTree(rightTreeIndex, mid+1, r)

	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentTree) GetSize() int {
	return len(st.data)
}

func (st *SegmentTree) Get(index int) interface{} {
	if index < 0 || index >= len(st.data) {
		panic("Index is illegal.")
	}
	return st.data[index]
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func rightChild(index int) int {
	return index*2 + 2
}

// 返回区间[queryL, queryR]的值
func (st *SegmentTree) Query(queryL, queryR int) interface{} {
	if queryL < 0 || queryL >= len(st.data) ||
		queryR < 0 || queryR >= len(st.data) {
		panic("Index is illegal.")
	}
	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

// 在以treeID为根的线段树中[l...r]的范围里，搜索区间[queryL, queryR]的值
func (st *SegmentTree) query(treeIndex, l, r, queryL, queryR int) interface{} {
	if l == queryL && r == queryR {
		return st.tree[treeIndex]
	}
	mid := l + (r-l)/2
	// treeIndex的节点分为[l...mid]和[mid+1...r]两部分
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)

	if queryL >= mid+1 {
		return st.query(rightTreeIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid {
		return st.query(leftTreeIndex, l, mid, queryL, queryR)
	}

	leftResult := st.query(leftTreeIndex, l, mid, queryL, mid)
	rightResult := st.query(rightTreeIndex, mid+1, r, mid+1, queryR)

	return st.merger(leftResult, rightResult)
}

func (st *SegmentTree) String() string {
	buffer := bytes.Buffer{}

	buffer.WriteString("[")
	for i := 0; i < len(st.tree); i++ {
		if st.tree[i] != nil {
			buffer.WriteString(fmt.Sprint(st.tree[i]))
		} else {
			buffer.WriteString("nil")
		}

		if i != len(st.tree)-1 {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("]")
	return buffer.String()
}
