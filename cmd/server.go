package main

import (
	"fmt"
	"github.com/mngoutham/golang_basics/proto"
)

func main()  {

	fmt.Println("Hello World!")

	todo := &hello.Todo{Text:"Goutham"}
	fmt.Println("TODO", todo)
}
