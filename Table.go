package gocsv

import (
	"bytes"
	"errors"
	"regexp"
	"strings"
)

const (
	keyNotFound    = "key not found in csv:"
	targetNotFound = "target not found in csv"
)

// Table ...
type Table struct {
	Name  string
	Size  int
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
				// 在标题行未处理过的情况下，!视为键行处理，否则不处理
				if v[1] == '!' && newTable.Keys == nil {
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
	newTable.Size = len(newTable.Lines)

	return newTable
}

// GetLine ...
func (table *Table) GetLine(keyName, keyValue string) (Line, bool) {
	if table.Keys == nil {
		return Line{}, false
	}

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
	if table.Keys == nil {
		return []Line{}, false
	}

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

// MatchLine ...
func (table *Table) MatchLine(keyName, matchValue string) (Line, bool) {
	if table.Keys == nil {
		return Line{}, false
	}

	n, result := table.Keys.GetIndex(keyName)
	if !result {
		return Line{}, false
	}

	for _, v := range table.Lines {
		matched, e := regexp.MatchString(matchValue, v.Values[n])
		if e != nil {
			return Line{}, false
		}

		if matched {
			return v, true
		}
	}

	return Line{}, false
}

// MatchLines return multip line
func (table *Table) MatchLines(keyName, matchValue string) ([]Line, bool) {
	if table.Keys == nil {
		return []Line{}, false
	}

	n, result := table.Keys.GetIndex(keyName)
	if !result { // 没有这个keyName
		return []Line{}, false
	}

	lines := []Line{}
	for _, v := range table.Lines {
		matched, e := regexp.MatchString(matchValue, v.Values[n])
		if e != nil {
			return []Line{}, false
		}

		if matched {
			lines = append(lines, v)
		}
	}

	return lines, true
}

// GetValuesByKey ...
func (table *Table) GetValuesByKey(key string) ([]string, bool) {
	if table.Keys == nil {
		return []string{}, false
	}

	n, result := table.Keys.GetIndex(key)
	if !result {
		return []string{}, false
	}

	values := []string{}
	for _, v := range table.Lines {
		value, result := v.GetValueByN(n)
		if result {
			values = append(values, value)
		}
	}

	return values, true
}

// GetN Get Element N
func (table *Table) GetN(n int) (Line, bool) {
	if len(table.Lines) >= (n - 1) {
		return table.Lines[n], true
	}

	return Line{}, false
}

// GetInt64ByKey 获取指定键值对的int64的值
func (table *Table) GetInt64ByKey(key, value, target string) (int64, error) {
	if table.Keys == nil {
		return int64(0), errors.New(keyNotFound + key)
	}
	n, result := table.Keys.GetIndex(key)
	if !result {
		return int64(0), errors.New(keyNotFound + key)
	}
	for _, v := range table.Lines {
		if v.Values[n] == value {
			return v.GetInt64(target)
		}
	}
	return int64(0), errors.New(keyNotFound + key)
}
