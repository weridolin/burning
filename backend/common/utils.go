package common

import (
	"fmt"
	"strconv"
)

func Str2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("string to int error", err)

	}
	return i
}
