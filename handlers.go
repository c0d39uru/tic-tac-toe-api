package main

import (
	"encoding/json"
	"github.com/cleancode4ever/tic-tac-toe-api/internal/board"
	"github.com/cleancode4ever/tic-tac-toe-api/internal/cell"
	"github.com/cleancode4ever/tic-tac-toe-api/internal/game"
	"net/http"
)

const (
	contentTypeJson = "application/health+json"
)

func (s *Server) handlePing() http.HandlerFunc {
	code := http.StatusOK
	pong, _ := json.Marshal(struct {
		Status  string `json:"status"`
		Version string `json:"version"`
	}{
		Status:  "pass",
		Version: s.version,
	})

	return func(w http.ResponseWriter, r *http.Request) {
		setContentType(w, contentTypeJson)
		w.WriteHeader(code)
		w.Write(pong)
	}
}

func (s *Server) handlePlay() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var status int
		var result []byte
		var req struct {
			Row byte
			Col byte
			Content [board.Dimension][board.Dimension]cell.Value
		}
		json.NewDecoder(r.Body).Decode(&req)

		g := game.New()
		err := g.SetBoard(&req.Content, req.Row, req.Col)

		setContentType(w, contentTypeJson)

		if err != nil {
			status = http.StatusBadRequest
		} else {
			status = http.StatusOK
		}

		w.WriteHeader(status)
		result, _ = json.Marshal(g)
		w.Write(result)
	}
}

func setContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}
