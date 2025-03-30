package api

type GuessRequest struct {
	Guess string `json:"guess"`
}

const (
	GuessRoute = "/games/{" + GameID + "}"
)
