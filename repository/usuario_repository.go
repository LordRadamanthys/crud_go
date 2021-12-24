package repository

import (
	"fmt"
	"log"

	"github.com/LordRadamanthys/crud_go/config"
	_ "github.com/go-sql-driver/mysql"
)

const (
	QUERY_INSERT     = "INSERT INTO	usuarios (nome) VALUES (?)"
	QUERY_UPDATE     = "UPDATE usuarios SET nome=%s WHERE id = %d"
	QUERY_SELECT_ONE = "SELECT id, nome from usuarios where id=%d"
	QUERY_DELETE     = "DELETE FROM usuarios WHERE id=%d"
)

var db, err = config.CreateConection("usuarios")

func init() {
	if err != nil {
		panic(err)
	}
}

type usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func InsertUser(nome string) string {

	defer db.Close()
	result, err := db.Exec(QUERY_INSERT, nome)
	if err != nil {
		panic(err)
	}

	rows, erro := result.RowsAffected()

	if erro != nil {
		return err.Error()
		log.Fatal(err.Error())
	} else if rows > 0 {
		return "Inserido com sucesso"
	}
	return "erro ao inserir"
}

func DeleteUser(id int) string {

	defer db.Close()
	result, err := db.Exec(fmt.Sprintf(QUERY_DELETE, id))
	if err != nil {
		panic(err)
	}

	rows, erro := result.RowsAffected()

	if erro != nil {
		return erro.Error()
		log.Fatal(err.Error())
	} else if rows > 0 {
		return "Usuario apagado com exito!"
	}
	return "Não foi possivel apagar este usuario!"
}

func GetUser(id int) usuario {
	var user usuario
	defer db.Close()
	rows, err := db.Query(fmt.Sprintf(QUERY_SELECT_ONE, id))
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Nome)
	}

	return user

}
