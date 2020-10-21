package main

import (
	"fmt"
)

type man struct {
	name     string
	lastname string
}
type phone struct {
	number string
	model  string
}

func main() {
	fio := new(man)
	fmt.Println(fio)
}
