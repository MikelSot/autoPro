package route

import (
	"github.com/MikelSot/autoPro/handler"
	"github.com/MikelSot/autoPro/middleware"
	"github.com/labstack/echo/v4"
)

type iProduct handler.IProductCRUDQuery
type iEmployee handler.IEmployeeCRUDExists
type iService handler.IServiceCRUDQuery
type iWorkshop handler.IWorkshopCRUD

func Login(e *echo.Echo, storage handler.IClientCRUDExists) {
	h := handler.NewLogin(storage)
	c := handler.NewClientHd(storage)
	e.POST("/v1/register", c.SingIn)
	e.POST("/v1/login", h.Login)
}

func Home(e *echo.Echo, iProd iProduct, iEmp iEmployee, ISer iService, iWork iWorkshop, appoint handler.IAppointmentCRUDQuery) {
	workshop := handler.NewWorkshopHd(iWork)
	service := handler.NewServiceHd(ISer)
	employee := handler.NewEmployeeHd(iEmp)
	product := handler.NewProductHd(iProd)
	appointment := handler.NewAppointmentHd(appoint)


	e.GET("/v1/product-home/:max", product.GetAll)
	e.GET("/v1/workshop-home/:max", workshop.GetAll)
	e.GET("/v1/service-home/:max", service.GetAll)
	e.GET("/v1/employee-home/:max", employee.DataEmployeeHome)
	e.GET("/v1/order-available-home/", appointment.AllOrderAttentionAvailable)

}

func Client(e *echo.Echo, storage handler.IClientCRUDExists) {
	c := handler.NewClientHd(storage)
	client := e.Group("/v1/client")
	client.Use(middleware.Authentication)

	client.PUT("/edit-profile", c.EditClient)
	//client.PUT("/edit-profile/:id", c.EditClient)
	client.GET("/:id", c.GetById)
	client.POST("/upload-file", c.UploadAvatar)
	client.GET("/all-client/:max", c.GetAll)
	client.DELETE("/:id", c.DeleteSoft)
	client.GET("/id-name", c.SelectNameID)
}

func AppointmentReview(e *echo.Echo, rev handler.ITechnicalReviewCRUDQuery, appoint handler.IAppointmentCRUDQuery) {
	r := handler.NewTechnicalReviewHd(rev)
	a := handler.NewAppointmentHd(appoint)
	appointment := e.Group("/v1/appointment")
	appointment.Use(middleware.Authentication)
	appointment.GET("/:id/:max", a.AllAppointmentClient)
	appointment.GET("/all/:max", a.GetAll)
	appointment.POST("", a.Create)
	appointment.DELETE("/:id", a.DeleteSoft)
	appointment.PUT("/:id", a.Update)
	appointment.PUT("/state/:id", a.UpdateState)
	appointment.GET("/order-available/", a.AllOrderAttentionAvailable)
	appointment.GET("/review/:id/:max", r.AllReviewClient)
	appointment.GET("/review/:id", r.GetById)
	appointment.POST("/review", r.Create)
	appointment.PUT("/review/:id", r.Update)
	appointment.DELETE("/review/:id", r.DeleteSoft)
	appointment.GET("/review-all/:max", r.GetAll)
}

func Blog(e *echo.Echo, storage handler.IBlogCRUDQuery, comm handler.ICommentCRUDQuery) {
	b := handler.NewBlogHd(storage)
	c := handler.NewCommentHd(comm)
	blog := e.Group("/v1/blog")
	blog.Use(middleware.Authentication)

	blog.GET("/:max", b.GetAll)
	blog.GET("/view-blog/:id", b.GetById)
	blog.PUT("/view-blog/:id", b.Update)
	blog.DELETE("/view-blog/:id", b.DeleteSoft)
	blog.GET("/view-blog-category/:id/:max", b.AllBlogCategory)
	blog.GET("/view-blog-employee/:id/:max", b.AllBlogEmployee)
	blog.GET("/view-comment-blog/:id/:max", c.AllCommentBlog)
	blog.POST("/create-comment", c.Create)
	blog.PUT("/create-comment/:id", c.Update)
	blog.DELETE("/create-comment/:id", c.DeleteSoft)
}

func Invoice(e *echo.Echo, inv handler.IInvoiceCRUDQuery, item handler.IInvoiceItemCRUDQuery) {
	i := handler.NewInvoiceHd(inv)
	it := handler.NewInvoiceItemHd(item)
	invoice := e.Group("/v1/invoice")
	invoice.Use(middleware.Authentication)

	invoice.GET("/:id/:max", i.AllInvoiceClient)
	invoice.GET("/workshop/:id/:max", i.AllInvoiceWorkshop)
	invoice.DELETE("/delete/:id", i.DeleteSoft)
	invoice.GET("/id/:id", i.GetById)
	invoice.PUT("/id/:id", i.Update)
	invoice.POST("/create", i.Create)
	invoice.GET("/all-item/:id/:max", it.AllInvoiceItemInvoice)
	invoice.DELETE("/item/:id", it.DeleteSoft)
	invoice.PUT("/item/:id", it.Update)
	invoice.POST("/item", it.Create)
}

func Product(e *echo.Echo, prod handler.IProductCRUDQuery, comm handler.ICommentCRUDQuery) {
	p := handler.NewProductHd(prod)
	c := handler.NewCommentHd(comm)
	product := e.Group("/v1/product")
	product.Use(middleware.Authentication)

	product.POST("/create", p.Create)
	product.GET("/:max", p.GetAll)
	product.PUT("/id/:id", p.Update)
	product.GET("/id/:id", p.GetById)
	product.DELETE("/product-id/:id", p.DeleteSoft)
	product.GET("/category/:id/:max", p.AllProductsCategory)
	product.GET("/workshop/:id/:max", p.AllProductsWorkshop)
	product.GET("/all-comment/:id/:max", c.AllCommentProduct)
	product.POST("/comment", c.Create)
	product.PUT("/comment", c.Update)
	product.DELETE("/comment", c.DeleteSoft)
}

func Employee(e *echo.Echo, emp handler.IEmployeeCRUDExists)  {
	em := handler.NewEmployeeHd(emp)
	employee := e.Group("/v1/employee")
	employee.Use(middleware.Authentication)

	employee.POST("/create", em.Create)
	employee.PUT("/update/:id", em.Update)
	employee.GET("/id/:id", em.GetById)
	employee.GET("/all/:max", em.GetAll)
	employee.DELETE("/all/:max", em.DeleteSoft)
	employee.GET("/all-data/:max", em.DataEmployeeHome)
}

func PaymentMethod(e *echo.Echo, pm handler.IPaymentMethodCRUD)  {
	p := handler.NewPaymentMethodHd(pm)
	method := e.Group("/v1/method")
	method.Use(middleware.Authentication)

	method.POST("", p.Create)
	method.PUT("/:id", p.Update)
	method.DELETE("/:id", p.DeleteSoft)
	method.GET("/id/:id", p.GetById)
	method.GET("/all/:id", p.GetAll)
}

func Role(e *echo.Echo, rl handler.IRoleCRUD)  {
	r := handler.NewRoleHd(rl)
	role := e.Group("/v1/role")
	role.Use(middleware.Authentication)

	role.POST("/:id", r.Create)
	role.PUT("/:id", r.Update)
	role.DELETE("/:id", r.DeleteSoft)
	role.GET("/id/:id", r.GetById)
	role.GET("/all/:id", r.GetAll)
}

func Service(e *echo.Echo, sr handler.IServiceCRUDQuery)  {
	s := handler.NewServiceHd(sr)
	service := e.Group("/v1/service")
	service.Use(middleware.Authentication)

	service.POST("/:id", s.Create)
	service.PUT("/:id", s.Update)
	service.DELETE("/:id", s.DeleteSoft)
	service.GET("/id/:id", s.GetById)
	service.GET("/all/:id", s.GetAll)
}