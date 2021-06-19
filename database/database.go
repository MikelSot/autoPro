package database

import (
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
	return DB().AutoMigrate(
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
}


// DB returna una unica instancia de la conexion a la base de datos
func DB() *gorm.DB {
	return db
}

func Migration()  {
	err := createTable()
	if err != nil {
		log.Fatalf("error al hace la migracion de los modelos a trablas %v", err)
	}
}



