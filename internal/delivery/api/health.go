package api

import "net/http"

func (api *API) Health(rw http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "Alive", nil
}
