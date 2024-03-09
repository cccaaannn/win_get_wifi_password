package util

import (
	"fmt"
	"strconv"
)

func GetUserInputAsInt() (int, error) {
	var selectedIndexInput string
	fmt.Printf(":%s", selectedIndexInput)
	fmt.Scanln(&selectedIndexInput)
	return strconv.Atoi(selectedIndexInput)
}
