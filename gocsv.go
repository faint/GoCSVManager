package gocsv

var self *List

// GetInstance ...
func GetInstance() *List {
	if self == nil {
		l := new(List)
		self = l
	}
	return self
}
