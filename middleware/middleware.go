package middleware

import "net/http"

func Authentication(fun func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request){
		token := r.Header.Get("Authorization")
		if token == ""{
			forbidden(w, r)
		}
		fun(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"message_type":""error", "message":"NO TIENES AUTORIZACION"}`))
}