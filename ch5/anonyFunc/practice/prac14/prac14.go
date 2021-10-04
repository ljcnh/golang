package main

import "fmt"

var prereqs = map[string][]string{
	"root": {
		"algorithms",
		"calculus",
		"compilers",
		"databases",
		"networks",
		"programming languages",
	},
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	s := breadthFirst(prereqs, "compilers")
	fmt.Println(s)
	fmt.Println(len(s))
}

func breadthFirst(m map[string][]string, start string) []string {
	var ans []string

	seen := make(map[string]bool)
	seen[start] = true

	var que []string
	que = append(que, start)

	for len(que) > 0 {
		s := que[0]
		que = que[1:]
		ans = append(ans, s)
		for _, t := range m[s] {
			if !seen[t] {
				seen[t] = true
				que = append(que, t)
			}
		}
	}

	return ans
}
