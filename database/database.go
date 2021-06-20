package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once

)

const (
	ZeroRowsAffected = 0
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


// Migration crea las tablas y ademas ingresa datos
func Migration() error  {
	err,_ :=migrationAndInsert()
	if err != nil {
		log.Fatalf("%v", err)
		return errors.New("Error al crear las tablas o al ingresar datos a las tablas")
	}
	return nil
}

// DB returna una unica instancia de la conexion a la base de datos
func DB() *gorm.DB {
	return db
}

