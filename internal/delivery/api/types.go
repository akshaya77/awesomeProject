package api

import v1 "awesomeProject/internal/usecase/v1"

type API struct {
	v1Api v1.ApiV1
}

type (

	Base struct {
		ServerProcessTime string `json:"server_process_time"`
		Status       string `json:"status"`
		ErrorMessage []string `json:message_error,omitempty`
	}

	Response struct {
		Base
		Data interface{} `json:"data"`
	}
)
