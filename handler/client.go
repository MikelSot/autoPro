package handler

import (
	"encoding/json"
	"github.com/MikelSot/autoPro/model/dto"
	"net/http"
	"strconv"
)

type client struct {
	query IQueryExists
	crud IClientCRUD
}

func NewClient(q IQueryExists, c IClientCRUD) client {
	return client{
		q,
		c,
	}
}


func (c *client) singIn(w http.ResponseWriter, r *http.Request) {
	// validar con regex
	if r.Method != http.MethodPost {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
		return
	}

	data := dto.SignInClient{}
	err := json.NewDecoder(r.Body).Decode(&data) // de json a estructura
	if err != nil {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
		return
	}

	err = c.crud.Create(&data)
	if err != nil {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
		return
	}

	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type":""message", "message":"correcto"}`))
}

func (c *client) editClient(w http.ResponseWriter, r *http.Request)  {
	// validar con regex
	if r.Method != http.MethodPut {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
		return
	}

	data := dto.EditClient{}
	err  = json.NewDecoder(r.Body).Decode(&data) // de json a estructura
	if err != nil {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
		return
	}

	err = c.crud.Update(uint(ID),&data)
	if err != nil {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type":""error", "message":"Metodo no permitido"}`))
		return
	}

	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type":""message", "message":"correcto"}`))
}

func (c *client) getById(w http.ResponseWriter, r *http.Request)  {

}

func (c *client) getAll(w http.ResponseWriter, r *http.Request)  {

}