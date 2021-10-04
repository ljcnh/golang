package main

func main() {
	//ages := make(map[string]int)
	/*	ages := map[string]int{
			"alice":   31,
			"charlie": 34,
		}
		ages["bob"] = ages["bob"] + 1
		//fmt.Println(ages)
		for name, age := range ages {
			fmt.Printf("%s\t%d\n", name, age)
		}
		var names []string
		names := make([]string, 0, len(ages))
		for name := range ages {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Printf("%s\t%d\n", name, ages[name])
		}
		var ages map[string]int

		if age, ok := ages["bob"]; !ok {
			fmt.Println(age)
		}
		fmt.Println(equal(map[string]int{"A": 42}, map[string]int{"A": 42}))*/
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
