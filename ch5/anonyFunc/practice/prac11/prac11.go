package main

import "fmt"

var prereqs = map[string][]string{
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
	"linear algebra":        {"calculus"},
}

func main() {
	a, b := topoSort(prereqs)
	fmt.Println(b)
	if b {
		for i, s := range a {
			println(i, s)
		}
	}
}

func topoSort(m map[string][]string) ([]string, bool) {
	var ans []string
	mp := make(map[string]map[string]bool)
	in := make(map[string]int)
	seen := make(map[string]bool)

	for _, value := range m {
		for _, val := range value {
			mp[val] = make(map[string]bool)
		}
	}

	for key, value := range m {
		for _, val := range value {
			mp[val][key] = true
			_, ok := in[val]
			if !ok {
				in[val] = 0
			}
		}
		in[key] = len(value)
	}

	f := false
	for {
		f = false
		for key, inNums := range in {
			if inNums == 0 && !seen[key] {
				f = true
				ans = append(ans, key)
				seen[key] = true
				for ik := range mp[key] {
					in[ik]--
				}
			}
		}
		if len(ans) == len(in) {
			f = true
			break
		}
		if !f {
			break
		}
	}
	if !f {
		return nil, false
	} else {
		return ans, true
	}
}
