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
	fmt.Printf(`
		Bienvenido %s!
	`, user)
}

func login(user string, password string) {
	fmt.Println("Connecting to postgres database...")
	db, err = sql.Open("postgres", "user="+user+" password="+password+" host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to postgres!")
}

func exit() {
	fmt.Println("Closing connection...")
	db.Close()
	fmt.Println("Closed!")
}

func createTables() {
	fmt.Println("Creating tables...")
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
		log.Fatal(err)
	} else {
		fmt.Println("Tables created succesfully!")
	}
}

func populateDatabase() {
	fmt.Println("Populating Database...")
	addClients()
	addBusiness()
	addTarjetas()
	generateCierres()
	fmt.Println("Database populated!")
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
						insert into cliente values (20, 'Lautaro', 	'Rolon', 	'Azcuenaga 1913', 	'541194127656');`)

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

func addTarjetas() {
	_, err = db.Exec(`	insert into tarjeta values ('5555899304583399', 1, 	'200911', '250221',	'1234', 100000.90, 'vigente');
						insert into tarjeta values ('5269399188431044', 2, 	'190918', '240928',	'0334', 50000.55, 	'vigente');
						insert into tarjeta values ('8680402479723030', 3, 	'180322', '230322',	'8214', 700000.12, 	'vigente');
						insert into tarjeta values ('7760048064179840', 4, 	'170211', '220221',	'4134', 100000.85, 	'vigente');
						insert into tarjeta values ('6317807399246634', 5, 	'200121', '250121',	'2324', 800000.22, 	'vigente');
						insert into tarjeta values ('2913395189972781', 6, 	'180819', '230828',	'4321', 900000.38, 	'vigente');
						insert into tarjeta values ('4681981280484337', 7,	'201121', '251121',	'8765', 100000.58, 	'vigente');
						insert into tarjeta values ('9387191057338602', 8, 	'160910', '210920',	'1253', 650000.85, 'vigente');
						insert into tarjeta values ('2503782418139215', 9, 	'161226', '211226',	'8367', 100000.87, 	'vigente');
						insert into tarjeta values ('4462725109757091', 10, '200901', '250921',	'6754', 20000.14, 	'vigente');
						insert into tarjeta values ('2954596377708750', 11, '180911', '230921',	'7852', 200000.50, 'vigente');
						insert into tarjeta values ('6231348143458624', 12, '161221', '211221',	'9873', 54000.25, 	'vigente');
						insert into tarjeta values ('4919235066192653', 13, '190911', '240921',	'6753', 10000.00, 	'vigente');
						insert into tarjeta values ('3742481627352427', 14, '170928', '220928',	'9801', 45000.56, 	'vigente');
						insert into tarjeta values ('2884720084187620', 15, '180111', '230121',	'9876', 500000.75, 	'vigente');
						insert into tarjeta values ('2340669528486435', 16, '170923', '220923',	'6752', 9000.80, 	'vigente');
						insert into tarjeta values ('2377527131015460', 17, '190912', '240922',	'0987', 100000.23, 	'vigente');
						insert into tarjeta values ('8472072142547842', 18, '200421', '250421',	'6987', 650000.00, 	'vigente');
						insert into tarjeta values ('3573172713553770', 19, '180216', '230226',	'0981', 220000.25, 	'vigente');
						insert into tarjeta values ('5552648744023638', 20, '170425', '220425',	'8974', 100000.45, 	'vigente');
						insert into tarjeta values ('6326855100263642', 1, 	'180607', '230627',	'9821', 450000.78, 	'suspendida');
						insert into tarjeta values ('8203564386694367', 2, 	'140728', '190728',	'0912', 9000.99, 	'anulada');`)
	if err != nil {
		log.Fatal(err)
	}
}

func addPKandFK() {
	fmt.Println("Adding PKs and FKs...")
	addPKs()
	addFKs()
	fmt.Println("PKs and FKs added succesfully!")
}

func addPKs() {
	_, err = db.Exec(`	alter table cliente add constraint cliente_pk primary key (nrocliente);
						alter table tarjeta add constraint tarjeta_pk primary key (nrotarjeta);
						alter table comercio add constraint comercio_pk primary key (nrocomercio);
						alter table compra add constraint compra_pk primary key (nrooperacion);
						alter table rechazo add constraint rechazo_pk primary key (nrorechazo);
						alter table cierre add constraint cierre_pk primary key (anio, mes, terminacion);
						alter table cabecera add constraint cabecera_pk primary key (nroresumen);
						alter table detalle add constraint detalle_pk primary key (nroresumen, nrolinea);
						alter table alerta add constraint alerta_pk primary key (nroalerta);`)
	if err != nil {
		log.Fatal(err)
	}
}

func addFKs() {
	_, err = db.Exec(`	alter table tarjeta add constraint tarjeta_nrocliente_fk foreign key (nrocliente) references cliente (nrocliente);
						--alter table rechazo add constraint rechazo_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
						alter table compra add constraint compra_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
						alter table alerta add constraint alerta_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
						alter table cabecera add constraint cabecera_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
						alter table alerta add constraint alerta_nrorechazo_fk foreign key (nrorechazo) references rechazo (nrorechazo);
						alter table rechazo add constraint rechazo_nrocomercio_fk foreign key (nrocomercio) references comercio (nrocomercio);
						alter table compra add constraint compra_nrocomercio_fk foreign key (nrocomercio) references comercio (nrocomercio);`)
	if err != nil {
		log.Fatal(err)
	}
}

func dropPKandFK() {
	fmt.Println("Removing PKs and FKs...")
	dropFKs()
	dropPKs()
	fmt.Println("PKs and FKs removed succesfully!")
}

func dropPKs() {
	_, err = db.Exec(`	alter table cliente drop  constraint cliente_pk;
						alter table tarjeta drop constraint tarjeta_pk;
						alter table comercio drop constraint comercio_pk;
						alter table compra drop constraint compra_pk;
						alter table rechazo drop constraint rechazo_pk;
						alter table cierre drop constraint cierre_pk;
						alter table cabecera drop constraint cabecera_pk;
						alter table detalle drop constraint detalle_pk;
						alter table alerta drop constraint alerta_pk;`)
	if err != nil {
		log.Fatal(err)
	}
}

func dropFKs() {
	_, err = db.Exec(`	alter table tarjeta drop constraint tarjeta_nrocliente_fk;
						--alter table rechazo drop constraint rechazo_nrotarjeta_fk;
						alter table compra drop constraint compra_nrotarjeta_fk;
						alter table alerta drop constraint alerta_nrotarjeta_fk;
						alter table cabecera drop constraint cabecera_nrotarjeta_fk;
						alter table alerta drop constraint alerta_nrorechazo_fk;
						alter table rechazo drop constraint rechazo_nrocomercio_fk;
						alter table compra drop constraint compra_nrocomercio_fk;`)
	if err != nil {
		log.Fatal(err)
	}
}

func generateCierres() {
	for nMes := 1; nMes <= 12; nMes++ {
		for terminacion := 1; terminacion <= 9; terminacion++ {
			if nMes == 12 {
				_, err = db.Exec(fmt.Sprintf("insert into cierre values (2020, %v, %v, '2020%v0%v', '2020010%v', '202001%v');", nMes, terminacion, nMes, terminacion, terminacion, terminacion+9))
				if err != nil {
					log.Fatal(err)
				}
			} else if nMes > 9 {
				_, err = db.Exec(fmt.Sprintf("insert into cierre values (2020, %v, %v, '2020%v0%v', '2020010%v', '202001%v');", nMes, terminacion, nMes, terminacion, terminacion, terminacion+9))
				if err != nil {
					log.Fatal(err)
				}
			} else {
				_, err = db.Exec(fmt.Sprintf("insert into cierre values (2020, %v, %v, '20200%v0%v', '20200%v0%v', '20200%v%v');", nMes, terminacion, nMes, terminacion, nMes+1, terminacion, nMes+1, terminacion+9))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		if nMes == 12 {
			_, err = db.Exec(fmt.Sprintf("insert into cierre values (2020, %v, 0, '2020%v10', '20200110', '20200119');", nMes, nMes))
			if err != nil {
				log.Fatal(err)
			}
		} else if nMes > 9 {
			_, err = db.Exec(fmt.Sprintf("insert into cierre values (2020, %v, 0, '2020%v10', '2020%v10', '2020%v19');", nMes, nMes, nMes+1, nMes+1))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_, err = db.Exec(fmt.Sprintf("insert into cierre values (2020, %v, 0, '20200%v10', '20200%v10', '20200%v19');", nMes, nMes, nMes+1, nMes+1))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func addStoredProceduresTriggers() {
	fmt.Println("Adding Stored Procedures and Triggers...")
	addAutorizacionDeCompra()
	//addOtroTrigger()
	fmt.Println("Done adding Stored Procedures and Triggers!")
}

func addAutorizacionDeCompra() {
	fmt.Println(" Adding 'Autorizacion De Compra' Procedure")
	_, err = db.Exec(`	create or replace function a_de_compra(nrotarjetax char , codseguridadx char , nrocomerciox int , montox decimal) returns boolean as $$
						declare
							montoCompraSum int;
							tarjetaRecord record;
							fechaActual date;
							timeActual timestamp;
							nrechazo int;
							noperacion int;
							montoTotal int;
						
						begin
							select count (nrooperacion) into noperacion from compra;
							select count(nrorechazo) into nrechazo from rechazo;
							select current_date into fechaActual;
						
							select * from tarjeta into tarjetaRecord where nrotarjeta = nrotarjetax;
						
							if not found then
								select current_timestamp into timeActual;
								insert into rechazo values (nrechazo,nrotarjetax,nrocomerciox,timeActual,montox,'tarjeta no valida o no vigente');
								return false;
							elsif tarjetaRecord.codseguridad != codseguridadx then
								select current_timestamp into timeActual;
								insert into rechazo values (nrechazo,nrotarjetax,nrocomerciox,timeActual,montox,'codigo de seguridad invalido');
								return false;
							elsif CAST(tarjetaRecord.validahasta as date) < fechaActual then /* arreglar */
								select current_timestamp into timeActual;
								insert into rechazo values (nrechazo,nrotarjetax,nrocomerciox,timeActual,montox,'plazo de vigencia expirado');
								return false;
							elsif tarjetaRecord.estado = 'suspendida' then
								select current_timestamp into timeActual;
								insert into rechazo values (nrechazo,nrotarjetax,nrocomerciox,timeActual,montox,'la tarjeta se encuentra suspendida');
								return false;
							end if;
							
							select sum(monto) into montoCompraSum from compra where nrotarjeta=nrotarjetax and pagado = false;
							montoTotal := montoCompraSum + montox;
						
							if tarjetaRecord.limitecompra < montoCompraSum then
								select current_timestamp into timeActual;
								insert into rechazo values (nrechazo,nrotarjetax,nrocomerciox,timeActual,montox,'supera limite de tarjeta');
								return false;
							end if;
							
							select current_timestamp into timeActual;
							insert into compra values (noperacion,nrotarjetax,nrocomerciox,timeActual,montox,false);
							return true;
						
						end;
						$$language plpgsql;`)
	if err != nil {
		log.Fatal(err)
	}
}

func menu() {
	menuString :=
		`
			Menu principal
		[ 1 ] Crear Base tpgossz (Auto)
		[ 2 ] Crear Base tpgossz (Manual)

		[ 0 ] Salir
		
		Elige una opción
		`
	fmt.Printf(menuString)

	var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
	fmt.Scan(&eleccion)

	switch eleccion {
	case 1:
		autoCreateDatabase()
	case 2:
		menuCreacionMnual()
	case 0:
		exitBool = true
		fmt.Println("Hasta Luego")
	default:
		fmt.Println("No elegiste ninguno")
	}
}
func menuCreacionMnual() {
	menuString :=
		`
			Menu de creacion Manual
		[ 1 ] Eliminar Base tpgossz
		[ 2 ] Crear Base tpgossz
		[ 3 ] Conectar con Base tpgossz
		[ 4 ] Crear tablas
		[ 5 ] Agregar PKs y FKs
		[ 6 ] Popular Base de datos
		[ 7 ] Remover PKs y FKs
		[ 8 ] Agregar Stored Procedures y Triggers

		[ 0 ] Volver
		
		Elige una opción
		`
	fmt.Printf(menuString)

	var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
	fmt.Scan(&eleccion)

	switch eleccion {
	case 1:
		dropDatabase()
	case 2:
		createDatabase()
	case 3:
		connectDatabase()
	case 4:
		createTables()
	case 5:
		addPKandFK()
	case 6:
		populateDatabase()
	case 7:
		dropPKandFK()
	case 8:
		addStoredProceduresTriggers()
	case 0:
		menu()
	default:
		fmt.Println("No elegiste ninguno")
	}
}

func autoCreateDatabase() {
	dropDatabase()
	createDatabase()
	connectDatabase()
	createTables()
	addPKandFK()
	populateDatabase()
	addStoredProceduresTriggers()
	fmt.Println("\nReady to work!")
}

func dropDatabase() {
	fmt.Println("Dropping tpgossz database if exists...")
	checkIfUsersConnected()
	_, err = db.Exec(`drop database if exists tpgossz;`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tpgossz database dropped!")
}

func createDatabase() {
	fmt.Println("Creating tpgossz Database...")
	_, err = db.Exec(`CREATE DATABASE tpgossz;`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tpgossz database created succesfully!")
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
		fmt.Println("tpgossz database doesn't exist!")
		createDatabase()

	} else {
		db, err = sql.Open("postgres", "user="+user+" password="+password+" host=localhost dbname=tpgossz sslmode=disable")
		if err != nil {
			log.Fatal(err)
			exit()
		}
		fmt.Println("Connected tpgossz!")
	}
}

func checkIfUsersConnected() {
	fmt.Println(" Checking if there are users connected berfore dropping...")
	var count int
	row := db.QueryRow(`SELECT count(*) FROM pg_stat_activity WHERE datname = 'tpgossz';`)
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		concatenated := fmt.Sprintf("  Found %d users connected", count)
		fmt.Println(concatenated)
		disconnectUsers()
	} else {
		fmt.Println(" No users connected")
	}

}

func disconnectUsers() {
	connectPostgres()
	fmt.Println("   Disconnecting users...")
	_, err = db.Exec(`REVOKE CONNECT ON DATABASE tpgossz FROM public;`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`SELECT pg_terminate_backend(pg_stat_activity.pid)
					  FROM pg_stat_activity
				      WHERE pg_stat_activity.datname = 'tpgossz';`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("   Disconnected users succesfully!")
}

func connectPostgres() {
	fmt.Println("   Connecting to postgres database before disconnecting tpgossz users")
	db, err = sql.Open("postgres", "user="+user+" password="+password+" host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("   Connected to postgres!")
}
