package util

import "strconv"

func ConvertBinary(n int) string {
	result := ""

	for ; n > 0; n /= 2 {
		result = strconv.Itoa(n%2) + result
	}

	return result
}
