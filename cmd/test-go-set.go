package main

import (
	"gopkg.in/fatih/set.v0"
	"fmt"
)

func main() {
	s := set.New()

	s.Add("hehaitao")
	size := s.Size()
	fmt.Println(size)
}

