package main

import (
	"fmt"
	"golang/ch7/eval"
)

// 在func (c call) Eval(env Env) float64中
// 添加了 case:min
func main() {
	expr, _ := eval.Parse("pow(x,min(1.5,2))")
	fmt.Println(expr.Eval(eval.Env{"x": 4}))
	expr, _ = eval.Parse("log(10)")
	fmt.Println(expr.Eval(eval.Env{"x": 4}))
}
