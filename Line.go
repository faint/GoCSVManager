package GoCSVManager

import "strings"

// Line ...
type Line struct {
	Keys   *Key
	Values []string
}

// SetKey ...
func (line *Line) SetKey(k *Key) {
	line.Keys = k
}

// SetValue ...
func (line *Line) SetValue(s string) {
	line.Values = strings.Split(s, ",")
}

// GetValueBy ...
func (line *Line) GetValueBy(key string) (string, bool) {
	n, result := line.Keys.GetIndex(key)
	if !result {
		return "", false
	}

	return line.Values[n], true
}
