package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"encoding/json"
	"strconv"
	bolt "github.com/coreos/bbolt"
	//bolt "go.etcd.io/bbolt"
)

var (
	db               *sql.DB
	err              error
	user             = "postgres"
	password         = "1234"
	exitBool         = false
	advancedMenuBool = false
)

func main() {
	defer exit()
	login(user, password)
	bienvenida()
	for {
		if advancedMenuBool {
			advancedMenu()
		} else {
			menu()
		}
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
	_, err = db.Exec(`	CREATE TABLE cliente (nrocliente int, nombre text, apellido text, domicilio text, telefono varchar(12));
						CREATE TABLE tarjeta (nrotarjeta varchar(16), nrocliente int, validadesde varchar(6), validahasta varchar(6),codseguridad varchar(4), limitecompra decimal(8,2), estado varchar(10));						
						CREATE TABLE comercio (nrocomercio int, nombre text, domicilio text, codigopostal varchar(8), telefono varchar(12));
						CREATE TABLE compra (nrooperacion int, nrotarjeta varchar(16), nrocomercio int, fecha timestamp, monto decimal(7,2), pagado boolean);
						CREATE TABLE rechazo (nrorechazo int, nrotarjeta varchar(16), nrocomercio int, fecha timestamp, monto decimal(7,2), motivo text, codmotivo int);
						CREATE TABLE cierre (anio int, mes int, terminacion int, fechainicio date, fechacierre date, fechavto date);
						CREATE TABLE cabecera (nroresumen int, nombre text, apellido text, domicilio text, nrotarjeta varchar(16), desde date, hasta date, vence date, total decimal(8,2));
						CREATE TABLE detalle (nroresumen int, nrolinea int, fecha date, nombrecomercio text, monto decimal(7,2));
						CREATE TABLE alerta (nroalerta int, nrotarjeta varchar(16), fecha timestamp, nrorechazo int, codalerta int, descripcion text);
						CREATE TABLE consumo (nrotarjeta varchar(16), codseguridad varchar(4), nrocomercio int, monto decimal(7,2));`)
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
	//addConsumos()
	fmt.Println("Database populated!")
}

func addClients() {
	_, err = db.Exec(`	INSERT INTO cliente VALUES (1, 'Leandro', 	'Sosa', 	'Marco Sastre 4540',	'541152774600');
						INSERT INTO cliente VALUES (2, 'Leonardo', 	'Sanabria', 'Gaspar Campos 1815',	'541148611570');
						INSERT INTO cliente VALUES (3, 'Florencia', 'Knol', 	'Zapiola 2825', 		'541148913800');
						INSERT INTO cliente VALUES (4, 'Romina', 	'Subelza', 	'Libertad 3113', 		'541149422726');
						INSERT INTO cliente VALUES (5, 'Marisa', 	'Sanchez', 	'Italia 812', 			'541143819523');
						INSERT INTO cliente VALUES (6, 'Leonardo', 	'Sanabria', 'Gaspar Campos 1815',	'541143344001');
						INSERT INTO cliente VALUES (7, 'Sebastian', 'Saavedra', 'Juncal 1139', 			'541147735133');
						INSERT INTO cliente VALUES (8, 'Matias', 	'Palermo', 	'Godoy Cruz 2725', 		'541143344001');
						INSERT INTO cliente VALUES (9, 'Alejandro', 'Belgrano', 'Obligado 2727', 		'541152774600');
						INSERT INTO cliente VALUES (10, 'Florencia', 'Diotallevi', 'Ecuador 282', 		'541148341571');
						INSERT INTO cliente VALUES (11, 'Camila', 	'Pipke', 	'Reconquista 914', 		'541148913800');
						INSERT INTO cliente VALUES (12, 'Melisa', 	'Quevedo', 	'La Plata 4215', 		'541149422726');
						INSERT INTO cliente VALUES (13, 'Micaela', 	'Valle', 	'Pasco 860', 			'541162722494');
						INSERT INTO cliente VALUES (14, 'Abigail', 	'Gerez', 	'Pellegrini 2312',		'541143344057');
						INSERT INTO cliente VALUES (15, 'Celeste', 	'Herenu', 	'Rivadavia 1592', 		'541172422755');
						INSERT INTO cliente VALUES (16, 'Andrea', 	'Bernal', 	'Alvear 4215', 			'541143123003');
						INSERT INTO cliente VALUES (17, 'Aldana', 	'Ramos', 	'Cevallos 261', 		'541143727636');
						INSERT INTO cliente VALUES (18, 'Antonella', 'Herrera', 'Gascon 1241', 			'541148631232');
						INSERT INTO cliente VALUES (19, 'Pedro', 	'Rafele', 	'Urquiza 1241', 		'541144927876');
						INSERT INTO cliente VALUES (20, 'Lautaro', 	'Rolon', 	'Azcuenaga 1913', 		'541194127656');`)
	if err != nil {
		log.Fatal(err)
	}
}

func addBusiness() {
	_, err = db.Exec(`	INSERT INTO comercio VALUES (1, 'Farmacia Tell','Juncal 699',		'B1663',	'541157274612');
						INSERT INTO comercio VALUES (2, 'Optica Bedini','Peron 781', 		'B1871',	'541174654172');
						INSERT INTO comercio VALUES (3, 'Terravision',	'Urquiza 1361',	 	'B1221',	'541183910808');
						INSERT INTO comercio VALUES (4, 'Optica Lutz', 	'Libertad 3113', 	'B1636',	'541149476322');
						INSERT INTO comercio VALUES (5, 'Chatelet', 	'Italia 812', 		'B1663',	'541140715725');
						INSERT INTO comercio VALUES (6, 'Magoya', 		'Peron 1601', 		'B1221',	'541153682324');
						INSERT INTO comercio VALUES (7, 'Mayo Resto', 	'Mitre 1319', 		'B1613',	'541198035313');
						INSERT INTO comercio VALUES (8, 'Macowens', 	'Gascon 1481', 		'B1850', 	'541143565021');
						INSERT INTO comercio VALUES (9, 'Mundo Peluche','Balbin 1645', 		'B1613',	'541152604684');
						INSERT INTO comercio VALUES (10, 'Sonia Novias','Sarmiento 1468', 	'C1827',	'541158573111');
						INSERT INTO comercio VALUES (11, 'Lentes Novar','Rivadavia 5802', 	'B1221',	'541141213088');
						INSERT INTO comercio VALUES (12, 'TatuArte', 	'Paunero 1564', 	'C1012',	'541149433826');
						INSERT INTO comercio VALUES (13, 'Kosiuko', 	'Marco Sastre 1840','C1026',	'541180712494');
						INSERT INTO comercio VALUES (14, 'Ossira', 		'Paunero 545', 		'C1008',	'541143314057');
						INSERT INTO comercio VALUES (15, 'Blindado Bar','Ecuador 5451', 	'B1221',	'541105927551');
						INSERT INTO comercio VALUES (16, 'Epic Shop', 	'Alvear 6014', 		'C1017',	'541143128703');
						INSERT INTO comercio VALUES (17, 'XS Resto', 	'Pasco 1261', 		'C1222',	'541143027636');
						INSERT INTO comercio VALUES (18, 'Hipervision', 'Libertad 1241', 	'B1221',	'541189151232');
						INSERT INTO comercio VALUES (19, 'Cibernet', 	'Urquiza 1241', 	'B1224',	'541144945876');
						INSERT INTO comercio VALUES (20, 'Crazy World', 'Zapiola 1086', 	'B1199',	'541175085786');
						INSERT INTO comercio VALUES (21, 'Piero', 		'Tribulato 1333', 	'B1201',	'541142147877');`)
	if err != nil {
		log.Fatal(err)
	}
}

func addTarjetas() {
	_, err = db.Exec(`	INSERT INTO tarjeta VALUES ('5555899304583399', 1, 	'200911', '250221',	'1234', 100000.90, 'vigente');
						INSERT INTO tarjeta VALUES ('5269399188431044', 2, 	'190918', '240928',	'0334', 50000, 	'vigente');
						INSERT INTO tarjeta VALUES ('8680402479723030', 3, 	'180322', '230322',	'8214', 700000.12, 	'vigente');
						INSERT INTO tarjeta VALUES ('7760048064179840', 4, 	'170211', '220221',	'4134', 100000.85, 	'vigente');
						INSERT INTO tarjeta VALUES ('6317807399246634', 5, 	'200121', '250121',	'2324', 800000.22, 	'vigente');
						INSERT INTO tarjeta VALUES ('2913395189972781', 6, 	'180819', '230828',	'4321', 900000.38, 	'vigente');
						INSERT INTO tarjeta VALUES ('4681981280484337', 7,	'201121', '251121',	'8765', 100000.58, 	'vigente');
						INSERT INTO tarjeta VALUES ('9387191057338602', 8, 	'160910', '210920',	'1253', 650000.85, 'vigente');
						INSERT INTO tarjeta VALUES ('2503782418139215', 9, 	'161226', '211226',	'8367', 100000.87, 	'vigente');
						INSERT INTO tarjeta VALUES ('4462725109757091', 10, '200901', '250921',	'6754', 20000.14, 	'vigente');
						INSERT INTO tarjeta VALUES ('2954596377708750', 11, '180911', '230921',	'7852', 200000.50, 'vigente');
						INSERT INTO tarjeta VALUES ('6231348143458624', 12, '161221', '211221',	'9873', 54000.25, 	'vigente');
						INSERT INTO tarjeta VALUES ('4919235066192653', 13, '190911', '240921',	'6753', 10000.00, 	'vigente');
						INSERT INTO tarjeta VALUES ('3742481627352427', 14, '170928', '220928',	'9801', 45000.56, 	'vigente');
						INSERT INTO tarjeta VALUES ('2884720084187620', 15, '180111', '230121',	'9876', 500000.75, 	'vigente');
						INSERT INTO tarjeta VALUES ('2340669528486435', 16, '170923', '220923',	'6752', 9000.80, 	'vigente');
						INSERT INTO tarjeta VALUES ('2377527131015460', 17, '190912', '240922',	'0987', 100000.23, 	'vigente');
						INSERT INTO tarjeta VALUES ('8472072142547842', 18, '200421', '250421',	'6987', 650000.00, 	'vigente');
						INSERT INTO tarjeta VALUES ('3573172713553770', 19, '180216', '230226',	'0981', 220000.25, 	'vigente');
						INSERT INTO tarjeta VALUES ('5552648744023638', 20, '170425', '220425',	'8974', 100000.45, 	'vigente');
						INSERT INTO tarjeta VALUES ('6326855100263642', 1, 	'180607', '230627',	'9821', 450000.78, 	'suspendida');
						INSERT INTO tarjeta VALUES ('8203564386694367', 2, 	'140728', '190728',	'0912', 9000.99, 	'anulada');`)
	if err != nil {
		log.Fatal(err)
	}
}

func addConsumos() {
	_, err = db.Exec(`  INSERT INTO consumo VALUES ('8680402479723030', '1'    , 10 , 600); --codigo de seguridad invalido
						INSERT INTO consumo VALUES ('8680402479723055', '8214' , 10 , 600); --tarjeta no valida o no vigente
						INSERT INTO consumo VALUES ('6326855100263642', '9821' , 10 , 600); --tarjeta suspendida
						INSERT INTO consumo VALUES ('8203564386694367', '0912' , 10 , 600); --tarjeta plazo de vigencia expirado
						INSERT INTO consumo VALUES ('5269399188431044', '0334' , 10 , 50001); --supera el limite de tarjeta
						INSERT INTO consumo VALUES ('8680402479723030', '8214' , 3  , 600); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('8680402479723030', '8214' , 11 , 600); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('8680402479723030', '8214' , 15 , 600); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('8680402479723030', '8214' , 16 , 600); --compra realizada correctamente cp C1017
						INSERT INTO consumo VALUES ('8680402479723030', '8214' , 10 , 600); --compra realizada correctamente cp C1827
						INSERT INTO consumo VALUES ('8680402479723030', '8214' , 15 , 600); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('7760048064179840', '4134' , 12 , 2000); --compra realizada correctamente cp C1012
						INSERT INTO consumo VALUES ('7760048064179840', '1111' , 2  , 5000); --codigo de seguridad invalido
						INSERT INTO consumo VALUES ('7760048064179840', '1111' , 4  , 66000.90); --supera el limite de tarjeta
						INSERT INTO consumo VALUES ('2913395189972781', '4321' , 13 , 20560.00); --compra realizada correctamente cp C1026
						INSERT INTO consumo VALUES ('4681981280484337', '8765' , 14 , 15000.50); --compra realizada correctamente cp C1008
						INSERT INTO consumo VALUES ('9387191057338602', '1253' , 15 , 600.00); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('2503782418139215', '8367' , 16 , 6500.45); --compra realizada correctamente cp C1017
						INSERT INTO consumo VALUES ('4462725109757091', '6754' , 17 , 8001.45); --compra realizada correctamente cp C1222
						INSERT INTO consumo VALUES ('2954596377708750', '7852' , 18 , 12000.70); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('6231348143458624', '9873' , 19 , 900.55); --compra realizada correctamente cp B1224
						INSERT INTO consumo VALUES ('4919235066192653', '6753' , 20 , 7000.90); --compra realizada correctamente cp B1199
						INSERT INTO consumo VALUES ('3742481627352427', '9801' , 1  , 700.95); --compra realizada correctamente cp B1663
						INSERT INTO consumo VALUES ('2884720084187620', '9876' , 2  , 1300.70); --compra realizada correctamente cp B1871
						INSERT INTO consumo VALUES ('2340669528486435', '6752' , 3  , 66600.20); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('2377527131015460', '0987' , 4  , 9000.00); --compra realizada correctamente cp B1636
						INSERT INTO consumo VALUES ('8472072142547842', '6987' , 5  , 7240.70); --compra realizada correctamente cp B1663
						INSERT INTO consumo VALUES ('3573172713553770', '0981' , 6  , 700.95); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('5552648744023638', '8974' , 7  , 3100.70); --compra realizada correctamente cp B1613
						INSERT INTO consumo VALUES ('6326855100263642', '9821' , 8  , 50200.40); --tarjeta suspendida
						INSERT INTO consumo VALUES ('8203564386694367', '0912' , 9  , 16500.00); --tarjeta anulada
						INSERT INTO consumo VALUES ('5555899304583399', '6987' , 11 , 18500.80); --compra realizada correctamente cp C1827
						INSERT INTO consumo VALUES ('5555899304583399', '6987' , 12 , 26000.00); --supera el limite de tarjeta
						INSERT INTO consumo VALUES ('5555899304583399', '6987' , 13 , 2540.90); --compra realizada correctamente cp C1026
						INSERT INTO consumo VALUES ('5269399188431044', '0334' , 14 , 5600.50); --compra realizada correctamente cp C1008
						INSERT INTO consumo VALUES ('7760048064179840', '4134' , 15 , 8000.00); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('6317807399246634', '2324' , 16 , 5000.40); --compra realizada correctamente cp C1017
						INSERT INTO consumo VALUES ('2913395189972781', '4321' , 17 , 50500.20); --compra realizada correctamente cp C1222
						INSERT INTO consumo VALUES ('4681981280484337', '8765' , 18 , 5440.10); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('9387191057338602', '1253' , 19 , 5000.40); --compra realizada correctamente cp B1224
						INSERT INTO consumo VALUES ('2503782418139215', '8367' , 20 , 5000.20); --compra realizada correctamente cp B1199
						INSERT INTO consumo VALUES ('4462725109757091', '6754' , 21 , 5440.10); --compra realizada correctamente cp B1201
						INSERT INTO consumo VALUES ('2954596377708750', '7852' , 1  , 2000.20); --compra realizada correctamente cp B1663
						INSERT INTO consumo VALUES ('6231348143458624', '9873' , 2  , 7440.10); --compra realizada correctamente cp B1871
						INSERT INTO consumo VALUES ('4919235066192653', '6753' , 3  , 2000.40); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('3742481627352427', '9801' , 4  , 50.50); --compra realizada correctamente cp B1636
						INSERT INTO consumo VALUES ('2884720084187620', '9876' , 5  , 440.80); --compra realizada correctamente cp B1663
						INSERT INTO consumo VALUES ('2340669528486435', '6752' , 6  , 4000.20); --compra realizada correctamente cp B1221
						INSERT INTO consumo VALUES ('2377527131015460', '0987' , 7  , 880.16); --compra realizada correctamente cp B1613
						INSERT INTO consumo VALUES ('8472072142547842', '6987' , 8  , 7000.40); --compra realizada correctamente cp B1850
						INSERT INTO consumo VALUES ('3573172713553770', '0981' , 9  , 950.60); --compra realizada correctamente cp B1613
						INSERT INTO consumo VALUES ('5552648744023638', '8974' , 10 , 1990.00); --compra realizada correctamente cp C1827
						INSERT INTO consumo VALUES ('6326855100263642', '9821' , 11 , 400.40); --tarjeta suspendida
						INSERT INTO consumo VALUES ('8203564386694367', '0912' , 12 , 800.16);`)
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
	_, err = db.Exec(`	ALTER TABLE cliente ADD CONSTRAINT cliente_pk PRIMARY KEY (nrocliente);
						ALTER TABLE tarjeta ADD CONSTRAINT tarjeta_pk PRIMARY KEY (nrotarjeta);
						ALTER TABLE comercio ADD CONSTRAINT comercio_pk PRIMARY KEY (nrocomercio);
						ALTER TABLE compra ADD CONSTRAINT compra_pk PRIMARY KEY (nrooperacion);
						ALTER TABLE rechazo ADD CONSTRAINT rechazo_pk PRIMARY KEY (nrorechazo);
						ALTER TABLE cierre ADD CONSTRAINT cierre_pk PRIMARY KEY (anio, mes, terminacion);
						ALTER TABLE cabecera ADD CONSTRAINT cabecera_pk PRIMARY KEY (nroresumen);
						ALTER TABLE detalle ADD CONSTRAINT detalle_pk PRIMARY KEY (nroresumen, nrolinea);
						ALTER TABLE alerta ADD CONSTRAINT alerta_pk PRIMARY KEY (nroalerta);`)
	if err != nil {
		log.Fatal(err)
	}
}

func addFKs() {
	_, err = db.Exec(`	ALTER TABLE 	tarjeta ADD CONSTRAINT tarjeta_nrocliente_fk 	FOREIGN KEY (nrocliente) REFERENCES cliente (nrocliente);
						--ALTER TABLE 	rechazo ADD CONSTRAINT rechazo_nrotarjeta_fk 	FOREIGN KEY (nrotarjeta) REFERENCES tarjeta (nrotarjeta);
						ALTER TABLE 	compra ADD CONSTRAINT compra_nrotarjeta_fk 		FOREIGN KEY (nrotarjeta) REFERENCES tarjeta (nrotarjeta);
						--ALTER TABLE 	alerta ADD CONSTRAINT alerta_nrotarjeta_fk 		FOREIGN KEY (nrotarjeta) REFERENCES tarjeta (nrotarjeta);
						ALTER TABLE 	cabecera ADD CONSTRAINT cabecera_nrotarjeta_fk 	FOREIGN KEY (nrotarjeta) REFERENCES tarjeta (nrotarjeta);
						--ALTER TABLE 	alerta ADD CONSTRAINT alerta_nrorechazo_fk 		FOREIGN KEY (nrorechazo) REFERENCES rechazo (nrorechazo);
						ALTER TABLE 	rechazo ADD CONSTRAINT rechazo_nrocomercio_fk 	FOREIGN KEY (nrocomercio) REFERENCES comercio (nrocomercio);
						ALTER TABLE 	compra ADD CONSTRAINT compra_nrocomercio_fk 	FOREIGN KEY (nrocomercio) REFERENCES comercio (nrocomercio);`)
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
	_, err = db.Exec(`	ALTER TABLE cliente 	DROP CONSTRAINT cliente_pk;
						ALTER TABLE tarjeta 	DROP CONSTRAINT tarjeta_pk;
						ALTER TABLE comercio 	DROP CONSTRAINT comercio_pk;
						ALTER TABLE compra 		DROP CONSTRAINT compra_pk;
						ALTER TABLE rechazo 	DROP CONSTRAINT rechazo_pk;
						ALTER TABLE cierre 		DROP CONSTRAINT cierre_pk;
						ALTER TABLE cabecera 	DROP CONSTRAINT cabecera_pk;
						ALTER TABLE detalle 	DROP CONSTRAINT detalle_pk;
						ALTER TABLE alerta 		DROP CONSTRAINT alerta_pk;`)
	if err != nil {
		log.Fatal(err)
	}
}

func dropFKs() {
	_, err = db.Exec(`	ALTER TABLE tarjeta 	DROP CONSTRAINT tarjeta_nrocliente_fk;
						--ALTER TABLE rechazo 	DROP CONSTRAINT rechazo_nrotarjeta_fk;
						ALTER TABLE compra 		DROP CONSTRAINT compra_nrotarjeta_fk;
						--ALTER TABLE alerta 	DROP CONSTRAINT alerta_nrotarjeta_fk;
						ALTER TABLE cabecera 	DROP CONSTRAINT cabecera_nrotarjeta_fk;
						--ALTER TABLE alerta 	DROP CONSTRAINT alerta_nrorechazo_fk;
						ALTER TABLE rechazo 	DROP CONSTRAINT rechazo_nrocomercio_fk;
						ALTER TABLE compra 		DROP CONSTRAINT compra_nrocomercio_fk;`)
	if err != nil {
		log.Fatal(err)
	}
}

func generateCierres() {
	for nMes := 1; nMes <= 12; nMes++ {
		for terminacion := 0; terminacion <= 9; terminacion++ {
			var fInicio string
			var fCierre string
			var fVto string

			fInicio = fmt.Sprintf("2020-%v-%v", nMes, terminacion+2)
			if nMes == 12 {
				fCierre = fmt.Sprintf("2021-%v-%v", 1, terminacion+1)
				fVto = fmt.Sprintf("2021-%v-%v", 2, terminacion+11)
			} else {
				fCierre = fmt.Sprintf("2020-%v-%v", nMes+1, terminacion+1)
				fVto = fmt.Sprintf("2020-%v-%v", nMes+1, terminacion+11)
			}
			_, err = db.Exec(fmt.Sprintf("INSERT INTO cierre VALUES (2020, %v, %v, '%v', '%v', '%v');", nMes, terminacion, fInicio, fCierre, fVto))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func addStoredProceduresTriggers() {
	fmt.Println("Adding Stored Procedures and Triggers...")
	addAutorizacionDeCompra()
	addGenerarResumen()
	addCompraRechazadaTrigger()
	add2Compras1mMismoCpTrigger()
	add2Compras5mDistintoCpTrigger()
	add2RechazosPorExcesoLimiteTrigger()
	//addOtroTrigger()
	fmt.Println("Done adding Stored Procedures and Triggers!")
}




func addAutorizacionDeCompra() {
	fmt.Println(" Adding 'Autorizacion De Compra' Procedure")
	_, err = db.Exec(`	CREATE OR REPLACE FUNCTION autorizacion_de_compra(nrotarjetax char , codseguridadx char , nrocomerciox int , montox decimal) returns boolean as $$
						declare
							montoCompraSum int;
							tarjetaRecord record;
							fechaActual date;
							timeActual timestamp;
							nrechazo int;
							noperacion int;
							montoTotal int;
						
						begin
							SELECT COUNT(nrooperacion)+1 INTO noperacion FROM compra;
							SELECT COUNT(nrorechazo)+1 INTO nrechazo FROM rechazo;
							SELECT current_date INTO fechaActual;
						
							SELECT * FROM tarjeta INTO tarjetaRecord WHERE nrotarjeta = nrotarjetax;
						
							if not found then
								SELECT current_timestamp INTO timeActual;
								INSERT INTO rechazo VALUES (nrechazo, nrotarjetax, nrocomerciox, timeActual, montox, 'tarjeta no valida o no vigente', 0);
								return false;
							elsif tarjetaRecord.codseguridad != codseguridadx THEN
								SELECT current_timestamp INTO timeActual;
								INSERT INTO rechazo VALUES (nrechazo, nrotarjetax, nrocomerciox, timeActual, montox, 'codigo de seguridad invalido', 1);
								RETURN false;
							elsif CAST(tarjetaRecord.validahasta as date) < fechaActual THEN /* arreglar */
								SELECT current_timestamp INTO timeActual;
								INSERT INTO rechazo VALUES (nrechazo, nrotarjetax, nrocomerciox, timeActual, montox, 'plazo de vigencia expirado', 2);
								RETURN false;
							elsif tarjetaRecord.estado = 'suspendida' THEN
								SELECT current_timestamp INTO timeActual;
								INSERT INTO rechazo VALUES (nrechazo, nrotarjetax, nrocomerciox, timeActual, montox, 'la tarjeta se encuentra suspendida', 3);
								RETURN false;
							end if;
							
							SELECT SUM(monto) INTO montoCompraSum FROM compra WHERE nrotarjeta=nrotarjetax and pagado = false;
							montoTotal := montoCompraSum + montox;
						
							IF tarjetaRecord.limitecompra < montoTotal THEN
								SELECT current_timestamp into timeActual;
								INSERT INTO rechazo VALUES (nrechazo, nrotarjetax, nrocomerciox, timeActual, montox,'supera limite de tarjeta', 4);
								return false;
							END IF;
							
							SELECT current_timestamp INTO timeActual;
							INSERT INTO compra VALUES (noperacion, nrotarjetax, nrocomerciox, timeActual, montox, false);
							RETURN true;
						
						END;
						$$language plpgsql;`)
	if err != nil {
		log.Fatal(err)
	}
}

func addGenerarResumen() {
	fmt.Println(" Adding 'Generar resumen' Procedure")
	_, err = db.Exec(`  CREATE OR REPLACE FUNCTION generar_resumen(nroclientex int , mesx int , aniox int) returns void as $$
						declare 
							ncliente record;
							ntarjeta record;
							ncierre record;
							ncomercio record;
							unaCompra record;

							fechaEnDate date;

							tarjetaEnText text;
							ultimoDigito text;
							deudaTotal int;
							nresumen int;
							nlinea int;
							digito int;

						begin 
							SELECT COUNT(nroresumen) INTO nresumen FROM cabecera;

							SELECT * INTO ncliente FROM cliente WHERE nrocliente = nroclientex ;
							SELECT * INTO ntarjeta FROM tarjeta WHERE nrocliente = nroclientex and estado = 'vigente'; 
							tarjetaEnText := text (ntarjeta.nrotarjeta); /* paso a texto el numero de tarjeta*/
							SELECT right(tarjetaEnText,1) INTO ultimoDigito; /*el ultimo digito*/
							digito := to_number(ultimoDigito,'9');    /*9 es formato de mascara*/


							SELECT * into ncierre from cierre where anio = aniox and mes = mesx and terminacion = digito; 
							SELECT sum(monto) into deudaTotal from compra where nrotarjeta = ntarjeta.nrotarjeta and pagado = false;

							INSERT INTO cabecera VALUES (nresumen,ncliente.nombre,ncliente.apellido,ncliente.domicilio,ntarjeta.nrotarjeta,ncierre.fechainicio,ncierre.fechacierre,ncierre.fechavto,deudaTotal);


							FOR unaCompra IN SELECT * FROM compra WHERE nrotarjeta = ntarjeta.nrotarjeta loop
		
								SELECT * INTO ncomercio FROM comercio WHERE nrocomercio = unaCompra.nrocomercio;
								SELECT CAST (unaCompra.fecha AS date) INTO fechaEnDate;
								SELECT COUNT(nrolinea) INTO nlinea FROM detalle;
								INSERT INTO detalle VALUES (nresumen,nlinea,fechaEnDate,ncomercio.nombre,unaCompra.monto);
								unaCompra.pagado := true;
							END loop;
						END;
						$$ language plpgsql;`)
	if err != nil {
		log.Fatal(err)
	}
}

func addCompraRechazadaTrigger() {
	fmt.Println(" Adding 'Alerta Compra Rechazada' Procedure and trigger")
	_, err = db.Exec(`  CREATE OR REPLACE FUNCTION alerta_compra_rechazada() RETURNS TRIGGER AS $$
						DECLARE
							nalerta int;
						BEGIN
							SELECT MAX(nroalerta) + 1 INTO nalerta FROM alerta;
							IF nalerta ISNULL THEN 
								nalerta := 1; 
							END IF;
								INSERT INTO alerta VALUES (nalerta, new.nrotarjeta, new.fecha, new.nrorechazo, 0, 'Compra Rechazada');
							RETURN new;
						END;
						$$ language plpgsql;
						
						CREATE TRIGGER compra_rechazada
						BEFORE INSERT ON rechazo
						FOR EACH ROW
						EXECUTE PROCEDURE alerta_compra_rechazada();`)
	if err != nil {
		log.Fatal(err)
	}
}

func add2Compras1mMismoCpTrigger() {
	fmt.Println(" Adding 'Alerta Compra 1m mismo CP' Procedure and trigger")
	_, err = db.Exec(`  CREATE OR REPLACE FUNCTION alerta_compra_1m_mismoCP() RETURNS TRIGGER AS $$
						DECLARE
							nalerta int;
							ncompras int;
						BEGIN
							SELECT count(*) INTO ncompras 
							FROM compra AS cp
							JOIN comercio AS cm on cm.nrocomercio = cp.nrocomercio
							WHERE cp.nrotarjeta = new.nrotarjeta AND cp.nrocomercio != new.nrocomercio  AND cm.codigopostal = (SELECT codigopostal 
																														FROM comercio
																														WHERE new.nrocomercio = nrocomercio) AND new.fecha - cp.fecha <= INTERVAL '1' MINUTE;						
							IF ncompras = 1 then
								SELECT MAX(nroalerta)+1 INTO nalerta FROM alerta;
								IF nalerta ISNULL THEN 
									nalerta := 1; 
								END IF;
									INSERT INTO alerta VALUES (nalerta, new.nrotarjeta, new.fecha, null, 1, 'Se registraron dos compras en un lapso menor de un minuto en comercios distintos ubicados en el mismo codigo postal');
							END IF;
							RETURN new;
						END;
						$$ language plpgsql;
						
						CREATE TRIGGER compra_1m_mismoCP
						AFTER INSERT ON compra
						FOR EACH ROW
						EXECUTE PROCEDURE alerta_compra_1m_mismoCP();`)
	if err != nil {
		log.Fatal(err)
	}
}

func add2Compras5mDistintoCpTrigger() {
	fmt.Println(" Adding 'Alerta Compra 5m distinto CP' Procedure and trigger")
	_, err = db.Exec(`  CREATE OR REPLACE FUNCTION alerta_compra_5m_distintoCP() returns trigger as $$
						declare
							nalerta int;
							ncompras int;
						begin
							SELECT COUNT(*) INTO ncompras 
							FROM compra AS cp
							JOIN comercio AS cm on cm.nrocomercio = cp.nrocomercio
							WHERE cp.nrotarjeta = new.nrotarjeta AND cm.codigopostal != (SELECT codigopostal 
																						 FROM comercio
																						 WHERE new.nrocomercio = nrocomercio) AND new.fecha - fecha <= interval '5' minute;						
							if ncompras = 1 then
								SELECT MAX(nroalerta)+1 into nalerta from alerta;
								IF nalerta ISNULL THEN 
									nalerta := 1; 
								end if;
									INSERT INTO alerta VALUES (nalerta, new.nrotarjeta, new.fecha, null, 5, 'Se registraron dos compras en un lapso menor a 5 minutos en comercios con diferentes codigos postales');
							END IF;
							RETURN new;
						END;
						$$ language plpgsql;
						
						CREATE TRIGGER compra_5m_distintoCP
						AFTER INSERT ON compra
						FOR EACH ROW
						EXECUTE PROCEDURE alerta_compra_5m_distintoCP();`)
	if err != nil {
		log.Fatal(err)
	}
}

func add2RechazosPorExcesoLimiteTrigger() {
	fmt.Println(" Adding 'Alerta 2 compras rechazadas exceso limite' Procedure and trigger")
	_, err = db.Exec(`  CREATE OR REPLACE FUNCTION alerta_dos_rechazos_por_execeso_limite() returns trigger as $$
						DECLARE
							nalerta int;
							nrechazos int;
						BEGIN						
							SELECT COUNT(*) INTO nrechazos
							FROM rechazo AS rz
							WHERE rz.nrotarjeta = new.nrotarjeta AND 
								rz.codmotivo = 4 AND 
								rz.fecha BETWEEN date(new.fecha) AND date(new.fecha) + INTERVAL '23:59:59';
						
							IF nrechazos = 1 then
								UPDATE tarjeta SET estado = 'suspendida' where nrotarjeta = new.nrotarjeta;
								SELECT MAX(nroalerta)+1 INTO nalerta from alerta;
								IF nalerta ISNULL THEN 
									nalerta := 1; 
								end if;
									INSERT INTO alerta VALUES (nalerta, new.nrotarjeta, new.fecha, new.nrorechazo, 32, 'Se registraron dos rechazos por exceso de limite en el dia. La tarjeta ha sido suspendida preventivamente');
							END IF;
							RETURN new;
						END;
						$$ language plpgsql;
						
						CREATE TRIGGER compra_rechazada_exceso
						BEFORE INSERT ON rechazo
						FOR EACH ROW
						EXECUTE PROCEDURE alerta_dos_rechazos_por_execeso_limite();`)
	if err != nil {
		log.Fatal(err)
	}
}

func testeo(){
	fmt.Println("inciando testeteo de Consumos virtuales:")
	_, err = db.Exec(` 
						select consumos_virtuales(); `)

	if err != nil {
		log.Fatal(err)
	}
}

func testeoConsumosVirtuales() {
	fmt.Println(" Cargando Consumos Virtuales")
	_, err = db.Exec(`  CREATE OR REPLACE FUNCTION consumos_virtuales() returns void as $$
						DECLARE
							unConsumo record;
						BEGIN						
							for unConsumo in select * from consumo loop
								PERFORM autorizacion_de_compra(unConsumo.nrotarjeta,unConsumo.codseguridad,										unConsumo.nrocomercio,unConsumo.monto);
							end loop;
						END;
						$$ language plpgsql;`)
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
		[ 3 ] Remover PKs y FKs
		[ 4 ] test
		[ 5 ] Guardar los datos (BoltDB)

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
		advancedMenuBool = true
	case 3:
		dropPKandFK()
	case 4:
		fmt.Println("Hola, Test!")
		testeo()
	case 5:
		generarBoltDB()
	case 0:
		exitBool = true
		fmt.Println("Hasta Luego")
	default:
		fmt.Println("No elegiste ninguno")
	}
}
func advancedMenu() {
	menuString :=
		`
			Menu de creacion Manual
		[ 1 ] Eliminar Base tpgossz
		[ 2 ] Crear Base tpgossz
		[ 3 ] Conectar con Base tpgossz
		[ 4 ] Crear tablas
		[ 5 ] Agregar PKs y FKs
		[ 6 ] Popular Base de datos
		[ 7 ] Agregar Stored Procedures y Triggers
		[ 8 ] test

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
		addStoredProceduresTriggers()
	case 8:
		fmt.Println("Hola, Test!")
		//funcion a testear
	case 0:
		advancedMenuBool = false
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

//////////////////////////////////////////////////////////////////////////////5. JSON y Bases de datos NoSQL

//Cliente estructura
type Cliente struct {
	Nrocliente int
	Nombre     string
	Apellido   string
	Domicilio  string
	Telefono   int
}

//Tarjeta estructura
type Tarjeta struct {
	Nrotarjeta   int
	Nrocliente   int
	Validadesde  int `json:"Desde: "`
	Validahasta  int `json:"Hasta: "`
	Codseguridad int `json:"Codigo: "`
	Limitecompra float64
	Estado       string
}

//Comercio estructura
type Comercio struct {
	Nrocomercio int
	Nombre      string
	Domicilio   string
	Codpostal   string
	Telefono    int
}

//Compra estructura
type Compra struct {
	Nrooperacion int
	Nrotarjeta   int
	Nrocomercio  int
	Fecha        string
	Monto        float64
	Pagado       bool
}

func generarBoltDB() {
	db, err := bolt.Open("tpgossz.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//////////////////////////CLIENTES///////////////////////////////////////

	cliente1 := Cliente{1, "Leandro", "Sosa", "Marco Sastre 4540", 541152774600}

	dataCl1, err := json.Marshal(cliente1)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Clientes", []byte(strconv.Itoa(cliente1.Nrocliente)), dataCl1)
	resultadoCl1, err := readUnique(db, "Clientes", []byte(strconv.Itoa(cliente1.Nrocliente)))
	fmt.Printf("%s\n", resultadoCl1)

	cliente2 := Cliente{2, "Leonardo", "Sanabria", "Gaspar Campos 1815", 541148611570}

	dataCl2, err := json.Marshal(cliente2)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Clientes", []byte(strconv.Itoa(cliente2.Nrocliente)), dataCl2)
	resultadoCl2, err := readUnique(db, "Clientes", []byte(strconv.Itoa(cliente2.Nrocliente)))
	fmt.Printf("%s\n", resultadoCl2)

	cliente3 := Cliente{3, "Florencia", "Knol", "Zapiola 2825", 541148913800}

	dataCl3, err := json.Marshal(cliente3)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Clientes", []byte(strconv.Itoa(cliente3.Nrocliente)), dataCl3)
	resultadoCl3, err := readUnique(db, "Clientes", []byte(strconv.Itoa(cliente3.Nrocliente)))
	fmt.Printf("%s\n", resultadoCl3)

	//////////////////////////TARJETAS///////////////////////////////////////

	tarjeta1 := Tarjeta{5555899304583399, 1, 200911, 250221, 1234, 100000.90, "vigente"}

	dataT1, err := json.Marshal(tarjeta1)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Tarjetas", []byte(strconv.Itoa(tarjeta1.Nrotarjeta)), dataT1)
	resultadoT1, err := readUnique(db, "Tarjetas", []byte(strconv.Itoa(tarjeta1.Nrotarjeta)))
	fmt.Printf("%s\n", resultadoT1)

	tarjeta2 := Tarjeta{5269399188431044, 2, 190918, 240928, 0334, 50000, "vigente"}

	dataT2, err := json.Marshal(tarjeta2)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Tarjetas", []byte(strconv.Itoa(tarjeta2.Nrotarjeta)), dataT2)
	resultadoT2, err := readUnique(db, "Tarjetas", []byte(strconv.Itoa(tarjeta2.Nrotarjeta)))
	fmt.Printf("%s\n", resultadoT2)

	tarjeta3 := Tarjeta{8680402479723030, 3, 180322, 230322, 8214, 700000.12, "vigente"}

	dataT3, err := json.Marshal(tarjeta3)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Tarjetas", []byte(strconv.Itoa(tarjeta3.Nrotarjeta)), dataT3)
	resultadoT3, err := readUnique(db, "Tarjetas", []byte(strconv.Itoa(tarjeta3.Nrotarjeta)))
	fmt.Printf("%s\n", resultadoT3)

	//////////////////////////COMERCIOS///////////////////////////////////////

	comercio1 := Comercio{1, "Farmacia Tell", "Juncal 699", "B1663", 541157274612}

	dataCom1, err := json.Marshal(comercio1)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Comercios", []byte(strconv.Itoa(comercio1.Nrocomercio)), dataCom1)
	resultadoCom1, err := readUnique(db, "Comercios", []byte(strconv.Itoa(comercio1.Nrocomercio)))
	fmt.Printf("%s\n", resultadoCom1)

	comercio2 := Comercio{2, "Optica Bedini", "Peron 781", "B1871", 541174654172}

	dataCom2, err := json.Marshal(comercio2)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Comercios", []byte(strconv.Itoa(comercio2.Nrocomercio)), dataCom2)
	resultadoCom2, err := readUnique(db, "Comercios", []byte(strconv.Itoa(comercio2.Nrocomercio)))
	fmt.Printf("%s\n", resultadoCom2)

	comercio3 := Comercio{3, "Terravision", "Urquiza 1361", "B1221", 541183910808}

	dataCom3, err := json.Marshal(comercio3)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Comercios", []byte(strconv.Itoa(comercio3.Nrocomercio)), dataCom3)
	resultadoCom3, err := readUnique(db, "Comercios", []byte(strconv.Itoa(comercio3.Nrocomercio)))
	fmt.Printf("%s\n", resultadoCom3)

	/////////////////////////COMPRAS///////////////////////////////////////

	compra1 := Compra{1, 5555899304583399, 12, "2020-12-31 22:55:40", 2009.99, true}

	dataCpr1, err := json.Marshal(compra1)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Compras", []byte(strconv.Itoa(compra1.Nrooperacion)), dataCpr1)
	resultadoCpr1, err := readUnique(db, "Compras", []byte(strconv.Itoa(compra1.Nrooperacion)))
	fmt.Printf("%s\n", resultadoCpr1)

	compra2 := Compra{2, 5269399188431044, 15, "2020-04-16 12:25:40", 500.45, true}

	dataCpr2, err := json.Marshal(compra2)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Compras", []byte(strconv.Itoa(compra2.Nrooperacion)), dataCpr2)
	resultadoCpr2, err := readUnique(db, "Compras", []byte(strconv.Itoa(compra2.Nrooperacion)))
	fmt.Printf("%s\n", resultadoCpr2)

	compra3 := Compra{3, 8680402479723030, 7, "2020-09-01 20:16:40", 1000000.00, false}

	dataCpr3, err := json.Marshal(compra3)

	if err != nil {
		log.Fatal(err)
	}

	createUpdate(db, "Compras", []byte(strconv.Itoa(compra3.Nrooperacion)), dataCpr3)
	resultadoCpr3, err := readUnique(db, "Compras", []byte(strconv.Itoa(compra3.Nrooperacion)))
	fmt.Printf("%s\n", resultadoCpr3)

}

func createUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
	// abre transacción de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
	if err != nil {
		return err
	}

	err = b.Put(key, val)
	if err != nil {
		return err
	}

	// cierra transacción
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func readUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
	var buf []byte

	// abre una transacción de lectura
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})

	return buf, err
}
