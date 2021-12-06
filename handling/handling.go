package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type IsEvenRequest struct {
	Value int `json:"value"`
}

type IsEvenResponse struct {
	Even bool `json:"even"`
}

func IsEven(req IsEvenRequest, _ http.ResponseWriter, _ *http.Request) (*IsEvenResponse, int, error) {
	return &IsEvenResponse{Even: req.Value%2 == 0}, http.StatusOK, nil
}

func main() {
	go http.ListenAndServe(":8080", Handler(IsEven))

	var body bytes.Buffer
	json.NewEncoder(&body).Encode(map[string]interface{}{
		"value": 2,
	})

	resp, err := http.Post("http://localhost:8080/", "application/json", &body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(result)
}

type Error struct {
	Message string `json:"msg"`
}

type HandlerFunc[Req any, Res any] func(Req, http.ResponseWriter, *http.Request) (*Res, int, error)

func Handler[T, U any](h HandlerFunc[T, U]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req T

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			write(w, http.StatusBadRequest, Error{Message: err.Error()})
			return
		}

		res, code, err := h(req, w, r)
		if err != nil {
			if code == 0 {
				code = http.StatusInternalServerError
			}

			write(w, code, Error{Message: err.Error()})
			return
		}

		if code == 0 {
			code = http.StatusOK
		}

		write(w, code, res)
	})
}

func write(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
