package api

import (
	"fmt"
	"net/http"
)

func (s *APIServer) handleIsPlateAvailable(w http.ResponseWriter, r *http.Request) error {
	plate := r.PathValue("plate")

	if len(plate) == 0 {
		return fmt.Errorf("provide a validate plate")
	}

	car, err := s.store.GetCarByPlate(plate)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, car)
}
