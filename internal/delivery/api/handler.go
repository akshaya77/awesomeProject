package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HandlerFunc func(rw http.ResponseWriter, r *http.Request) (interface{}, error)


func (fn HandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := Response{}

	response.Base.Status = "OK"

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "GET,PUT,OPTIONS")

	if r.Method == "OPTIONS" {
		rw.WriteHeader(http.StatusOK)
		return
	}

	start := time.Now()

	var data interface{}
	var err error
	var buf []byte

	data,err = fn(rw, r)

	response.Base.ServerProcessTime = time.Since(start).String()

	if data != nil && err == nil {
		response.Data = data
	} else if err != nil {
		response.Base.ErrorMessage = []string {err.Error()}
	}

	if buf,err = json.Marshal(response); err == nil {
		_,err := rw.Write(buf)
		if err != nil {
			log.Println("error in unmarshalling")
		}
	}

	return



}