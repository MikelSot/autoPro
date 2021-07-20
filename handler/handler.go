package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
)


// IClientCRUDExists embebida de cliente
type IClientCRUDExists interface {
	IClientCRUD
	IQueryExists
}

// IClientCRUDExists embebida de empleado
type IEmployeeCRUDExists interface {
	IEmployeeCRUD
}

// IClientCRUDExists embebida de cita
type IAppointmentCRUDQuery interface {
	IAppointmentCRUD
	IQueryAppointment
}

type IServiceCRUDQuery interface {
	IServiceCRUD
}

type IProductCRUDQuery interface {
	IProductCRUD
	IQueryProduct
}

type ITechnicalReviewCRUDQuery interface {
	ITechnicalReviewCRUD
	IQueryReview
}

type IBlogCRUDQuery interface {
	IBlogCRUD
	IQueryBlog
}

type ICommentCRUDQuery interface {
	ICommentCRUD
	IQueryComment
}

type IInvoiceCRUDQuery interface {
	IInvoiceCRUD
	IQueryInvoice
}

type IInvoiceItemCRUDQuery interface {
	IInvoiceItemCRUD
	IQueryInvoiceItem
}

// IClient interface de CRUD
type IClientCRUD interface {
	Create(client *dto.SignInClient) error
	Update(ID uint, client *dto.EditClient) error
	GetByID(ID uint) (*model.Client, error)
	GetAll(max int) (*model.Clients, error)
	UpdatePicture(ID uint, rute string) error
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

// IQueryExists esta interface tiene metodos para validar si ya existe determinado atributo.
type IQueryExists interface {
	QueryEmailExists(email string) (bool,model.Client, error)
	QueryDniExists(dni string) (bool,uint, error)
	QueryUriExists(uri string) (bool, error)
}

// IEmployeeCRUD interface crud de empleado
type IEmployeeCRUD interface {
	Create(employee *model.Employee) error
	Update(ID uint, employee *model.Employee) error
	GetByID(ID uint) (*dto.AllDataEmployee, error)
	GetAll(max int) (*model.Employees, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
	DataEmployeeHome(max int) (*dto.DataEmployeeHomes,error)
	QueryEmailEqualsClient(email string) (uint,error)
	QueryEmailExists(email string) (bool, model.Employee, error)
}

// IAppointmentCRUD interface crud de citas
type IAppointmentCRUD interface {
	Create(appointment *dto.AppointmentCreate) error
	Update(ID uint, appointment *dto.AppointmentUpdate) error
	UpdateState(ID uint, dto *dto.AppointmentUpdateState) error
	GetAll(max int) (*model.Appointments, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

// IQueryAppointment consultas para citas
type IQueryAppointment interface {
	AllOrderAttentionAvailable() (map[int]string, error)
	AllAppointmentClient(ID uint, max int) (*model.Appointments, error)
	QueryWorkshopExists(name string) (bool, error)
	QueryServiceExists(name string) (bool, error)
}

// IBlogCRUD interface drud de blog
type IBlogCRUD interface {
	Create( blog *model.Blog) error
	Update(ID uint,  blog *model.Blog) error
	GetByID(ID uint) (*model.Blog, error)
	GetAll(max int) (*model.Blogs, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

// IQueryBlog interface de consulta de blog
type IQueryBlog interface {
	AllBlogCategory(ID, max int) (*model.Blogs, error)
	AllBlogEmployee(ID uint, max int) (*model.Blogs, error)
}

// ICommentCRUD interface crud de comentarios
type ICommentCRUD interface {
	Create(comment *model.Comment) error
	Update(ID uint,comment *model.Comment) error
	DeleteSoft(ID uint) error
}

// IQueryComment interface de consulta de blog
type IQueryComment interface {
	AllCommentBlog(ID, max int) (*dto.CommentClients, error) // ID blog
	AllCommentProduct(ID, max int) (*dto.CommentClients, error) // ID product
}


// IInvoiceCRUD interface crud de factura
type IInvoiceCRUD interface {
	Create(invoice *model.Invoice) error
	Update(ID uint, invoice *model.Invoice) error
	GetByID(ID uint) (*model.Invoice, error)
	DeleteSoft(ID uint) error
}

// IQueryInvoice interface consulta de factura
type IQueryInvoice interface {
	AllInvoiceClient(ID, max int) (*dto.InvoiceClients, error)
	AllInvoiceWorkshop(ID uint, max int) (*dto.InvoiceWorkshops, error)
}

// IInvoiceItemCRUD interface item de factura
type IInvoiceItemCRUD interface {
	Create(item *model.InvoiceItem) error
	Update(ID uint, item *model.InvoiceItem) error
	GetByID(ID uint) (*model.InvoiceItem,error)
	UpdateStock(ID ,stock uint) error
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

// IQueryInvoiceItem interface de consulta de item de factura
type IQueryInvoiceItem interface {
	AllInvoiceItemInvoice(ID, max int) (*model.InvoiceItems, error)
}

// IPaymentMethodCRUD interface crud de metodo de pago
type IPaymentMethodCRUD interface {
	Create(method *model.PaymentMethod) error
	Update(ID uint8, method *model.PaymentMethod) error
	GetByID(ID uint8) (*model.PaymentMethod, error)
	GetAll(max int) (*model.PaymentMethods, error)
	DeleteSoft(ID uint8) error
	DeletePermanent(ID uint8) error
}

// IProductCRUD interface crud de producto
type IProductCRUD interface {
	Create( product *model.Product) error
	Update(ID uint,  product *model.Product) error
	GetByID(ID uint) (*model.Product, error)
	GetAll(max int) (*model.Products, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

// IQueryProduct interface de consulta de producto
type IQueryProduct interface {
	AllProductsCategory(ID uint, max int) (dto.ProductClients, error)
	AllProductsWorkshop(ID uint, max int) (dto.ProductClients, error)
}


// IRoleCRUD  interface crud de role
type IRoleCRUD interface {
	Create(role *model.Role) error
	Update(ID uint8, role *model.Role) error
	GetByID(ID uint8) (*model.Role, error)
	GetAll(max int) (*model.Roles, error)
	DeleteSoft(ID uint8) error
	DeletePermanent(ID uint8) error
}

// IServiceCRUD interface crud de servicios
type IServiceCRUD interface {
	Create(service *model.Service) error
	Update(ID uint8,service *model.Service) error
	GetByID(ID uint8) (*model.Service, error)
	GetAll(max int) (*model.Services, error)
	DeleteSoft(ID uint8) error
	DeletePermanent(ID uint8) error
}

// ITechnicalReviewCRUD interface de revicion tecnica
type ITechnicalReviewCRUD interface {
	Create(review *model.TechnicalReview) error
	Update(ID uint, review *model.TechnicalReview) error
	GetByID(ID uint) (*model.TechnicalReview, error)
	GetAll(max int) (*model.TechnicalReviews, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

type IQueryReview interface {
	AllReviewClient(ID uint, max int) (*model.TechnicalReviews, error)
}


// IWorkshopCRUD interface crud de taller
type IWorkshopCRUD interface {
	Create(workshop *model.Workshop) error
	Update(ID uint8, workshop *model.Workshop) error
	GetByID(ID uint8) (*model.Workshop, error)
	GetAll(max int) (*model.Workshops, error)
	DeleteSoft(ID uint8) error
	DeletePermanent(ID uint8) error
}
