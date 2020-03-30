package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Name      string
	Usuario   string
	Contra    string
	Url       string
	Port      string
	BaseDatos string
}
type Userr struct {
	id         string `json:"id"`
	nombre     string `json:"nombre"`
	usuario    string `json:"usuario`
	contrasena string `json:"contrasena"`
}

func NewBD(name string, user string, contra string, url string, port string, baseDatos string) *Mysql {
	return &Mysql{
		Name:      name,
		Usuario:   user,
		Contra:    contra,
		Url:       url,
		Port:      port,
		BaseDatos: baseDatos,
	}
}

func (s *Mysql) NewConeccion() (*sql.DB, error) {
	db, err := sql.Open(s.Name, s.Usuario+":"+s.Contra+"@tcp("+s.Url+":"+s.Port+")/"+s.BaseDatos)
	if err != nil {
		return nil, err
	}
	fmt.Println("exito ")
	return db, nil
}

func obtenerUsuarios() ([]User, error) {
	contactos := []User{}
	bd := NewBD("mysql", "VtgHpFxzCP", "V7sFv16RgJ", "remotemysql.com", "3306", "VtgHpFxzCP")

	db, err := bd.NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT * FROM VtgHpFxzCP.Administrador")

	if err != nil {
		fmt.Println("error en la consulta")
		return nil, err
	}
	defer filas.Close()

	var c User

	for filas.Next() {

		err = filas.Scan(&c.id, &c.nombre, &c.usuario, &c.contrasena)
		if err != nil {
			fmt.Println("error al scanear")
			return nil, err
		}

		contactos = append(contactos, c)
	}
	return contactos, nil
}
