package users_transport_http

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	Fullname    string  `json:"fullname"`
	PhoneNumber *string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	Fullname    string  `json:"fullname"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

	}

}
