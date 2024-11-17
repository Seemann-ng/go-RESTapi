package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"

	"github.com/Seemann-ng/go-RESTapi/internal/app/model"
	"github.com/Seemann-ng/go-RESTapi/internal/app/store"
)

const (
	sessionName = "APIServerSession"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

// Server ...
type Server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

// NewServer ...
func NewServer(storage store.Store, sessionStore sessions.Store) *Server {
	server := &Server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        storage,
		sessionStore: sessionStore,
	}

	server.configureRouter()

	return server
}

// ServeHTTP ...
func (server *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}

// configureRouter ...
func (server *Server) configureRouter() {
	server.router.HandleFunc("/users", server.handleUsersCreate()).Methods("POST")
	server.router.HandleFunc("/sessions", server.handleSessionCreate()).Methods("POST")
}

// handleSessionCreate ...
func (server *Server) handleSessionCreate() http.HandlerFunc {
	type requestStruct struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		newRequest := &requestStruct{}
		if err := json.NewDecoder(request.Body).Decode(newRequest); err != nil {
			server.error(writer, request, http.StatusBadRequest, err)
			return
		}

		user, err := server.store.User().FindByEmail(newRequest.Email)
		if err != nil || !user.ComparePassword(newRequest.Password) {
			server.error(writer, request, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := server.sessionStore.Get(request, sessionName)
		if err != nil {
			server.error(writer, request, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = user.ID
		if err := session.Save(request, writer); err != nil {
			server.error(writer, request, http.StatusInternalServerError, err)
			return
		}

		server.respond(writer, request, http.StatusOK, nil)
	}
}

// handleUsersCreate ...
func (server *Server) handleUsersCreate() http.HandlerFunc {
	type requestStruct struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		newRequest := &requestStruct{}
		if err := json.NewDecoder(request.Body).Decode(newRequest); err != nil {
			server.error(writer, request, http.StatusBadRequest, err)
			return
		}

		user := &model.User{
			Email:    newRequest.Email,
			Password: newRequest.Password,
		}

		if err := server.store.User().Create(user); err != nil {
			server.error(writer, request, http.StatusUnprocessableEntity, err)
			return
		}

		user.Sanitize()
		server.respond(writer, request, http.StatusCreated, user)
	}
}

// error ...
func (server *Server) error(writer http.ResponseWriter, request *http.Request, responseCode int, err error) {
	server.respond(writer, request, responseCode, map[string]string{"error": err.Error()})
}

// respond ...
func (server *Server) respond(writer http.ResponseWriter, request *http.Request, responseCode int, data interface{}) {
	writer.WriteHeader(responseCode)
	if data != nil {
		err := json.NewEncoder(writer).Encode(data)
		if err != nil {
			logrus.Error(err)
		}
	}
}
