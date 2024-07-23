package keycloack

import (
	ctx "context"
	"net/http"
	"time"
)

func (k *Keycloak) KeycloakMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !k.CheckSession(r) {
			err := k.RefreshToken(w, r)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
		}

		next.ServeHTTP(w, r)
	}
}

func (k *Keycloak) SessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isLoggedIn := k.CheckSession(r)

		// Si no es la ruta de login y no está logueado, redirigir a login
		if r.URL.Path != "/login" && !isLoggedIn {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Si es la ruta de login y está logueado, redirigir a home
		if r.URL.Path == "/login" && isLoggedIn {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		// En cualquier otro caso, continuar con el siguiente handler
		next.ServeHTTP(w, r)
	}
}

// RefreshToken intenta renovar el token de acceso
func (k *Keycloak) RefreshToken(w http.ResponseWriter, r *http.Request) error {
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		return err
	}

	token, err := k.Gocloak.RefreshToken(ctx.Background(), refreshToken.Value, k.ClientId, k.ClientSecret, k.Realm)
	if err != nil {
		return err
	}

	// Establece las nuevas cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token.AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    token.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(30 * 24 * time.Hour), // 30 días, ajusta según necesites
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	return nil
}

// CheckSession verifica si la sesión actual es válida
func (k *Keycloak) CheckSession(r *http.Request) bool {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return false
	}

	_, _, err = k.Gocloak.DecodeAccessToken(ctx.Background(), cookie.Value, k.Realm)
	return err == nil
}
