package web

import (
	"api_music/server/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MusicHandlers struct {
	service *service.MusicService
}


func NewMusicHandler(service *service.MusicService) *MusicHandlers{
	return &MusicHandlers{service: service}
	
}

func (h *MusicHandlers) GetMusic(w http.ResponseWriter, r *http.Request) {
	music, err := h.service.GetMusic()

	if err != nil {
		http.Error(w, "Falha ao pegar a música: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(music)
}


func (h *MusicHandlers) SetMusic(w http.ResponseWriter, r *http.Request) {
	var music service.Music

	if err := json.NewDecoder(r.Body).Decode(&music); err != nil {
		http.Error(w, "Carga de solicitação inválida: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.SetMusic(&music); err != nil {
		http.Error(w, "Falha na Criação da música: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(music)
}

func (h *MusicHandlers) GetMusicByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idStr := vars["id"]

	id, err :=  strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID Inválido: ", http.StatusBadRequest)
		return
	}

	music, err := h.service.GetMusicByID(id)

	if err != nil {
		http.Error(w, "Falha em pegar a música: ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(music)

}

func (h *MusicHandlers) UpdateMusic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID Inválido", http.StatusBadRequest)
		return
	}

	var music service.Music

	if err := json.NewDecoder(r.Body).Decode(&music); err != nil {
		http.Error(w, "Carga de solicitação inválida", http.StatusBadRequest)
		return
	}

	music.ID = id

	if err := h.service.UpdateMusic(&music); err != nil {
		http.Error(w, "Falha na edição da música", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(music)
}


func (h *MusicHandlers) DeleteMusic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID Inválido", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteMusic(id); err != nil {
		http.Error(w, fmt.Sprintf("Falha na exclusão da música: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}