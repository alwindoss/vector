package server

import (
	"encoding/json"
	"net/http"

	"github.com/alwindoss/vector/internal/station"
)

type Handler interface {
	CreateStation(w http.ResponseWriter, r *http.Request)
}

func NewHandler(svc station.Service) Handler {
	return &vectorHandler{
		StnSvc: svc,
	}
}

type vectorHandler struct {
	StnSvc station.Service
}

// CreateStation implements Handler
func (h *vectorHandler) CreateStation(w http.ResponseWriter, r *http.Request) {
	stnReq := createStationRequest{}
	json.NewDecoder(r.Body).Decode(&stnReq)

	stn := &station.Station{
		Name: stnReq.Name,
	}
	resp := &createStationResponse{}
	stn, err := h.StnSvc.CreateStation(stn)
	if err != nil {
		errO := &ErrObj{
			ErrCode:    "ERR50001",
			ErrMessage: "unable to create station",
		}
		resp.Err = errO
		return
	}
	resp.ID = stn.ID
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&resp)
}
