package main

import (
	"fmt"

	"github.com/LordRadamanthys/crud_go/repository"
)

func main() {

	usuario := repository.Usuario{
		ID:   203,
		Nome: "Jorginho traz a doze",
	}

	// result := repository.InsertUser("Roger Habbit")
	// result := repository.DeleteUser(202)
	// result := repository.GetUser(2)
	result := repository.UpdateUser(usuario)

	fmt.Println(result)

}
