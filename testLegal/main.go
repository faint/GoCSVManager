package main

import (
	"fmt"
	"gocsv"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("检测CSV读取合法性")
	list := new(gocsv.List)

	dirList, e := ioutil.ReadDir("../csv/")
	if e != nil {
		fmt.Println("read dir error")
		return
	}

	count := 0
	for _, v := range dirList {
		fmt.Println("加载：", v.Name())

		e := list.Load("../csv/" + v.Name())
		if e != nil {
			fmt.Println(e)
		}

		csvName := strings.Split(v.Name(), ".")[0]
		fmt.Println("csvName:", csvName)
		table, result := list.GetTable(csvName)
		if !result {
			fmt.Println(csvName, ":", result)
			count++
		}

		fmt.Println("测试Key：", table.Keys)

	}
	fmt.Println("检测完毕，共有", count, "文件读取失败！!")
}
