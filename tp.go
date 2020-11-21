package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	db  *sql.DB
	err error
)

func main() {
	fmt.Println("Conectado con BD")
	db, err := sql.Open("postgres", "user=postgres password=1234 host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a BD")
	defer fmt.Println("Cerrrando conexion")
	defer db.Close()
	
	Bienvenida(db)
}

func CreateDatabase(db *sql.DB) {

	_, err = db.Exec(`drop database if exists tpgossz;`)
		
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("drop")
	
	_, err = db.Exec(`create database tpgossz;`)
		fmt.Println("create")
	if err != nil {
		log.Fatal(err)
	}
}

func Bienvenida(db *sql.DB) {
	menu :=
		`
	Bienvenido
	[ 1 ] Pizza
	[ 2 ] Tacos
	¿Qué prefieres?
`
	fmt.Print(menu)

	var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
	fmt.Scanln(&eleccion)

	switch eleccion {
	case 1:
		CreateDatabase(db)
		fmt.Println("Prefieres pizza")
	case 2:
		fmt.Println("Prefieres tacos")
	default:
		fmt.Println("No prefieres ninguno de ellos")
	}
}