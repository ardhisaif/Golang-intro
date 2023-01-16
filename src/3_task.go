package src

import (
	// "fmt"
	"errors"
	"strconv"
)

func CheckFilm(time int, arr []int) (string, error) {
	var result string
	for i, val1 := range arr {
		for j, val2 := range arr {
			if i != j { //Avoid repetition or duplication
				if (val1 + val2) == time {
					result += strconv.Itoa(val1) + " dan " + strconv.Itoa(val2)
					break
				}
			}
		}
		if len(result) > 0 {
			break
		}
	}

	if len(result) == 0 {
		return "", errors.New("tidak ada yang film yang sesuai")
	}
	return result, nil
}