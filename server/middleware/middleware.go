package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/models"
	"github.com/kachamaka/chaosgo/status"
	"github.com/kachamaka/chaosgo/tokens"
)

// Auth is a function that acts as a middleware for all handlers except login and register
// and is used to ensure that an authorized user is making requests
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")

		token := r.Header.Get("Authorization")
		claims, err := tokens.DecryptToken(token)

		if err != nil {
			encoder.Encode(models.BasicResponse{Success: false, Message: "error decrypting token", Status: status.TOKEN_ERROR})
			log.Println("middleware auth err: ", err)
			return
		}

		id, ok := claims["_id"]
		if !ok {
			encoder.Encode(models.BasicResponse{Success: false, Message: "something wrong with JWT", Status: status.TOKEN_ERROR})
			log.Println("middleware auth err: no _id in token claims")
			return
		}

		ctx := context.WithValue(r.Context(), "_id", id)
		r = r.WithContext(ctx)
		next(w, r)
	})

}

// CORS is a function that enables CORS policy for localhost applications
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization")
		next.ServeHTTP(w, r)
	})
}
