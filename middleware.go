package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func validateJSONMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.A == nil || req.B == nil || *req.A < 0 || *req.B < 0 {
			http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "request", req)
		r = r.WithContext(ctx)
		next(w, r, ps)
	}
}
