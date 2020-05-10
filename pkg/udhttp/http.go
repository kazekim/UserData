/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package udhttp

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	sv *http.Server
	port string
}

func NewServer(port string) *Server {
	s := &http.Server{
		Addr:           ":"+port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &Server{
		sv: s,
		port: port,
	}
}

func (s *Server) RegisterHandler(handler http.Handler) {
	s.sv.Handler = handler
}

func (s *Server) ListenAndServe() {
	log.Fatal(s.sv.ListenAndServe())
}


