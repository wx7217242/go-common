package util

import (
	"strconv"
	"time"
)

func ConvertBinary(n int) string {
	result := ""

	for ; n > 0; n /= 2 {
		result = strconv.Itoa(n%2) + result
	}

	return result
}

func MillisecondSinceEpoch(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 比较两个时间毫秒值
func TimeDiffer(t1 time.Time, t2 time.Time) int64 {
	return MillisecondSinceEpoch(t1) - MillisecondSinceEpoch(t2)
}

// 返回平均时间的时间戳(毫秒）
func AverageTime(t1 time.Time, t2 time.Time) int64 {

	a := MillisecondSinceEpoch(t1)
	b := MillisecondSinceEpoch(t2)

	return b + (a-b)/2
}
