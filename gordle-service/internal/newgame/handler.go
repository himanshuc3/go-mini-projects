package newgame

import (
	"encoding/json"
	"gordle-service/internal/api"
	"gordle-service/internal/session"
	"log"
	"net/http"
)

// NOTE:
// 1. curl -v -X POST localhost:8080 => verbose, method
// 2. Return status codes: 200 - ok, 201 - created
func Handle(w http.ResponseWriter, req *http.Request) {
	game, err := createGame()

	if err != nil {
		log.Printf("unable to create a new game: %s", err)
		http.Error(w, "failed to create a new game", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	apiGame := response(game)
	err = json.NewEncoder(w).Encode(apiGame)

	if err != nil {
		log.Printf("failed to writer response: %s", err)
	}

	// w.WriteHeader(http.StatusCreated)
	// _, _ = w.Write([]byte("Creating a new game"))
}

func createGame() (session.Game, error) {
	return session.Game{}, nil
}

func response(game session.Game) api.GameResponse {
	return api.GameResponse{}
}
