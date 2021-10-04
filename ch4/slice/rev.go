package main

func main() {
	/*	a := []int{0, 1, 2, 3, 4, 5}
		reverse(a[:])
		fmt.Println(a)
		var b [6]int = [6]int{1, 2, 3, 4, 5, 6}
		reverse(b[:])
		fmt.Println(b)
		fmt.Println(cap(b))*/
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
