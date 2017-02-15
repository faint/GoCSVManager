package goCsv

import (
	"io/ioutil"
	"path"
	"strings"
)

// List CSV文件列表
type List struct {
	Tables []Table // CSV文件表格数组
}

// Load 加载CSV文件：
// 当文件已经存在List结构内时，重新读取，更新既有内容。
// 当List结构未保存该文件时时，读取并加入List结构。
func (list *List) Load(pathAndFilename string) error {
	file, e := ioutil.ReadFile(pathAndFilename)
	if e != nil {
		return e
	}

	filename := strings.Split(path.Base(pathAndFilename), ".")[0]
	tableNew := createTable(filename, file)

	// 检查是否已存在同名csv
	n, exist := list.isExist(filename)
	if exist { // 如果存在，则更新列表里的CSV
		list.Tables[n] = tableNew
	} else { // 如果不存在，则创建新的CSV，并加入列表
		list.Tables = append(list.Tables, tableNew)
	}

	return nil
}

// isExist 检查指定CSV表名是否已经在结构内
// 返回(结构内的索引值, 存在性的布尔值)
func (list *List) isExist(name string) (int, bool) {
	for k, v := range list.Tables {
		if v.Name == name {
			return k, true
		}
	}

	return 0, false
}

// GetTable 返回(指定表名的Table结构,存在性的真假值)
func (list *List) GetTable(csvName string) (Table, bool) {
	for _, v := range list.Tables {
		if v.Name == csvName {
			return v, true
		}
	}

	return Table{}, false
}

// GetValueByFiled return only one line
func (list *List) GetValueByFiled(csvName, keyFiledName, keyFiledValue, needField string) (string, bool) {
	table, result := list.GetTable(csvName) //  读表
	if !result {
		return "", false
	}

	line, result := table.GetLine(keyFiledName, keyFiledValue) // 读行
	if !result {
		return "", false
	}

	value, result := line.GetValueBy(needField)
	if !result {
		return "", false
	}

	return value, true
}

// GetValuesByFiled return multi line
func (list *List) GetValuesByFiled(csvName, keyFiledName, keyFiledValue, needField string) ([]string, bool) {
	table, result := list.GetTable(csvName) //  读表
	if !result {
		return []string{}, false
	}

	lines, result := table.GetLines(keyFiledName, keyFiledValue)
	if !result {
		return []string{}, false
	}

	values := []string{}
	for _, v := range lines {
		value, result := v.GetValueBy(needField)
		if result {
			values = append(values, value)
		}
	}

	return values, true
}

// GetValueByN single column
func (list *List) GetValueByN(csvName string, n int) (Line, bool) {
	table, result := list.GetTable(csvName) //  读表
	if !result {
		return Line{}, false
	}

	line, result := table.GetN(n)
	if !result {
		return Line{}, false
	}

	return line, true
}

// GetFirstValueByN ..
func (list *List) GetFirstValueByN(csvName string, n int) (string, bool) {
	line, result := list.GetValueByN(csvName, n)
	if !result {
		return "", false
	}

	return line.Values[0], true
}
