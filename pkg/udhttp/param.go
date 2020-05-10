/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package udhttp

import (
	"net/http"
	"strings"
)

func GetRequestURIParam(r *http.Request, index int) string {
	pathReq := r.RequestURI
	ps := strings.Split(pathReq, "/")
	if index < len(ps) {
		//Use url.QueryEscape for pre go1.8
		value := ps[index + 1]
		return value
	}

	return ""
}

