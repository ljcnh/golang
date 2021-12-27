package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())
	fd := stdout.FieldByName("fd")
	fmt.Println(fd.Int()) // "1"
	//fd.SetInt(2)                           //
	fmt.Println(fd.CanAddr(), fd.CanSet()) // "true false"

	//x := 2                   // value   type    variable?
	//a := reflect.ValueOf(2)  // 2       int     no
	//fmt.Println(a)           // "3"
	//b := reflect.ValueOf(x)  // 2       int     no
	//fmt.Println(b)           // "3"
	//c := reflect.ValueOf(&x) // &x      *int    no
	//fmt.Println(c)           // "3"
	//d := c.Elem()            // 2       int     yes (x)
	//fmt.Println(d)           // "3"
	//
	//fmt.Println(a.CanAddr()) // "false"
	//fmt.Println(b.CanAddr()) // "false"
	//fmt.Println(c.CanAddr()) // "false"
	//fmt.Println(d.CanAddr()) // "true"
	//
	//x := 2
	//fmt.Println(&x)
	//d := reflect.ValueOf(&x).Elem()
	//c := reflect.ValueOf(&x)
	//fmt.Println(&d)
	//fmt.Println(&c)
	//
	//px := d.Addr().Interface().(*int)
	//fmt.Println(px)
	//*px = 3
	//fmt.Println(x)
	//d.Set(reflect.ValueOf(6))
	//fmt.Println(x)
	//v := reflect.ValueOf(3) // a reflect.Value
	//x := v.Interface()      // an interface{}
	//fmt.Println(x)          // "3"
	//i := x.(int)            // an int
	//fmt.Printf("%d\n", i)   // "3"
	//
	//t := v.Type()
	//fmt.Println(t)          // "int"
	//fmt.Println(t.String()) // "int"

}
