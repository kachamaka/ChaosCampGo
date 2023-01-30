package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/models"
	"github.com/kachamaka/chaosgo/tokens"
)

func Auth(next func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "\t")

		token := r.Header.Get("Authorization")
		claims, err := tokens.DecryptToken(token)

		if err != nil {
			encoder.Encode(models.BasicResponse{Success: false, Message: "error decrypting token"})
			log.Println("middleware auth err: ", err)
			return
		}

		id, ok := claims["_id"]
		if !ok {
			encoder.Encode(models.BasicResponse{Success: false, Message: "something wrong with JWT"})
			log.Println("middleware auth err: no _id in token claims")
			return
		}

		// context.Set(r, "_id", v)
		ctx := context.WithValue(r.Context(), "_id", id)
		r = r.WithContext(ctx)
		next(w, r)
	})

}