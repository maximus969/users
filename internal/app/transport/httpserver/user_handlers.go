package httpserver

import (
	"encoding/json"
	"errors"	
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/maximus969/users-app/internal/app/common/server"
	"github.com/maximus969/users-app/internal/app/domain"
)

// GetBook returns a book by ID
func (h HttpServer) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr, ok := vars["id"]
	if !ok {
		server.BadRequest("invalid-user-id", errors.New("no id"), w, r)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		server.BadRequest("invalid-user-id", errors.New("id not in correct format"), w, r)
		return
	}

	user, err := h.userService.GetUserById(r.Context(), userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseUser(user)

	server.RespondOK(response, w, r)
}

// CreateBook creates a new book
func (h HttpServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}

	if err := userRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	user, err := toDomainUser(userRequest)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	insertedUser, err := h.userService.CreateUser(r.Context(), user)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseUser(insertedUser)

	server.RespondOK(response, w, r)
}

// UpdateBook updates a book by ID
func (h HttpServer) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr, ok := vars["id"]
	if !ok {
		server.BadRequest("invalid-user-id", errors.New("no id"), w, r)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		server.BadRequest("invalid-user-id", errors.New("id not in correct format"), w, r)
		return
	}

	var userRequest UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}

	if err := userRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	_, err = h.userService.GetUserById(r.Context(), userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	user, err := domain.NewUser(domain.NewUserData{
		Id:        userID,
		Firstname: userRequest.Firstname,
		Lastname:  userRequest.Lastname,
		Email:     userRequest.Email,
		Age:       userRequest.Age,
	})
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	updatedUser, err := h.userService.UpdateUser(r.Context(), user)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseUser(updatedUser)

	server.RespondOK(response, w, r)
}

// DeleteBook deletes a book by ID
func (h HttpServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr, ok := vars["id"]
	if !ok {
		server.BadRequest("invalid-user-id", errors.New("no id"), w, r)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		server.BadRequest("invalid-user-id", errors.New("id not in correct format"), w, r)
		return
	}

	_, err = h.userService.GetUserById(r.Context(), userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	err = h.userService.DeleteUser(r.Context(), userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("user-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	server.RespondOK(map[string]bool{"deleted": true}, w, r)
}

func (h HttpServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	// filter by category IDs
	users, err := h.userService.GetUsers(r.Context())
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := make([]UserResponse, 0, len(users))
	for _, user := range users {
		response = append(response, toResponseUser(user))
	}

	server.RespondOK(response, w, r)
}
