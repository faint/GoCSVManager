package main

import (
	"fmt"
	"gocsv"
)

func main() {
	fmt.Println("检测CSV读取合法性")
	list := new(gocsv.List)

	n, err := list.LoadDir("./csv/")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("read successed:", n)
}
