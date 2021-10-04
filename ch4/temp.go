package main

import (
	"fmt"
	//"go/ch4/myJson"
	"golang/ch4/myJson/practice/prac10"
	"log"
	"os"
)

func main() {
	//myJson.SearchIssues(os.Args[1:])
	result, err := prac10.SearchIssuesLimited(os.Args[1:], -1, 0, 0, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%v \t", item.CreatedAt)
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User, item.Title)
	}
	//now := time.Now()
	//fmt.Println(now)
	//m, _ := time.ParseDuration("-1d")
	//fmt.Println(m)
	//m1 := now.Add(m)
	//fmt.Println(m1)
}

/*


const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)
type Employee myStruct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}
var dilbert Employee
type Point struct{ X, Y int }


p := Point{1, 2}
q := Point{2, 1}
fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
fmt.Println(p == q)
hits := make(map[Point]int)
hits[Point{23, 443}]++
fmt.Println(hits)
seen := make(map[string]struct{})
if _, ok := seen["s"]; !ok {
	seen["s"] = struct{}{}
	// ...first time seeing s...
}
fmt.Println(seen)
values := []int{21, 45, 5, 4, 3, 7, 23, 576, 3, 45, 6, 24, 25, 5, 6, 36}
myStruct.TreeSort(values)
fmt.Println(values)

	dilbert.Salary -= 5000
	fmt.Println(dilbert.Salary)
	position := &dilbert.Position
	*position = "Senior" + *position
	fmt.Println(dilbert.Position)
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += "(proactive team player)"
	fmt.Println(dilbert.Position)

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
	id := dilbert.ID
	EmployeeByID(id).Salary = 0*/
/*func EmployeeByID(id int) *Employee {
	return &dilbert
}*/
/*
func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

		months := [...]string{1: "January", 2: "February", 3: "Mar", 4: "Apr", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "Dece", 12: "December"}
		summer := months[6:9]
		fmt.Println(len(summer))
		fmt.Println(summer)
		endlessSummer := summer[:5] // extend a slice (within capacity)
		fmt.Println(len(endlessSummer))
		fmt.Println(endlessSummer)


		months := [...]string{1: "January", 2: "February", 3: "Mar", 4: "Apr", 5: "May", 6: "June", 7: "July", 8: "August", 9: "Swp", 12: "December"}

		Q2 := months[4:7]
		summer := months[6:9]
			fmt.Println(Q2)     // ["April" "May" "June"]
		fmt.Println(summer) // ["June" "July" "August"]
		for _, s := range summer {
			for _, q := range Q2 {
				if s == q {
					fmt.Printf("%s appears in both\n", s)
				}
			}
		}
		fmt.Println(summer[:20])
		symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
		fmt.Println(len(symbol))
		symbol1 := [...]string{"$", "€", "￡", "￥"}
		fmt.Println(symbol1 == symbol)
		symbol2 := [...]string{0: "$", 3: "€", 2: "￡", 1: "￥"}
		fmt.Println(symbol2[1])
		symbol3 := [...]string{0: "$", 3: "€", 8: "￡", 1: "￥"}
		fmt.Println(len(symbol3))
		var q [5]int = [5]int{1, 2, 3}
		var r [3]int = [3]int{1, 3}
		fmt.Println(r[2])              // "0"
		fmt.Println(r[1])              // "0"
		fmt.Println(reflect.TypeOf(q)) // "0"
}
*/
