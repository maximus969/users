package httpserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/maximus969/users-app/internal/app/domain"
	"github.com/maximus969/users-app/internal/app/transport/httpserver/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHttpServer_GetUser(t *testing.T) {
	userServiceMock := mocks.NewUserService(t)

	createdAt := time.Now()
	testCreatedUser, err := domain.NewUser(domain.NewUserData{
		Id:        uuid.MustParse("b123991c-a5b5-413f-8f78-070d2ae5f481"),
		Firstname: "Test",
		Lastname:  "Lastname",
		Email:     "test@gmail.com",
		Age:       30,
		Created:   createdAt,
	})
	require.NoError(t, err)

	userServiceMock.On("CreateUser", mock.Anything, mock.Anything).Return(testCreatedUser, nil)

	httpServer := NewHttpServer(userServiceMock)

	newUserRequest := []byte(`{
		"id": "b123991c-a5b5-413f-8f78-070d2ae5f481",
		"firstname": "Test",
		"lastname": "Lastname",
		"email": "test@gmail.com",
		"age": 30
	}`)

	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(newUserRequest))
	w := httptest.NewRecorder()

	httpServer.CreateUser(w, req)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)

	var createUserResponse UserResponse
	err = json.NewDecoder(res.Body).Decode(&createUserResponse)
	require.NoError(t, err)

	require.Equal(t, createUserResponse.Id.String(), testCreatedUser.Id().String())
	require.Equal(t, createUserResponse.Firstname, testCreatedUser.Firstname())
	require.Equal(t, createUserResponse.Lastname, testCreatedUser.Lastname())
	require.Equal(t, createUserResponse.Email, testCreatedUser.Email())
	require.Equal(t, createUserResponse.Age, testCreatedUser.Age())
}

func TestHttpServer_DeleteUser(t *testing.T) {
	userServiceMock := mocks.NewUserService(t)

	testUUID := uuid.MustParse("b123991c-a5b5-413f-8f78-070d2ae5f481")

	userServiceMock.On("GetUserById", mock.Anything, testUUID).Return(domain.User{}, nil)
	userServiceMock.On("DeleteUser", mock.Anything, testUUID).Return(nil)

	httpServer := NewHttpServer(userServiceMock)

	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", httpServer.DeleteUser).Methods(http.MethodDelete)

	req := httptest.NewRequest(http.MethodDelete, "/user/"+testUUID.String(), nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)

	var deleteResponse map[string]bool
	err := json.NewDecoder(res.Body).Decode(&deleteResponse)
	require.NoError(t, err)

	require.Equal(t, true, deleteResponse["deleted"])
}
