package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	s "post-example/service"
)

type IStarship interface {
	StarshipsHandler(w http.ResponseWriter, r *http.Request)
}

type Starship struct {
	service s.IStarships
}

func (h *Starship) StarshipsHandler(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.GetStarships()
	if err != nil {
		fmt.Printf("Error getting starships: %s", err.Error())
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func NewStarship() Starship {
	return Starship{
		service: s.NewStarships(),
	}
}