package getstatus

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
	apiGame := api.GameResponse{
		ID: id,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	// ...encode into JSON
	json.NewEncoder(writer).Encode(apiGame)

	// if err != nil {
	// 	http.Error(writer, "problem encoding the object", http.StatusInternalServerError)
	// 	return
	// }

}
