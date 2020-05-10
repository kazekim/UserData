/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"github.com/kazekim/UserData/pkg/udhttp"
	"net/http"
)

type Endpoint interface {
	RegisterUserService()
	Handler() http.Handler
}

type defaultEndpoint struct {
	mux udhttp.ServeMux
}

func NewEndpoint() Endpoint {

	mux := udhttp.NewServeMux()

	return &defaultEndpoint{
		mux: mux,
	}
}

func (ep *defaultEndpoint) RegisterUserService() {

	ep.GET("/users/{id}", getUserProfile)
	ep.PUT("/users/", createUserProfile)
	ep.POST("/users/{id}", createUserProfile)
	ep.DELETE("/users/{id}", deleteUserProfile)

	//ep.HandleFunc("/", getRoot)
	//ep.sv.Han
}

func (ep *defaultEndpoint) GET(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	ep.mux.GET(pattern, handler)
}

func (ep *defaultEndpoint) POST(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	ep.mux.POST(pattern, handler)
}

func (ep *defaultEndpoint) PUT(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	ep.mux.PUT(pattern, handler)
}

func (ep *defaultEndpoint) DELETE(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	ep.mux.DELETE(pattern, handler)
}

func (ep *defaultEndpoint) Handler() http.Handler {
	return ep.mux
}
