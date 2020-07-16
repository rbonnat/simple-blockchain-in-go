package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rbonnat/blockchain-in-go/service"
)

// Message contains data of Blockchain
type Message struct {
	Value int
}

// HandleGetBlockchain returns handler for GetBlockchain
func HandleGetBlockchain(s *service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.MarshalIndent(s.Blocks(r.Context()), "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		io.WriteString(w, string(bytes))
	}
}

// HandleWriteBlock returns handler for WriteBlock Route
func HandleWriteBlock(s *service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var m Message

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&m); err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, r.Body)
			return
		}
		defer r.Body.Close()

		newBlock, err := s.InsertNewBlock(r.Context(), m.Value)
		if err != nil {
			respondWithJSON(w, r, http.StatusInternalServerError, m)
			return
		}

		respondWithJSON(w, r, http.StatusCreated, *newBlock)
	}
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
