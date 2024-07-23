package api

import (
	"net/http"
	"time"
)

func (s *APIServer) logoutHandler(w http.ResponseWriter, r *http.Request) error {
	// Eliminar la cookie de auth_token
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   r.TLS != nil,
		SameSite: http.SameSiteStrictMode,
	})

	// Eliminar la cookie de refresh_token
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   r.TLS != nil,
		SameSite: http.SameSiteStrictMode,
	})

	// Redirigir al usuario a la página de login
	http.Redirect(w, r, "/login", http.StatusFound)
	return nil
}
