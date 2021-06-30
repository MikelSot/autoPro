package route

import (
	"github.com/MikelSot/autoPro/handler"
	"github.com/labstack/echo"
)

type iProduct handler.IProductCRUDQuery
type iEmployee handler.IEmployeeCRUDExists
type iService handler.IServiceCRUDQuery
type iWorkshop handler.IWorkshopCRUD

func homeRo(e *echo.Echo,iProd iProduct, iEmp iEmployee, ISer iService, iWork iWorkshop)  {
	workshop := handler.NewWorkshopHd(iWork)
	service := handler.NewServiceHd(ISer)
	employee := handler.NewEmployeeHd(iEmp)
	product := handler.NewProductHd(iProd)

	home := e.Group("/api/v1")
	// api/v1/ [lo que sigue]
	home.GET("/:max", workshop.GetAll)
	home.GET("/:max", service.GetAll)
	home.GET("/:max", employee.DataEmployeeHome)
	home.GET("/:max", product.GetAll)
}
