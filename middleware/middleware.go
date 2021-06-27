package middleware

import (
	"net/http"
	"github.com/MikelSot/autoPro/jwt"
)

func Authentication(fun func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){
		token := r.Header.Get("Authorization")
		_, err := jwt.ValidateToken(token)
		if err != nil {
			forbidden(w, r)
			return
		}
		fun(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"message_type":""error", "message":"NO TIENES AUTORIZACION"}`))
}