package api

import (
	"encoding/json"
	"net/http"

	"test-connect-db/pkg/logger"
	"test-connect-db/server/api/schemas/requests"
	"test-connect-db/server/api/schemas/responses"
	"test-connect-db/server/repository"
	"test-connect-db/server/repository/models"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	L    logger.Logger
	Repo *repository.Repo
}

func NewServer(r *repository.Repo) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		Repo:   r,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/index", s.healthz()).Methods("GET")
	s.HandleFunc("/users", s.createUser()).Methods("POST")
	s.HandleFunc("/users/{name}", s.getUsersByName()).Methods("GET")
}

func (s *Server) healthz() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		resp := responses.Healthz{
			Status: "OK",
		}
		if encodeErr := json.NewEncoder(w).Encode(resp); encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			mapErr := map[string]interface{}{
				"error": encodeErr.Error(),
			}
			jsonErr, err := json.Marshal(mapErr)
			if err != nil {
				s.L.Error(err, "failed Healthz")
				return
			}
			_, err = w.Write(jsonErr)
			if err != nil {
				s.L.Error(err, "failed Healthz")
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) createUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user requests.User
		if decodeErr := json.NewDecoder(r.Body).Decode(&user); decodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			mapErr := map[string]interface{}{
				"error": decodeErr.Error(),
			}
			jsonErr, err := json.Marshal(mapErr)
			if err != nil {
				s.L.Error(err, "failed create user")
				return
			}
			_, err = w.Write(jsonErr)
			if err != nil {
				s.L.Error(err, "failed create user")
				return
			}
		}
		normalized := s.normalizeUserModel(&user)
		if dbErr := s.Repo.Users.Save(normalized); dbErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			mapErr := map[string]interface{}{
				"error": dbErr.Error(),
			}
			jsonErr, err := json.Marshal(mapErr)
			if err != nil {
				s.L.Error(err, "failed create user")
				return
			}
			_, err = w.Write(jsonErr)
			if err != nil {
				s.L.Error(err, "failed create user")
				return
			}
		}

		w.Write([]byte("OK"))
		w.WriteHeader(http.StatusOK)
	}
}
func (s *Server) getUsersByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name, ok := vars["name"]
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			mapErr := map[string]interface{}{
				"error": "missing name parameter",
			}
			jsonErr, err := json.Marshal(mapErr)
			if err != nil {
				s.L.Error(err, "failed get users by name")
				return
			}
			_, err = w.Write(jsonErr)
			if err != nil {
				s.L.Error(err, "failed get users by name")
				return
			}
		}

		users, dbErr := s.Repo.Users.ManyByName(name)
		if dbErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			mapErr := map[string]interface{}{
				"error": dbErr.Error(),
			}
			jsonErr, err := json.Marshal(mapErr)
			if err != nil {
				s.L.Error(err, "failed get users by name")
				return
			}
			_, err = w.Write(jsonErr)
			if err != nil {
				s.L.Error(err, "failed get users by name")
				return
			}
		}

		resp, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			mapErr := map[string]interface{}{
				"error": dbErr.Error(),
			}
			jsonErr, err := json.Marshal(mapErr)
			if err != nil {
				s.L.Error(err, "failed get users by name")
				return
			}
			_, err = w.Write(jsonErr)
			if err != nil {
				s.L.Error(err, "failed get users by name")
				return
			}
		}

		w.Write(resp)
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) normalizeUserModel(request *requests.User) *models.Users {
	var normalized models.Users
	normalized.Name = request.Name
	normalized.Age = request.Age
	normalized.IsAdmin = request.IsAdmin

	return &normalized
}
