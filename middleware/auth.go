package middleware

import (
	"context"
	"net/http"

	"github.com/yigithancolak/custmate/graph/model"
	"github.com/yigithancolak/custmate/postgresdb"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"organization"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(s *postgresdb.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				next.ServeHTTP(w, r)
				return
			}
			auth := r.Header.Get("Authorization")

			if auth == "" {
				next.ServeHTTP(w, r)
				return
			}

			bearer := "Bearer "
			auth = auth[len(bearer):]

			payload, err := s.JWTMaker.VerifyToken(auth)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			// get the user from the database
			organization, err := s.GetOrganizationById(payload.OrganizationID)
			if err != nil {
				http.Error(w, "organization not exists", http.StatusForbidden)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, organization)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.Organization {
	raw, _ := ctx.Value(userCtxKey).(*model.Organization)

	return raw
}
