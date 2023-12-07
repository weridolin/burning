package common

import (
	"fmt"
	"strconv"

	"github.com/gofrs/uuid"
)

func Str2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("string to int error", err)
	}
	return i
}

func GetUUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
