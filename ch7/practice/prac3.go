package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func TreeSort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (c *tree) String() string {
	var res string
	var dfs func(c *tree)
	// 递归 中序遍历
	dfs = func(c *tree) {
		if c == nil {
			return
		}
		dfs(c.left)
		res = res + " " + strconv.Itoa(c.value)
		dfs(c.right)
	}
	dfs(c)
	return res
}

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	var root *tree
	for _, v := range data {
		root = add(root, v)
	}

	fmt.Println(root)
}
