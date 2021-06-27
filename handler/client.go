package handler

import (
	"encoding/json"
	"github.com/MikelSot/autoPro/model"
	"net/http"
)

type client struct {
	query IQueryExists
	crud IClientCRUD
}

func newClient(exists IQueryExists, crud IClientCRUD) client {
	return client{exists,crud}
}

func (c *client) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
	}

	data := model.Client{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
	}

	err = c.crud.Create(&data)
	if err != nil {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
	}

	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type":""message", "message":"correcto"}`))

}