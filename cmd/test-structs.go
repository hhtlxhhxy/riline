package main

import (
	"github.com/fatih/structs"
	"fmt"
)

type TestStruct struct {
	Name   string
	Id     int
	Age    int
	Enable bool
}

func main() {
	server := &TestStruct{
		Name:"hehaitao",
		Id:1,
		Age:25,
		Enable:true,
	}

	fmt.Println(structs.Map(server))
	fmt.Println(structs.Name(server))
	fmt.Println(structs.Values(server))
	fmt.Println(structs.Fields(server))
	fmt.Println(structs.Names(server))
	fmt.Println(structs.IsStruct(server))
	fmt.Println(structs.HasZero(server))

	s := structs.New(server)
	fmt.Println(s.Map())
	fmt.Println(s.Name())
	fmt.Println(s.Values())
	fmt.Println(s.Fields())
	fmt.Println(s.Names())
	fmt.Println(s.HasZero())

	name := s.Field("Name")
	fmt.Println(name.Value().(string))
	name.Set("haha")
	fmt.Println(name.Value().(string))
	fmt.Println(name.Kind())

}
