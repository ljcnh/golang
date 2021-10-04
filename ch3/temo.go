package main

import (
	"fmt"
	"reflect"
)

type Weekday int
type Flags uint

func main() {
	var f float64 = 3 + 0i
	f = 2
	f = 1e123
	var a = 1 + 0i
	a = 1e123
	fmt.Println(f, a)
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(a))
	var b int32 = 1
	var c int64 = 34
	var d = 234
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(d))
	fmt.Println(b & int32(c))
	/*	const (
			sundau Weekday = iota
			Monday
			Tuesday
			Wednesday
			Thursday
			Friday
			Saturday
		)
		fmt.Println(sundau, Monday, Tuesday, Wednesday, Tuesday, Friday, Saturday)
		const (
			FlagUp           Flags = 1 << iota // is up
			FlagBroadcast                      // supports broadcast access capability
			FlagLoopback                       // is a loopback interface
			FlagPointToPoint                   // belongs to a point-to-point link
			FlagMulticast                      // supports multicast access capability
		)
		fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)*/

	/*	const noDelay time.Duration = 0
		const timeout = 5 * time.Minute
		fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
		fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
		fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

		x := 123
		y := fmt.Sprintf("%d", x)
		z := strconv.Itoa(x)
		fmt.Println(y, z) // "123 123"

		fmt.Println(reflect.TypeOf(y) == reflect.TypeOf(z))

		a := strconv.FormatInt(int64(x), 2)
		fmt.Println(a)
		fmt.Println(reflect.TypeOf(a))
		fmt.Println(reflect.TypeOf(z))
		s := fmt.Sprintf("x=%b", x) // "x=1111011"
		fmt.Println(s)
		fmt.Println(reflect.TypeOf(s))

		x, err := strconv.Atoi("123")
		fmt.Println(x)
		fmt.Println(reflect.TypeOf(x))
		if err != nil {
			fmt.Println(err)
		}
		y, err := strconv.ParseInt("123", 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(y)
		fmt.Println(reflect.TypeOf(y))*/
	/*	var x uint8 = 1 << 1 | 1 << 5
		var y uint8 = 1 << 1 | 1 << 2
		fmt.Printf("%08b\n",x)   //00100010
		fmt.Printf("%08b\n",y)   //00000110
		fmt.Printf( "%08b\n" , x&y)   //00000010
		fmt.Printf( "%08b\n" , x|y)   //00100110
		fmt.Printf( "%08b\n" , x^y)   //00100100
		fmt.Printf( "%08b\n" , x&^y)   //00100000

		for i := uint(0);i<8;i++ {
			if x&(1<<i) != 0{
				fmt.Println(i)  // 1 5
			}
		}
		fmt.Printf( "%08b\n" , x<< 1 ) // "01000100", the set {2, 6}
		fmt.Printf( "%08b\n" , x>> 1 ) // "00010001", the set {0, 4}*/
	/*	medals := []string { "gold" , "silver" , "bronze"}
		for i:=len(medals)-1;i>=0;i--{
			fmt.Println(medals[i])
		}
		o := 0666
		fmt.Printf( "%d %[1]o %#[1]o\n" , o) // "438 666 0666"
		x := int64 ( 0xdeadbeef )
		fmt.Printf( "%d %[1]x %#[1]x %#[1]X\n" , x)
		ascii := 'a'
		unicode := '国'
		newline := '\n'
		fmt.Printf( "%d %[1]c %[1]q\n" , ascii) // "97 a 'a'"
		fmt.Printf( "%d %[1]c %[1]q\n" , unicode) // "22269 国 '国'"
		fmt.Printf( "%d %[1]q\n" , newline) // "10 '\n'
		for x := 0 ; x < 8 ; x++ {
			fmt.Printf( "x = %d e^x = %8.3f\n" , x, math.Exp( float64 (x)))
		}
		var z float64
		fmt.Println(z, -z, 1 /z, - 1 /z, z/z) // "0 -0 +Inf -Inf NaN"
		nan := math.NaN()
		fmt.Println(nan == nan, nan < nan, nan > nan)*/
	/*	s := "Hello, 世界"
		for i := 0; i < len(s); {
			r, size := utf8.DecodeRuneInString(s[i:])
			fmt.Printf("%d\t%c\t%d\n", i, r, size)
			i += size
		}
		fmt.Println(len(s))
		for i := 0; i < len(s); i++ {
			fmt.Println(s[i])
		}
		for i, r := range s {
			fmt.Printf("%d\t%c\n", i, r)
		}
		s := "Hello, 世界"
		n := 0
		for range s {
			n++
		}
		fmt.Println(n)
		fmt.Println(len(s))*/
}
