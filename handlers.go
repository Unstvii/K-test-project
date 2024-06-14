package main

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func calculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	req, ok := ctx.Value("request").(Request)
	if !ok {
		http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
		return
	}

	resultAChan := make(chan *big.Int)
	resultBChan := make(chan *big.Int)

	go func() {
		resultAChan <- factorial(*req.A)
	}()
	go func() {
		resultBChan <- factorial(*req.B)
	}()

	resultA := <-resultAChan
	resultB := <-resultBChan

	response := Response{
		FactorialA: resultA.String(),
		FactorialB: resultB.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
