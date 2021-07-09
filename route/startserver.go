package route

import (
	"github.com/MikelSot/autoPro/database"
	"github.com/MikelSot/autoPro/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func StartServer()  {
	err := jwt.LoadFiles("jwt/app.rsa", "jwt/app.rsa.public")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}

	database.Migration()
	client := database.NewClientDao()
	product := database.NewProductDao()
	service := database.NewServiceDao()
	workshop := database.NewWorkshopDao()
	blog := database.NewBloDao()
	invoice := database.NewInvoiceDao()
	invoiceItem := database.NewInvoiceItemDao()
	employee := database.NewEmployeeDao()
	comment := database.NewCommentDao()
	technicalReview := database.NewTechnicalReviewDao()
	appointment := database.NewAppointmentDao()
	method := database.NewPaymentMethodDao()
	role := database.NewRoleDao()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	Login(e, &client)
	Home(e, &product,&employee, &service, &workshop)
	Client(e, &client)
	Blog(e,&blog, &comment)
	Invoice(e, &invoice,&invoiceItem)
	Product(e, &product, &comment)
	AppointmentReview(e, &technicalReview, &appointment)
	PaymentMethod(e, &method)
	Employee(e, &employee)
	Role(e, &role)
	Service(e, &service)

	err = e.Start(":8080")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
