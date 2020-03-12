package main

import jsoniter "github.com/json-iterator/go"

func main() {
	a := []ExprNode{}
	for i := 0; i < 10; i++ {
		a = append(a, ExprNode{
			Left: i,
		})
	}
}

type ExprNode struct {
	Left          uint8  `json:"left"`
	CompOperator  uint8  `json:"comp_operator"`
	Value         string `json:"value"`
	LogicOperator string `json:"logic_operator"`
}

type Expr []ExprNode
