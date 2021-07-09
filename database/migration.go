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
		{Name: "Soldado", Description: " Soldar una parte dañada del vehículo (quebrada)", Picture: "/uploads/img/services/pexels-cottonbro-4489743.jpg"},
		{Name: "Creación de partes del vehículo ", Description: "Si una parte del vehículo no se encuentra en repuestos, se procede a crear uno nuevo", Picture: "/uploads/img/services/pexels-chevanon-photography-1108101.jpg"},
		{Name: "Cambio de llantas Llanta", Description: "Llanta desinfladas o dañadas se cambian por nuevas", Picture: "/uploads/img/services/pexels-andrea-piacquadio-3806249.jpg"},
		{Name: "mantenimiento del vehículo", Description: "Revisar solo motor del vehículo", Picture: "/uploads/img/services/auto-repair-3691962_1280.jpg"},
		{Name: "Reparación de partes del vehículo", Description: " Reparar la parte dañada del vehículo", Picture: "/uploads/img/services/auto-2861859_1920.jpg"},
		{Name: "Parchado", Description: "Rellenar las partes que faltan en el vehículo ", Picture: "/uploads/img/services/hands-of-car-mechanic-with-wrench-in-garage.jpg"},
		{Name: "Cambio de aceite", Description: "Cambiar el aceite del vehículo por uno nuevo"},
		{Name: "Auxilio mecánico", Description: ""},
		{Name: "Venta de repuestos", Description: "Repuestos nuevos en venta de las partes de vehículos"},
		{Name: "Cambio de muelle", Description: "Cambiar un muelle dañado o quebrado"},
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
		{Name: "Tubo de escape", Stock: 15, CategoryID: 2, UnitPrice: 100, Picture: "/uploads/img/products/exhaust-3538388_1920.jpg"},
		{Name: "Llantas", Stock: 20, CategoryID: 1, UnitPrice: 180, Picture: "/uploads/img/products/mature-2875251_1920.jpg"},
		{Name: "Tornillos", Stock: 100, CategoryID: 4, UnitPrice: 4.5, Picture: "/uploads/img/products/pexels-pixabay-162553.jpg"},
		{Name: "Filtro", Stock: 50, CategoryID: 4, UnitPrice: 1.5, Picture: "/uploads/img/products/mqdefault.jpg"},
		{Name: "Silenciador", Stock: 10, CategoryID: 2, UnitPrice: 80, Picture: "/uploads/img/products/H4cc16d632d9c4d3aa41ef40d3f1fb4008.jpg"},
		{Name: "Tuercas", Stock: 10, CategoryID: 2, UnitPrice: 80, Picture: "/uploads/img/products/pexels-cottonbro-4480531.jpg"},
		{Name: "Muelle", Stock: 30, CategoryID: 2, UnitPrice: 120},
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
