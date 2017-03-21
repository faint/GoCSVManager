package gocsv

import (
	"errors"
	"strconv"
	"strings"
)

const (
	keyNotFound = "key not found in csv:"
)

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

// GetValueByN ...
func (line *Line) GetValueByN(n int) (string, bool) {
	if line.Values[n] == "" {
		return "", false
	}
	return line.Values[n], true
}

// GetString return string
func (line *Line) GetString(key string) (string, error) {
	v, result := line.GetValueBy(key)
	if !result {
		return "", errors.New(keyNotFound + key)
	}
	return v, nil
}

// GetInt64 return int64
func (line *Line) GetInt64(key string) (int64, error) {
	v, result := line.GetValueBy(key)
	if !result {
		return int64(0), errors.New(keyNotFound + key)
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return int64(0), err
	}
	return i, nil
}

// GetInt32 return int32
func (line *Line) GetInt32(key string) (int32, error) {
	v, result := line.GetValueBy(key)
	if !result {
		return int32(0), errors.New(keyNotFound + key)
	}
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return int32(0), err
	}
	return int32(i), nil
}

// GetInt return int
func (line *Line) GetInt(key string) (int, error) {
	v, result := line.GetValueBy(key)
	if !result {
		return int(0), errors.New(keyNotFound + key)
	}
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return int(0), err
	}
	return int(i), nil
}

// GetIntSlice return []int
func (line *Line) GetIntSlice(key string) ([]int, error) {
	var intSlice []int
	v, result := line.GetValueBy(key)
	if !result {
		return intSlice, errors.New(keyNotFound + key)
	}
	slice := strings.Split(v, "|")
	for _, p := range slice {
		i, err := strconv.ParseInt(p, 10, 32)
		if err != nil {
			return []int{}, err
		}
		intSlice = append(intSlice, int(i))
	}
	return intSlice, nil
}
