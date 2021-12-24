package main

import (
	"fmt"

	"github.com/LordRadamanthys/crud_go/repository"
)

func main() {

	// result := repository.InsertUser("Roger Habbit")
	// result := repository.DeleteUser(202)
	result := repository.GetUser(2)

	fmt.Println(result)

}
