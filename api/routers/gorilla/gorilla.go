package gorilla

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golvl2/api/handlers"
	"io"
	"net/http"
)

type Router struct {
	*mux.Router
	h *handlers.Handlers
}

func NewRouter(handlers *handlers.Handlers) *Router {
	r := &Router{
		Router: mux.NewRouter(),
		h:      handlers,
	}

	r.HandleFunc("/create", r.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", r.GetUser).Methods(http.MethodGet)

	return r
}

func (r *Router) CreateUser(w http.ResponseWriter, req *http.Request) {
	u := handlers.User{}
	in, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(in, &u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := r.h.CreateUser(req.Context(), u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (r *Router) GetUser(w http.ResponseWriter, req *http.Request) {

	v := mux.Vars(req)

	resp, err := r.h.GetUser(req.Context(), v["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
