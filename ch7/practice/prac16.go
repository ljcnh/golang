package main

import (
	"golang/ch7/eval"
	"html/template"
	"log"
	"net/http"
)

var calc = template.Must(template.New("calc").Parse(`
<h1>计算器</h1>
<form method="get" action="/">
  <input type="text" name="expr" value="{{.ExprStr}}"/>
  <input type="submit" value="calc"/>
  <p>{{.Result}}</p>
</form>
`))

type Res struct {
	ExprStr string
	Result  float64
}

func main() {
	myHandle := func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		res := &Res{"", 0.0}
		if len(q["expr"]) > 0 {
			res.ExprStr = q["expr"][0]
			expr, err := eval.Parse(q["expr"][0])
			if err == nil {
				res.Result = expr.Eval(eval.Env{})
			}
		}
		calc.Execute(w, res)
	}
	http.HandleFunc("/", myHandle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
