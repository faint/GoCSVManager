package main

import (
	"fmt"

	"../../GoCSVManager"
)

func main() {
	fmt.Println("CSVManager Start...")

	list := new(GoCSVManager.List)
	list.Load("../csv/params.csv")

	fmt.Println(list.Tables)
	v, f := list.GetValueByFiled("params", "name", "BattleConfiguration", "desc")
	fmt.Println("v:", v, f)

	vs, fs := list.GetValuesByFiled("params", "name", "BattleConfiguration", "desc")
	fmt.Println("vs:", vs, fs)

	fmt.Println("CSVManager End...")
}
