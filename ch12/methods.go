package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// 显示一个类型的方法集

func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)
	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}
}

func main() {
	Print(time.Hour)
	Print(new(strings.Replacer))
}
