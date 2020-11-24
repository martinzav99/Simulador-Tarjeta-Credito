package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db       *sql.DB
	err      error
	user     = "postgres"
	password = "1234"
	exitBool = false
)

func main() {
	defer exit()
	login(user, password)
	bienvenida()
	for {
		menu()
		if exitBool == true {
			break
		}
	}
}

func bienvenida() {
	fmt.Printf(
		`
	Bienvenido %s
	`, user)
}

func login(user string, password string) {
	//var scanner = ""
	//fmt.Printf("Bienvenido\n\n")
	//fmt.Printf("Nombre de Usuario [" + user + "]:")
	//fmt.Scanln(&scanner)
	//if scanner != "" {
	//	user = scanner
	//}
	//fmt.Printf("host [" + host + "]:")
	//fmt.Scanln(&scanner)
	//if scanner != "" {
	//	host = scanner
	//}
	fmt.Println("Connecting to postgres database...")
	db, err = sql.Open("postgres", "user="+user+" password="+password+" host=host dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected postgres!")
}

func exit() {
	fmt.Println("Closing connection...")
	db.Close()
	fmt.Println("Closed!")
}

func connectDatabase() {
	fmt.Println("Connecting to tpgossz database...")

	//https://notathoughtexperiment.me/blog/how-to-do-create-database-dbname-if-not-exists-in-postgres-in-golang/
	row := db.QueryRow(`SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = 'tpgossz');`)
	var exists bool
	err = row.Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	if exists == false {
		createDatabase()
		createTables()
		applyPKandFK()
		populateDatabase()
	} else {
		db, err = sql.Open("postgres", "user="+user+" password="+password+" host=host dbname=tpgossz sslmode=disable")
		if err != nil {
			log.Fatal(err)
			exit()
		}
		fmt.Println("Connected tpgossz!")
	}
}
func createDatabase() {
	fmt.Println("tpgossz database doesn't exists, creating database...")
	_, err = db.Exec(`CREATE DATABASE tpgossz;`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tpgossz database created succesfully!")
	connectDatabase()
}

func createTables() {
	fmt.Println("creating tables...")
	_, err = db.Exec(`	create table cliente (nrocliente int, nombre text, apellido text, domicilio text, telefono varchar(12));
						create table tarjeta (nrotarjeta varchar(16), nrocliente int, validadesde varchar(6), validahasta varchar(6),codseguridad varchar(4), limitecompra decimal(8,2), estado varchar(10));
						
						create table comercio (nrocomercio int, nombre text, domicilio text, codigopostal varchar(8), telefono varchar(12));
						create table compra (nrooperacion int, nrotarjeta varchar(16), nrocomercio int, fecha timestamp, monto decimal(7,2), pagado boolean);

						create table rechazo (nrorechazo int, nrotarjeta varchar(16), nrocomercio int, fecha timestamp, monto decimal(7,2), motivo text);

						create table cierre (anio int, mes int, terminacion int, fechainicio date, fechacierre date, fechavto date);
						create table cabecera (nroresumen int, nombre text, apellido text, domicilio text, nrotarjeta varchar(16), desde date, hasta date, vence date, total decimal(8,2));

						create table detalle (nroresumen int, nrolinea int, fecha date, nombrecomercio text, monto decimal(7,2));

						create table alerta (nroalerta int, nrotarjeta varchar(16), fecha timestamp, nrorechazo int, codalerta int, descripcion text);

						create table consumo (nrotarjeta varchar(16), codseguridad varchar(4), nrocomercio int, monto decimal(7,2));`)
	if err != nil {
		log.Fatal()
	} else {
		fmt.Println("Tables created succesfully!")
	}
}

func applyPKandFK() {

}

func populateDatabase() {
	addClients()
	addBusiness()
}

func addClients() {
	_, err = db.Exec(`	insert into cliente values (1, 'Leandro', 	'Sosa', 	'Marco Sastre 4540','541152774600');
						insert into cliente values (2, 'Leonardo', 	'Sanabria', 'Gaspar Campos 1815','541148611570');
						insert into cliente values (3, 'Florencia', 'Knol', 	'Zapiola 2825', 	'541148913800');
						insert into cliente values (4, 'Romina', 	'Subelza', 	'Libertad 3113', 	'541149422726');
						insert into cliente values (5, 'Marisa', 	'Sanchez', 	'Italia 812', 		'541143819523');
						insert into cliente values (6, 'Leonardo', 	'Sanabria', 'Gaspar Campos 1815','541143344001');
						insert into cliente values (7, 'Sebastian', 'Saavedra', 'Juncal 1139', 		'541147735133');
						insert into cliente values (8, 'Matias', 	'Palermo', 	'Godoy Cruz 2725', 	'541143344001');
						insert into cliente values (9, 'Alejandro', 'Belgrano', 'Obligado 2727', 	'541152774600');
						insert into cliente values (10, 'Florencia', 'Diotallevi', 'Ecuador 282', 	'541148341571');
						insert into cliente values (11, 'Camila', 	'Pipke', 	'Reconquista 914', 	'541148913800');
						insert into cliente values (12, 'Melisa', 	'Quevedo', 	'La Plata 4215', 	'541149422726');
						insert into cliente values (13, 'Micaela', 	'Valle', 	'Pasco 860', 		'541162722494');
						insert into cliente values (14, 'Abigail', 	'Gerez', 	'Pellegrini 2312',	'541143344057');
						insert into cliente values (15, 'Celeste', 	'Herenu', 	'Rivadavia 1592', 	'541172422755');
						insert into cliente values (16, 'Andrea', 	'Bernal', 	'Alvear 4215', 		'541143123003');
						insert into cliente values (17, 'Aldana', 	'Ramos', 	'Cevallos 261', 	'541143727636');
						insert into cliente values (18, 'Antonella', 'Herrera', 'Gascon 1241', 		'541148631232');
						insert into cliente values (19, 'Pedro', 	'Rafele', 	'Urquiza 1241', 	'541144927876');
						insert into cliente values (20, 'Lautaro', 	'Rolon', 	'Azcuenaga 1913', 	'541194127656');
						insert into cliente values (21, 'Ricardo', 	'Rueda', 	'Libertad 1252', 	'541147447171');`)

	if err != nil {
		log.Fatal(err)
	}
}

func addBusiness() {
	_, err = db.Exec(`	insert into comercio values (1, 'Farmacia Tell','Juncal 699',		'B1663',	'541157274612');
						insert into comercio values (2, 'Optica Bedini','Peron 781', 		'B1871',	'541174654172');
						insert into comercio values (3, 'Terravision',	'Urquiza 1361',	 	'B1221',	'541183910808');
						insert into comercio values (4, 'Optica Lutz', 	'Libertad 3113', 	'B1636',	'541149476322');
						insert into comercio values (5, 'Chatelet', 	'Italia 812', 		'B1663',	'541140715725');
						insert into comercio values (6, 'Magoya', 		'Peron 1601', 		'B1810',	'541153682324');
						insert into comercio values (7, 'Mayo Resto', 	'Mitre 1319', 		'B1613',	'541198035313');
						insert into comercio values (8, 'Macowens', 	'Gascon 1481', 		'B1850', 	'541143565021');
						insert into comercio values (9, 'Mundo Peluche','Balbin 1645', 		'B1613',	'541152604684');
						insert into comercio values (10, 'Sonia Novias','Sarmiento 1468', 	'C1827',	'541158573111');
						insert into comercio values (11, 'Lentes Novar','Rivadavia 5802', 	'C1002',	'541141213088');
						insert into comercio values (12, 'TatuArte', 	'Paunero 1564', 	'C1012',	'541149433826');
						insert into comercio values (13, 'Kosiuko', 	'Marco Sastre 1840','C1026',	'541180712494');
						insert into comercio values (14, 'Ossira', 		'Paunero 545', 		'C1008',	'541143314057');
						insert into comercio values (15, 'Blindado Bar','Ecuador 5451', 	'C1022',	'541105927551');
						insert into comercio values (16, 'Epic Shop', 	'Alvear 6014', 		'C1017',	'541143128703');
						insert into comercio values (17, 'XS Resto', 	'Pasco 1261', 		'C1222',	'541143027636');
						insert into comercio values (18, 'Hipervision', 'Libertad 1241', 	'C1244',	'541189151232');
						insert into comercio values (19, 'Cibernet', 	'Urquiza 1241', 	'B1224',	'541144945876');
						insert into comercio values (20, 'Crazy World', 'Zapiola 1086', 	'B1199',	'541175085786');
						insert into comercio values (21, 'Piero', 		'Tribulato 1333', 	'B1201',	'541142147877');`)

	if err != nil {
		log.Fatal(err)
	}
}

func menu() {
	menu :=
		`
	[ 1 ] Conectar con la Base
	[ 2 ] Opciones avanzadas
	[ 0 ] Salir
	
	Elige una opción
`
	fmt.Printf(menu)

	var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
	fmt.Scan(&eleccion)

	switch eleccion {
	case 1:
		connectDatabase()
	case 2:
		menuAdvanced()
	case 0:
		exitBool = true
		fmt.Println("Hasta Luego")
	default:
		fmt.Println("No elegiste ninguno")
	}
}

func menuAdvanced() {
	menuString :=
		`
		*Opciones Avanzadas*

	[ 1 ] Crear Base tpgossz
	[ 2 ] Crear tablas
	[ 3 ] Crear PK y FK
	[ 4 ] Popular Base de datos
	[ 5 ] Popular solo clientes
	[ 6 ] Popular solo negocios
	[ 0 ] Volver
	
	Elige una opción
`
	fmt.Printf(menuString)

	var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
	fmt.Scan(&eleccion)

	switch eleccion {
	case 1:
		createDatabase()
	case 2:
		createTables()
	case 3:
		applyPKandFK()
	case 4:
		populateDatabase()
	case 5:
		addClients()
	case 6:
		addBusiness()
	case 0:
		menu()
	default:
		fmt.Println("No elegiste ninguno")
		menuAdvanced()
	}
}
