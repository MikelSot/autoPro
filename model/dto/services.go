package dto

// para intercambiar info (y ademas aqui directo le pasamos el id del taller y ya en la query lo separamos)
type ServiceWorkshops []*ServiceWorkshop
type ServiceWorkshop struct {
	ID          uint8  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Uri         string `json:"uri"`
	WorkshopID  uint8  `json:"workshop_id"`
}

// mostrar info de workshop
type WorkshopClients []*WorkshopClient
type WorkshopClient struct {
	Name    string `json:"name"`
	Active  bool   `json:"active"`
	Address string `json:"address"`
}
