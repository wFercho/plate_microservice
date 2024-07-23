package api

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func (s *APIServer) handleIsPlateAvailable(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	plate := r.FormValue("plate")
	//plate := r.PathValue("plate")

	fmt.Println(plate)
	if len(plate) == 0 {
		return fmt.Errorf("provide a validate plate")
	}

	car, err := s.store.GetCarByPlate(plate)

	if err != nil {
		return err
	}

	//tmpl, err := template.ParseFiles(filepath.Join("templates", "resultado.html"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return WriteJSON(w, http.StatusOK, nil)
	// }
	return WriteJSON(w, http.StatusOK, car)
}

func (s *APIServer) homePage(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return WriteJSON(w, http.StatusOK, nil)
	}
	tmpl.Execute(w, nil)

	return nil
}

func (s *APIServer) resultadoPage(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles(filepath.Join("templates", "resultado.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return WriteJSON(w, http.StatusOK, nil)
	}
	tmpl.Execute(w, nil)

	return nil
}
