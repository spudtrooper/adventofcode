package common

import "bytes"

func Repeat(s string, times int) string {
	var buf bytes.Buffer
	for i := 0; i < times; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}
