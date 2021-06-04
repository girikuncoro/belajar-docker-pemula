package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/girikuncoro/belajar-docker-pemula/backend/db"
	"github.com/girikuncoro/belajar-docker-pemula/backend/schema"
	"github.com/girikuncoro/belajar-docker-pemula/backend/service"
)

type todoHandler struct {
	postgres *db.Postgres
	static   *db.Static
}

func (h *todoHandler) GetStatic(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepo(r.Context(), h.static)

	todoList, err := service.GetAll(ctx)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseOK(w, todoList)
}

func (h *todoHandler) getAllTodo(w http.ResponseWriter, r *http.Request) {
	if h.postgres == nil {
		responseError(w, http.StatusInternalServerError, "must connect to postgres")
		return
	}
	ctx := db.SetRepo(r.Context(), h.postgres)

	todoList, err := service.GetAll(ctx)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseOK(w, todoList)
}

func (h *todoHandler) insertTodo(w http.ResponseWriter, r *http.Request) {
	if h.postgres == nil {
		responseError(w, http.StatusInternalServerError, "must connect to postgres")
		return
	}
	ctx := db.SetRepo(r.Context(), h.postgres)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todo schema.Todo
	if err := json.Unmarshal(b, &todo); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := service.Insert(ctx, &todo)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOK(w, id)
}

func (h *todoHandler) updateTodo(w http.ResponseWriter, r *http.Request) {
	if h.postgres == nil {
		responseError(w, http.StatusInternalServerError, "must connect to postgres")
		return
	}
	ctx := db.SetRepo(r.Context(), h.postgres)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todo schema.Todo
	if err := json.Unmarshal(b, &todo); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = service.Update(ctx, &todo)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOK(w, todo.ID)
}

func (h *todoHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	if h.postgres == nil {
		responseError(w, http.StatusInternalServerError, "must connect to postgres")
		return
	}
	ctx := db.SetRepo(r.Context(), h.postgres)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var req struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(b, &req); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := service.Delete(ctx, req.ID); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func responseOK(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
