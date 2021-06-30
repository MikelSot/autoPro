package route

import (
	"github.com/MikelSot/autoPro/handler"
	"github.com/MikelSot/autoPro/middleware"
	"github.com/labstack/echo"
)

type iProduct handler.IProductCRUDQuery
type iEmployee handler.IEmployeeCRUDExists
type iService handler.IServiceCRUDQuery
type iWorkshop handler.IWorkshopCRUD

func Login(e *echo.Echo,  storage handler.IClientCRUDExists) {
	h := handler.NewLogin(storage)
	c := handler.NewClientHd(storage)
	e.POST("/v1/register",c.SingIn)
	e.POST("/v1/login", h.Login)
}

func Home(e *echo.Echo,iProd iProduct, iEmp iEmployee, ISer iService, iWork iWorkshop)  {
	workshop := handler.NewWorkshopHd(iWork)
	service := handler.NewServiceHd(ISer)
	employee := handler.NewEmployeeHd(iEmp)
	product := handler.NewProductHd(iProd)

	e.GET("/v1/product-home/:max", product.GetAll)
	e.GET("/v1/workshop-home/:max", workshop.GetAll)
	e.GET("/v1/service-home/:max", service.GetAll)
	e.GET("/v1/employee-home/:max", employee.DataEmployeeHome)
}

func Client(e *echo.Echo, storage handler.IClientCRUDExists)  {
	c := handler.NewClientHd(storage)
	client := e.Group("/v1/client")
	client.Use(middleware.Authentication)
	client.PUT("/edit-profile",c.EditClient)
	e.GET("/:id", c.GetById)
	e.GET("/all-client/:max", c.GetAll)
	e.DELETE("/:id", c.DeleteSoft)
}

func Blog(e *echo.Echo, storage handler.IBlogCRUDQuery, comm handler.ICommentCRUDQuery)  {
	b := handler.NewBlogHd(storage)
	c := handler.NewCommentHd(comm)
	blog := e.Group("/v1/blog")
	blog.Use(middleware.Authentication)
	
	blog.GET("/:max", b.GetAll)
	blog.GET("/view-blog/:id", b.GetById)
	blog.PUT("/view-blog/:id", b.Update)
	blog.DELETE("/view-blog/:id", b.DeleteSoft)
	blog.GET("/view-comment-blog/:id/:max", c.AllCommentBlog)
	blog.GET("/view-comment-product/:id/:max", c.AllCommentProduct)
	blog.POST("/create-comment", c.Create)
	blog.DELETE("/create-comment", c.DeleteSoft)
}

func Invoice(e *echo.Echo, inv handler.IInvoiceCRUDQuery, item handler.IInvoiceItemCRUDQuery)  {
	i := handler.NewInvoiceHd(inv)
	it := handler.NewInvoiceItemHd(item)
	blog := e.Group("/v1/invoice")
	blog.Use(middleware.Authentication)

	blog.GET("/:id/:max", i.AllInvoiceClient)
	blog.DELETE("/:id", i.DeleteSoft)
	blog.GET("/id/:id", i.GetById)
	blog.PUT("/id/:id", i.Update)
	blog.GET("/id-item/:id/:max", it.AllInvoiceItemInvoice)
	blog.DELETE("/id-item/:id", it.DeleteSoft)
	blog.PUT("/id-item/:id", it.Update)
}

func Product(e *echo.Echo, prod handler.IProductCRUDQuery)  {
	p := handler.NewProductHd(prod)
	product := e.Group("/v1/product")
	product.Use(middleware.Authentication)

	product.GET("/:max", p.GetAll)
	product.GET("/product-id/:id", p.GetAll)
	product.PUT("/product-id/:id", p.Update)
	product.GET("/product-id/:id/:max", p.AllProductsCategory)
}

func AppointmentReview(e *echo.Echo, rev handler.ITechnicalReviewCRUDQuery, appoint handler.IAppointmentCRUDQuery)  {
	r := handler.NewTechnicalReviewHd(rev)
	a := handler.NewAppointmentHd(appoint)
	appointment := e.Group("/v1/appointment")
	appointment.GET("/:id/:max", a.AllAppointmentClient)
	appointment.POST("", a.Create)
	appointment.DELETE("/:id", a.DeleteSoft)
	appointment.PUT("/:id", a.Update)
	appointment.GET("/order-available/", a.AllOrderAttentionAvailable)
	appointment.GET("/review/:id/:max", r.AllReviewClient)
	appointment.GET("/review/:id", r.GetById)
	appointment.POST("/review", r.Create)

}