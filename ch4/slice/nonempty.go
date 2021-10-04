package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	t := nonempty(data)
	fmt.Printf("%q\t%d\t%d\n", t, len(t), cap(t))
	fmt.Printf("%q\t%d\t%d\n", data, len(data), cap(data))
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
