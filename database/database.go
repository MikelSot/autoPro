package database

import (
	"errors"
	"fmt"
	"github.com/MikelSot/autoPro/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once

	ErrorWhenInsertingIntoTable = errors.New("error al insertar datos en la tabla")
)

// connectionDB conexion a la base de datos, singleton
func connectionDB() {
	once.Do(func() {
		var err error
		dsn := "host=localhost user=me-postgresql password=cmd.08miguel dbname=autoProdb port=5432 sslmode=disable TimeZone=America/Lima"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("no se pudo conectar a la base de datos --> %v", err)
		}

		fmt.Println("CONECTADO!!")
	})
}

//createTable esta funcion nos crehara las tablas en la base de datos postgresql
func createTable() error {
	connectionDB()
	err := DB().AutoMigrate(
		&model.Role{},
		&model.Workshop{},
		&model.Service{},
		&model.PaymentMethod{},
		&model.Category{},
		&model.Employee{},
		&model.Product{},
		&model.Client{},
		&model.Appointment{},
		&model.Blog{},
		&model.Comment{},
		&model.Invoice{},
		&model.InvoiceItem{},
		&model.TechnicalReview{},
	)

	if err != nil{
		return err
	}
	return nil
}


// DB returna una unica instancia de la conexion a la base de datos
func DB() *gorm.DB {
	return db
}


func insertDataRole() error {
	roles := model.Roles{
		{Name: "admin"},
		{Name: "cliente"},
		{Name: "invitado"},
		{Name: "empleado-normal"},
	}
	DB().Select("Name").Create(&roles)
	return nil
}

func insertDataWorkshop() error {

	workshops := model.Workshops{
		{Name:"Mario", Address: "Jr. Los platos - Mala"},
		{Name:"Jesús" , Address: "Av. Progreso - Mala"},
		{Name:"Alejandro" , Address: "Av 9 de octubre - Mala"},
	}
	DB().Select("Name", "Address").Create(&workshops)
	return nil
}


//func insertDataService() error {
//	services := model.Services{
//		{},
//	}
//}

func migrationAndInsert() (error,bool) {
	var flag = false
	var err error
	err = createTable()
	if err != nil {
		log.Fatalf("error al hace la migracion de los modelos a trablas %v", err)
		return err, flag
	}

	//err = insertDataRole()
	err = insertDataWorkshop()

	if err != nil {
		return err, flag
	}

	flag = true
	return nil, flag
}

func Migration() error  {
	err,_ :=migrationAndInsert()
	if err != nil {
		log.Fatalf("%v", err)
		return errors.New("Error al crear las tablas o al ingresar datos a las tablas")
	}
	return nil
}

