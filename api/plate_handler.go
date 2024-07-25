package api

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

func (s *APIServer) handleIsPlateAvailable(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	plate := r.FormValue("plate")
	//plate := r.PathValue("plate")
	fmt.Println(plate)

	if !isValidColombianPlate(plate) || len(plate) == 0 {
		return WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Formato invalido de placa"})
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

func isValidColombianPlate(plate string) bool {
	plate = strings.ToUpper(strings.TrimSpace(plate))

	// Patrones de placas colombianas
	patterns := map[string]*regexp.Regexp{
		"Vehículos particulares":            regexp.MustCompile(`^[A-Z]{3}\d{3}$`),
		"Vehículos públicos":                regexp.MustCompile(`^[A-Z]{3}\d{3}$`),
		"Motos":                             regexp.MustCompile(`^[A-Z]{3}\d{2}[A-Z]$`),
		"Remolques y semirremolques":        regexp.MustCompile(`^R\d{5}$`),
		"Vehículos diplomáticos":            regexp.MustCompile(`^(CD|CC|AT)\d{4}$`),
		"Vehículos de internación temporal": regexp.MustCompile(`^[A-Z]{2}\d{4}$`),
	}

	for _, pattern := range patterns {
		if pattern.MatchString(plate) {
			return true
		}
	}

	return false
}
