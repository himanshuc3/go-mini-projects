package api

import "gordle-service/internal/session"

// NOTE:
// 1. Test UT coverage: go test -cover
func ToGameResponse(g session.Game) GameResponse {
	apiGame := GameResponse{
		ID:           string(g.ID),
		AttemptsLeft: g.AttemptsLeft,
		Guesses:      make([]Guess, len(g.Guesses)),
		Status:       string(g.Status),
	}

	for index := 0; index < len(g.Guesses); index++ {
		apiGame.Guesses[index].Word = g.Guesses[index].Word
		apiGame.Guesses[index].Feedback = g.Guesses[index].Feedback
	}

	if g.AttemptsLeft == 0 {
		apiGame.Solution = ""
	}
	return apiGame
}
