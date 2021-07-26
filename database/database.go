package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
	onceMigration sync.Once
)

const (
	ZeroRowsAffected = 0
	MaxGetAll        = 10
	MaxComment       = 10
	MaxGetAllHome	 = 6
)


// connectionDB conexion a la base de datos, singleton
func connectionDB() {
	once.Do(func() {
		var err error
		dsn := "host=localhost user=mike password=cuUlLyVD9kS4V39qm1tmpU5S4MvWiUiHhU8 dbname=autoprodb port=5432 sslmode=disable TimeZone=America/Lima"
		//dsn := "host=localhost user=me-postgresql password=cmd.08miguel dbname=autoProdb port=5432 sslmode=disable TimeZone=America/Lima"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("no se pudo conectar a la base de datos --> %v", err)
		}

		fmt.Println("CONECTADO!!")
	})
}


// Migration crea las tablas y ademas ingresa datos
func Migration() error  {
	var err error
	onceMigration.Do(func() {
		err,_ =migrationAndInsert()
	})

	fmt.Println("DATABASE --> MIGRATION")
	return nil
}

// DB returna una unica instancia de la conexion a la base de datos
func DB() *gorm.DB {
	return db
}

