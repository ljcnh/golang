package main

import (
	"fmt"
	"golang/ch7/eval"
)

// Expr合约：添加的 String 方法
// strings 文件
// 这里只是个测试
func main() {
	expr, _ := eval.Parse("pow(2 + pow(3))")
	fmt.Println(expr.String())
}
