package main

import (
	"bufio"
	"fmt"
	"golang/ch7/eval"
	"os"
	"strconv"
	"strings"
)

func main() {
	exitCode := 0
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("表达式: ")
	stdin.Scan()
	expr, err := eval.Parse(stdin.Text())
	if err != nil {
		fmt.Fprintf(os.Stderr, "表达式错误: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("输入 exit 退出\n")
	fmt.Printf("输入变量的值 ( <var>=<val>, 如: x=3 y=5 ...): ")

	env := eval.Env{}
	for true {
		stdin.Scan()
		envStr := stdin.Text()
		if envStr == "exit" {
			fmt.Printf("退出")
			os.Exit(0)
		}
		if stdin.Err() != nil {
			fmt.Fprintln(os.Stderr, stdin.Err())
			os.Exit(1)
		}
		assignments := strings.Fields(envStr)
		for _, a := range assignments {
			fields := strings.Split(a, "=")
			if len(fields) != 2 {
				fmt.Fprintf(os.Stderr, "输入 变量=值 错误: %s\n", a)
				exitCode = 2
			}
			ident, valStr := fields[0], fields[1]
			val, err := strconv.ParseFloat(valStr, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "变量值 %s 解析错误, 默认使用0: %s\n", ident, err)
				exitCode = 2
			}
			env[eval.Var(ident)] = val
		}
		fmt.Println(expr.Eval(env))
		fmt.Printf("exitcode：%d\n", exitCode)
	}
}
