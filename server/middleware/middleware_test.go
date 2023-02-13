package middleware_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kachamaka/chaosgo/database"
	"github.com/kachamaka/chaosgo/middleware"
	"github.com/spf13/viper"
)

var TOKEN string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2M2VhNmQ4ZGI4YjUwODdhNDNkNjcwZmMifQ.8XDgxvu5sQJzHCYSTPOLJPyWDBrXljzJCnRASG3pPfQ"

func BlankHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	_, err := database.GetHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func TestAuthMiddleware(t *testing.T) {
	r, err := http.NewRequest("GET", "/getEvents", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set("Authorization", TOKEN)
	viper.Set("TOKEN_SECRET", "9JbQi751WkxVStyLhQxLOA2YExjYL4MK")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware.Auth(BlankHandler))

	handler.ServeHTTP(rr, r)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
