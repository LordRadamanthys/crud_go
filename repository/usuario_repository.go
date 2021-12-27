package repository

import (
	"fmt"
	"log"

	"github.com/LordRadamanthys/crud_go/config"
	_ "github.com/go-sql-driver/mysql"
)

const (
	QUERY_INSERT     = "INSERT INTO	usuarios (nome) VALUES (?)"
	QUERY_UPDATE     = "UPDATE usuarios SET nome=? WHERE id = ?"
	QUERY_SELECT_ONE = "SELECT id, nome from usuarios where id=%d"
	QUERY_DELETE     = "DELETE FROM usuarios WHERE id=%d"
)

var db, err = config.CreateConection("usuarios")

func init() {
	if err != nil {
		panic(err)
	}
}

type Usuario struct {
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
		log.Fatal(err.Error())
		return err.Error()
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
		log.Fatal(err)
		return erro.Error()
	} else if rows > 0 {
		return "Usuario apagado com exito!"
	}
	return "Não foi possivel apagar este usuario!"
}

func GetUser(id int) Usuario {
	var user Usuario
	defer db.Close()
	rows, err := db.Query(fmt.Sprintf(QUERY_SELECT_ONE, id))

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.ID, &user.Nome)
	}

	return user

}

func UpdateUser(user Usuario) Usuario {
	defer db.Close()
	_, err := db.Exec(QUERY_UPDATE, user.Nome, user.ID)
	if err != nil {
		panic(err)
	}

	updateUser := GetUser(user.ID)
	return updateUser
}
