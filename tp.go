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

func CrearTablar(db *sql.DB) {

	_, err = db.Exec(`	create table cliente (nrocliente int, nombre text, apellido text, domicilio text, telefono varhar(12))
						create table tarjeta (nrotarjeta varchar(16), nrocliente int, validadesde varchar(6), validahasta varchar(6),codseguridad varchar(4), limitecompra decimal(8,2), estado varchar(10))
						
						create table comercio (nrocomercio int, nombre text, domicilio text, codigopostal varchar(8), telefono varchar(12))
						create table compra (nrooperacion int, nrotarjeta varchar(16), nrocomercio int, fecha timestamp, monto decimal(7,2), pagado boolean)

						create table rechazo (nrorechazo int, nrotarjeta varchar(16), nrocomercio int, fecha timestamp, monto decimal(7,2), motivo text)

						create table cierre (anio int, mes int, terminacion int, fechainicio date, fechacierre date, fechavto date)
						create table cabecera (nroresumen int, nombre text, apellido text, domicilio text, nrotarjeta varchar(16), desde date, hasta date, vence date, total decimal(8,2))

						create table detalle (nroresumen int, nrolinea int, fecha date, nombrecomercio text, monto decimal(7,2))

						create table alerta (nroalerta int, nrotarjeta varchar(16), fecha timestamp, nrorechazo int, codalerta int, descripcion text)

						create table consumo (nrotarjeta varchar(16), codseguridad varchar(4), nrocomercio int, monto decimal(7,2))`)
    if err != nil {
        log.Fatal(err)
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