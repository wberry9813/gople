package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var symbol, positive, decimal string

	if s[0] == '-' || s[0] == '+' {
		symbol = string(s[0])
	}

	if dot := bytes.IndexByte([]byte(s), '.'); dot >= 0 {
		if symbol != "" {
			positive = s[1:dot]
		} else {
			positive = s[:dot]
		}

		decimal = s[dot+1:]
	} else {
		if symbol != "" {
			positive = s[1:]
		} else {
			positive = s
		}
	}

	var buf bytes.Buffer
	buf.WriteString(symbol)
	buf.WriteString(positive[:len(positive)%3])

	for i := len(positive) % 3; i < len(positive); i += 3 {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(positive[i : i+3])
	}

	if decimal != "" {
		buf.WriteString(".")

		for i := 0; i < len(decimal)-len(decimal)%3; i += 3 {
			buf.WriteString(decimal[i : i+3])
			if i+3 < len(decimal) {
				buf.WriteString(",")
			}
		}
		buf.WriteString(decimal[len(decimal)-len(decimal)%3:])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("123654"))
	fmt.Println(comma("123654.45457"))
	fmt.Println(comma("-123654"))
	fmt.Println(comma("-123654.4847698"))
}
