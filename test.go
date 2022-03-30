package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  hello world! sfa "
	sub := strings.Split(s, " ")

	res := []string{}
	for i := len(sub) - 1; i >= 0; i-- {
		res = append(res, sub[i])
		fmt.Println(sub[i])
	}
	// fmt.Println(sub)
	// joinstr := strings.Join(res, " ")
	// sub = strings.TrimSpace(sub)
	// fmt.Println(res)
	// fmt.Println(joinstr)
}
