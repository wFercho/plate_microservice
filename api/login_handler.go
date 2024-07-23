package api

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	ctx "context"
)

func (s *APIServer) loginHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("METODO:", r.Method)
	tmpl, err := template.ParseFiles(filepath.Join("templates", "login.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	if r.Method == http.MethodGet {
		return tmpl.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		usernameOrEmail := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println("DATA", s.kclk.ClientId, s.kclk.ClientSecret, s.kclk.Realm, usernameOrEmail, password)

		token, err := s.kclk.Gocloak.Login(ctx.Background(), s.kclk.ClientId, s.kclk.ClientSecret, s.kclk.Realm, usernameOrEmail, password)
		if err != nil {
			fmt.Println("ERR:", err)
			tmpl.Execute(w, map[string]string{"error": "Invalid credentials"})
			return nil
		}

		// Establecer la cookie del token de acceso
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    token.AccessToken,
			Path:     "/",
			Expires:  time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
			HttpOnly: true,
			Secure:   r.TLS != nil, // Habilitar en producción con HTTPS
			SameSite: http.SameSiteStrictMode,
		})

		// Establecer la cookie del token de actualización
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    token.RefreshToken,
			Path:     "/",
			Expires:  time.Now().Add(30 * 24 * time.Hour), // 30 días, ajusta según necesites
			HttpOnly: true,
			Secure:   r.TLS != nil, // Habilitar en producción con HTTPS
			SameSite: http.SameSiteStrictMode,
		})

		http.Redirect(w, r, "/", http.StatusFound)
		return nil
	}

	return WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method Not Allowed"})
}
