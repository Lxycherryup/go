package main

import "strings"

func buildString() string {
	var str strings.Builder
	for range 5 {
		str.WriteString("Hello, World!\n")
	}
	return str.String()
}
