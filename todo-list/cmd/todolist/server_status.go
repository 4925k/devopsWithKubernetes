package main

import (
	"encoding/json"
	"hash/fnv"
	"net/http"
	"time"
)

func (app *application) hashHandler(w http.ResponseWriter, r *http.Request) {
	appHash := fnv.New32a()
	appHash.Write([]byte("app"))

	reqHash := fnv.New32a()
	reqHash.Write([]byte(time.Now().String()))

	type out struct {
		App_hash uint32 `json:"app_hash"`
		Req_hash uint32 `json:"req_hash"`
	}

	res := out{appHash.Sum32(), reqHash.Sum32()}

	js, err := json.MarshalIndent(res, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
