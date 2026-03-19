package models

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int                 `json:"-"`
	Headers map[string][]string `json:"-"`
	Cookies []*http.Cookie      `json:"-"`
	Data    any                 `json:"data,omitempty"`
	Error   string              `json:"error,omitempty"`
	Meta    map[string]any      `json:"meta,omitempty"`
}

type JSON_Body struct {
	Data  any            `json:"data,omitempty"`
	Error string         `json:"error,omitempty"`
	Meta  map[string]any `json:"meta,omitempty"`
}

// HTTP writer wrapper function to simply responses
func (r *Response) WriteHTTP(writer http.ResponseWriter) {
	// Order of Operations: Headers -> Cookies -> Status -> Body
	for keys, vals := range r.Headers {
		for _, val := range vals {
			writer.Header().Add(keys, val)
		}
	}

	for _, cookie := range r.Cookies {
		http.SetCookie(writer, cookie)
	}

	// assume if status not initialized, 200 OK
	if r.Status == 0 {
		r.Status = http.StatusOK
	}
	writer.WriteHeader(r.Status)

	json_encoder := json.NewEncoder(writer)
	json_encoder.SetEscapeHTML(false)
	_ = json_encoder.Encode(
		JSON_Body{
			Data:  r.Data,
			Error: r.Error,
			Meta:  r.Meta,
		},
	)
}
