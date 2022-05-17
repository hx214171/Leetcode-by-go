package main

import (
	"fmt"
)

type Stu struct {
	Name string
}

func main() {
	fmt.Printf("%v\n", Stu{"tom"})
	fmt.Printf("%+v\n", Stu{"tom"})
}
