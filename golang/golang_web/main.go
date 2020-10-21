package main

import (
	"fmt"
)

type man struct {
	name     string
	lastname string
}

func main() {
	fio := new(man)
	fmt.Println(fio)
}
