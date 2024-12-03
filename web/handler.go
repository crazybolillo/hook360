package web

import (
	"bytes"
	"encoding/json"
	"github.com/crazybolillo/hook360/repo"
	"io"
	"log/slog"
	"net/http"
)

type eventPayload struct {
	ID    string `json:"id"`
	Event string `json:"event"`
}

type Handler struct {
	store *repo.Event
}

func NewHandler(store *repo.Event) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Failed to read request body", "reason", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var payload eventPayload
	decoder := json.NewDecoder(bytes.NewReader(body))
	err = decoder.Decode(&payload)
	if err != nil {
		slog.Error("Failed to decode request body", "reason", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	slog.Info("Storing event", "id", payload.ID, "event", payload.Event)
	err = h.store.Save(r.Context(), body)
	if err != nil {
		slog.Error("Failed to save event", "reason", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
