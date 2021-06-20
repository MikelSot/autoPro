package database

import (
	"github.com/MikelSot/autoPro/model"
	"log"
)

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

	if err != nil {
		return err
	}
	return nil
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
		{Name: "Mario", Address: "Jr. Los platos - Mala"},
		{Name: "Jesús", Address: "Av. Progreso - Mala"},
		{Name: "Alejandro", Address: "Av 9 de octubre - Mala"},
	}
	DB().Select("Name", "Address").Create(&workshops)
	return nil
}

func insertDataService() error {
	services := model.Services{
		{Name: "mantenimiento del vehículo", Description: "Revisar solo motor del vehículo"},
		{Name: "Reparación de partes del vehículo", Description: " Reparar la parte dañada del vehículo"},
		{Name: "Soldado", Description: " Soldar una parte dañada del vehículo (quebrada)"},
		{Name: "Parchado", Description: "Rellenar las partes que faltan en el vehículo "},
		{Name: "Cambio de aceite", Description: "Cambiar el aceite del vehículo por uno nuevo"},
		{Name: "Venta de repuestos", Description: "Repuestos nuevos en venta de las partes de vehículos"},
		{Name: "Cambio de muelle", Description: "Cambiar un muelle dañado o quebrado"},
		{Name: "Cambio de llantas Llanta", Description: "Llanta desinfladas o dañadas se cambian por nuevas"},
		{Name: "Creación de partes del vehículo ", Description: "Si una parte del vehículo no se encuentra en repuestos, se procede a crear uno nuevo"},
		{Name: "Auxilio mecánico", Description: ""},
	}

	DB().Select("Name", "Description").Create(&services)

	service_workshops := []model.Service_Workshops{
		{1, 1}, {1, 2}, {1, 3},
		{2, 1}, {2, 2}, {2, 3},
		{3, 1}, {3, 2}, {3, 3},
		{4, 1}, {4, 2}, {4, 3},
		{5, 1}, {5, 2}, {5, 3},
		{6, 1}, {6, 2}, {6, 3},
		{7, 1}, {7, 2}, {7, 3},
		{8, 1}, {8, 2}, {8, 3},
		{9, 1}, {9, 2}, {9, 3},
		{10, 1}, {10, 2}, {10, 3},
	}

	DB().Create(&service_workshops)
	return nil
}

func insertDataCategory() error {
	categories := model.Categories{
		{Name: "partes del vehículo"},
		{Name: "repuestos para vehículo"},
		{Name: "repuesto en general"},
		{Name: "Tornillos"},
	}
	DB().Select("Name").Create(&categories)
	return nil
}

func insertDataProduct() error {
	products := model.Products{
		{Name: "Llantas", Stock: 20, CategoryID: 1, UnitPrice: 180},
		{Name: "Muelle", Stock: 30, CategoryID: 2, UnitPrice: 120},
		{Name: "Tornillos", Stock: 100, CategoryID: 4, UnitPrice: 4.5},
		{Name: "Tuercas", Stock: 50, CategoryID: 4, UnitPrice: 1.5},
		{Name: "Tubo de escape", Stock: 15, CategoryID: 2, UnitPrice: 100},
		{Name: "Silenciador", Stock: 10, CategoryID: 2, UnitPrice: 80},
	}

	DB().Select("Name", "Stock", "CategoryID", "UnitPrice").Create(&products)

	return nil
}

func migrationAndInsert() (error, bool) {
	var flag = false
	var err error
	err = createTable()
	if err != nil {
		log.Fatalf("error al hace la migracion de los modelos a trablas %v", err)
		return err, flag
	}

	err = insertDataRole()
	err = insertDataWorkshop()
	err = insertDataService()
	err = insertDataCategory()
	err = insertDataProduct()

	if err != nil {
		return err, flag
	}

	flag = true
	return nil, flag
}
