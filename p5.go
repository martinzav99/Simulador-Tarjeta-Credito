package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	bolt "github.com/coreos/bbolt"
	//bolt "go.etcd.io/bbolt"
)

//////////////////////////////////5. JSON y Bases de datos NoSQL

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

	createUpdate(db, "Clientes", []byte(strconv.Itoa(cliente1.Nrocliente)), data_cl1)
	resultadoCl1, err := readUnique(db, "Clientes", []byte(strconv.Itoa(cliente1.Nrocliente)))
	fmt.Printf("%s\n", resultado_cl1)

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
