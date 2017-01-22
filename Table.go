package GoCSVManager

import (
	"bytes"
	"strings"
)

// Table ...
type Table struct {
	Name  string
	Keys  *Key // 标题行
	Lines []Line
}

// createTable ...
func createTable(name string, fileBytes []byte) Table {
	// init content
	newTable := Table{}
	newTable.Name = name
	newTable.Lines = []Line{}
	lines := bytes.Split(fileBytes, []byte{'\n'})
	for _, v := range lines { // parse csv
		if len(v) <= 0 { // 空行过滤
			continue
		}

		if v[len(v)-1] == '\r' { // 去除尾部\r
			v = v[:len(v)-1]
		}

		if len(v) <= 0 { // 去除仅包含`\n\r`的空行
			continue
		}

		if v[0] == '#' { // 处理注释行
			if len(v) > 1 { // 处理非仅#的注释行
				if v[1] == '!' && newTable.Keys != nil { // 在标题行未处理过的情况下，!视为标题行处理，否则不处理
					newKeys := new(Key)
					newKeys.Value = strings.Split(string(v[2:]), ",")
					newTable.Keys = newKeys
				}
				continue
			}
		}

		line := createLine(newTable.Keys, strings.Split(string(v), ","))
		newTable.Lines = append(newTable.Lines, line)
	}

	return newTable
}

// GetLine ...
func (table *Table) GetLine(keyName, keyValue string) (Line, bool) {
	n, result := table.Keys.GetIndex(keyName)
	if !result {
		return Line{}, false
	}

	for _, v := range table.Lines {
		if v.Values[n] == keyValue {
			return v, true
		}
	}

	return Line{}, false
}

// GetLines return multip line
func (table *Table) GetLines(keyName, keyValue string) ([]Line, bool) {
	n, result := table.Keys.GetIndex(keyName)
	if !result { // 没有这个keyName
		return []Line{}, false
	}

	lines := []Line{}
	for _, v := range table.Lines {
		if v.Values[n] == keyValue {
			lines = append(lines, v)
		}
	}

	return lines, true
}

// GetN Get Element N
func (table *Table) GetN(n int) (Line, bool) {
	if len(table.Lines) >= (n - 1) {
		return table.Lines[n], true
	}

	return Line{}, false
}
