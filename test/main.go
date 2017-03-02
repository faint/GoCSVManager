package main

import (
	"fmt"
	"gocsv"
)

func main() {
	fmt.Println("CSVManager Start...")

	list := new(gocsv.List)
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

	params, _ := list.GetTable("params")
	// t := &params
	vvv, rrr := params.GetValuesByKey("desc")
	fmt.Println(rrr, vvv)
	fmt.Println("CSVManager End...")
}
