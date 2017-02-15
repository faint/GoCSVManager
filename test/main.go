package main

import (
	"fmt"

	"../../goCsv"
)

func main() {
	fmt.Println("CSVManager Start...")

	list := new(goCsv.List)
	list.Load("../csv/params.csv")
	list.Load("../csv/banwords.csv")

	// fmt.Println(list.Tables)
	v, f := list.GetValueByFiled("params", "name", "BattleConfiguration", "desc")
	fmt.Println("v:", v, f)

	vs, fs := list.GetValuesByFiled("params", "name", "BattleConfiguration", "desc")
	fmt.Println("vs:", vs, fs)

	n, nbool := list.GetFirstValueByN("banwords", 5)
	fmt.Println("n:", n, nbool)

	table, _ := list.GetTable("banwords")
	fmt.Println("Ban Line:", table.Size)
	fmt.Println("CSVManager End...")
}
