package guess

import (
	"encoding/json"
	"gordle-service/internal/api"
	"log"
	"net/http"
)

func Handle(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue(api.GameID)
	log.Printf("id: %v", id)
	if id == "" {
		http.Error(writer, "missing the id of the game", http.StatusBadRequest)
		return
	}

	// NOTE:
	// 1. log is not thread safe but then do we require our logs
	// in a particulate ordering
	log.Printf("retrieve status of game with id: %v", id)
	r := api.GuessRequest{}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	// ...encode into JSON
	err := json.NewDecoder(request.Body).Decode(&r)

	if err != nil {
		http.Error(writer, "problem encoding the object", http.StatusInternalServerError)
		return
	}
	log.Printf("Guess is: %v", r)

}
