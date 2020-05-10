/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package udhttp

import (
	"net/http"
)

type httpContext struct {
	handler func(http.ResponseWriter, *http.Request)
}

func (c *httpContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.handler(w, r)
}

func GET(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	c := httpContext{
		handler: h,
	}
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			c.ServeHTTP(w, r)
		} else {
			_ = ResponseJSON(w, http.StatusMethodNotAllowed, "Method not Allow")
		}
	}
}

func POST(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	c := httpContext{
		handler: h,
	}
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			c.ServeHTTP(w, r)
		} else {
			_ = ResponseJSON(w, http.StatusMethodNotAllowed, "Method not Allow")
		}
	}
}