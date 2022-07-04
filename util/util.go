package util

import "strconv"

func Atoi(s string, def int) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return def
	}
	return i
}
