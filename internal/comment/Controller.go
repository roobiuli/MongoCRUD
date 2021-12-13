package comment

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/roobiuli/MongoCRUD/internal/pkg/storage"
	"net/http"
	"time"
)

type Controller struct {
	Storage storage.CommentStorer
}


// Post /api/v1/comments


func(c Controller) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("POST Method required"))
		return
	}

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("GET Method required"))
		return
	}

	var req Create

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := uuid.New().String()

	if err := c.Storage.Insert(r.Context(), storage.Comment{
		UUID:      id,
		Text:      req.Text,
		CreatedAt: time.Now(),
	}) ; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
	return

}




// GET /api/v1/comments


func(c Controller) Find(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("GET Method required"))
		return
	}

	id := r.Form.Get("uuid")

	 com, err := c.Storage.Find(r.Context(), id)
	 if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	comresp := Response{
		UUID:      com.UUID,
		Text:      com.Text,
		CreatedAt: com.CreatedAt,
	}

	respbody, err := json.Marshal(&comresp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}


	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(respbody))
	return

}

func (c Controller) Delete(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodDelete{
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Delete Method required"))
		return
	}

	id :=  r.Form.Get("uuid")

	if err := c.Storage.Delete(r.Context(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
	return
}

func (c Controller) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("PATCH Method required"))
		return
	}
	var req Create
	id := r.Form.Get("uuid")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	com := storage.Comment{
		UUID:      id,
		Text:      req.Text,
	}

	if err := c.Storage.Update(r.Context(), com); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
	return
}