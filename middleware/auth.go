package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

type AuthMiddleware struct {
	verifier *oidc.IDTokenVerifier
}

func NewAuthMiddleware(provider *oidc.Provider, clientID string) *AuthMiddleware {
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	return &AuthMiddleware{
		verifier: provider.Verifier(oidcConfig),
	}
}

func (am *AuthMiddleware) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No authorization header", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := am.verifier.Verify(r.Context(), bearerToken)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
