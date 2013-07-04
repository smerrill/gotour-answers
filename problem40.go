package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	x := make(map[string]int)
	for _, v := range strings.Fields(s) {
		x[v]++
	}
	return x
}

func main() {
	wc.Test(WordCount)
}
