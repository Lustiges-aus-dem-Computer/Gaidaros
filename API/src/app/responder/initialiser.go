package responder

import (
	"app/responder/respondtypes"
	"net/http"

	"github.com/gorilla/mux"
)

func SetHandler() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", respondtypes.Get).Methods(http.MethodGet)
	api.HandleFunc("", respondtypes.Post).Methods(http.MethodPost)
	api.HandleFunc("", respondtypes.Put).Methods(http.MethodPut)
	api.HandleFunc("", respondtypes.Delete).Methods(http.MethodDelete)

	api.HandleFunc("/user/{userID}/comment/{commentID}", respondtypes.Params).Methods(http.MethodGet)

	return r
}
