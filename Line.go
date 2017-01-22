package GoCSVManager

// Line ...
type Line struct {
	Keys   *Key
	Values []string
}

func createLine(k *Key, s []string) Line {
	line := Line{}
	line.Keys = k
	line.Values = s
	return line
}

// GetValueBy ...
func (line *Line) GetValueBy(key string) (string, bool) {
	n, result := line.Keys.GetIndex(key)
	if !result {
		return "", false
	}

	return line.Values[n], true
}
