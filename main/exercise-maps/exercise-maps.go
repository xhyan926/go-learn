package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {

	m := make(map[string]int)
	ar := strings.Split(s, " ")

	for _, v := range ar {
		if value, exists := m[v]; exists {
			m[v] = value + 1
		} else {
			m[v] = 1
		}
	}

	return m
}

func main() {
	fmt.Println(wordCount("hello hello hello hi hi"))
}
