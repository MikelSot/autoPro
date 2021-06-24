package dto

// para intercanbiar info entre usuario y para mostrarlo al cliente
type ProductClients []*ProductClient
type ProductClient struct {
	Name         string  `json:"name"`
	ProductCode  string  `json:"product_code"`
	Measures     string  `json:"measure"`
	Stock        uint    `json:"stock"`
	UnitPrice    float32 `json:"unit_price"`
	Uri          string  `json:"uri"`
	Offer        bool    `json:"offer"`
	Picture      string  `json:"picture"`
	Description  string  `json:"description"`
}

