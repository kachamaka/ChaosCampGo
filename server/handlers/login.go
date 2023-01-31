package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kachamaka/chaosgo/models"
	"github.com/kachamaka/chaosgo/tokens"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	var req models.LoginRequest

	if err := decoder.Decode(&req); err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "can't read body"})
		log.Println("err by decoding body: ", err)
		return
	}

	ID, err := req.MatchPassword()
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: err.Error()})
		log.Println(err)
		return
	}

	token, err := tokens.GenerateToken(ID.Hex())
	if err != nil {
		encoder.Encode(models.BasicResponse{Success: false, Message: "couldn't generate token"})
		log.Println("err with generating the token")
		return
	}

	encoder.Encode(models.AuthResponse{Token: token, BasicResponse: models.BasicResponse{Success: true, Message: "all good"}})
}
