/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package udhttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

func ResponseJSON(w http.ResponseWriter, code int, i interface{}) error {

	response := Response{
		Code: code,
		Data: i,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}
