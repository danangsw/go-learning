package main

import (
	"golang.org/x/tour/wc";
	"strings"
)

func WordCount(s string) map[string]int {
	r := make(map[string]int)
	
	strs := strings.Fields(s)
	
	for _,v:=range(strs) {
		elem, ok := r[v]
		if ok {
			r[v] = 1 + elem
		} else {
			r[v] = 1
		}
	}
	
	return r
}

func main() {
	wc.Test(WordCount)
}
