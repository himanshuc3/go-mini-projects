package handlers

import (
	"gordle-service/internal/api"
	"gordle-service/internal/handlers/getstatus"
	"gordle-service/internal/handlers/guess"
	"gordle-service/internal/newgame"
	"net/http"
)

// NOTE:
// 1. Handler - controller combining api routes
// and handlers from newgame
func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(api.NewGameRoute, newgame.Handle)
	return mux
}

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handle)
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handle)
	r.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handle)
	return r
}
