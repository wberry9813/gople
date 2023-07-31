package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	buf.WriteString(s[:len(s)%3])
	for i := len(s) % 3; i < len(s); i += 3 {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("123654"))
}
