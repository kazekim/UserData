/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package udhttp

import (
	"net/http"
	"regexp"
	"strings"
)

type ServeMux interface {
	GET(pattern string, handler func(w http.ResponseWriter, r *http.Request))
	POST(pattern string, handler func(w http.ResponseWriter, r *http.Request))
	PUT(pattern string, handler func(w http.ResponseWriter, r *http.Request))
	DELETE(pattern string, handler func(w http.ResponseWriter, r *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type defaultServeMux struct {
	mux         *http.ServeMux
	methodFuncs []httpMethodFunc
}

type httpMethodFunc struct {
	path   string
	Get    []subPathHandler
	Post   []subPathHandler
	Put    []subPathHandler
	Delete []subPathHandler
}

type subPathHandler struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
}

func NewServeMux() ServeMux {

	var methodFuncs []httpMethodFunc
	mux := http.NewServeMux()
	return &defaultServeMux{
		mux:         mux,
		methodFuncs: methodFuncs,
	}
}

func (m *defaultServeMux) GET(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {

	mPtrn, sPtrn := m.splitPattern(pattern)
	sph := subPathHandler{
		sPtrn,
		handler,
	}
	for i := range m.methodFuncs {
		if m.methodFuncs[i].path == mPtrn {
			s := append(m.methodFuncs[i].Get, sph)
			m.methodFuncs[i].Get = s
			return
		}
	}

	var f httpMethodFunc
	f.path = mPtrn
	s := append(f.Get, sph)
	f.Get = s
	fs := append(m.methodFuncs, f)
	m.methodFuncs = fs
}

func (m *defaultServeMux) POST(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {

	mPtrn, sPtrn := m.splitPattern(pattern)
	sph := subPathHandler{
		sPtrn,
		handler,
	}
	for i := range m.methodFuncs {
		if m.methodFuncs[i].path == mPtrn {
			s := append(m.methodFuncs[i].Post, sph)
			m.methodFuncs[i].Post = s
			return
		}
	}

	var f httpMethodFunc
	f.path = mPtrn
	s := append(f.Post, sph)
	f.Post = s
	funcs := append(m.methodFuncs, f)
	m.methodFuncs = funcs
}

func (m *defaultServeMux) PUT(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {

	mPtrn, sPtrn := m.splitPattern(pattern)
	sph := subPathHandler{
		sPtrn,
		handler,
	}
	for i := range m.methodFuncs {
		if m.methodFuncs[i].path == mPtrn {
			s := append(m.methodFuncs[i].Put, sph)
			m.methodFuncs[i].Put = s
			return
		}
	}

	var f httpMethodFunc
	f.path = mPtrn
	s := append(f.Put, sph)
	f.Put = s
	funcs := append(m.methodFuncs, f)
	m.methodFuncs = funcs
}

func (m *defaultServeMux) DELETE(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {

	mPtrn, sPtrn := m.splitPattern(pattern)
	sph := subPathHandler{
		sPtrn,
		handler,
	}
	for i := range m.methodFuncs {
		if m.methodFuncs[i].path == mPtrn {
			s := append(m.methodFuncs[i].Delete, sph)
			m.methodFuncs[i].Delete = s
			return
		}
	}

	var f httpMethodFunc
	f.path = mPtrn
	s := append(f.Delete, sph)
	f.Delete = s
	funcs := append(m.methodFuncs, f)
	m.methodFuncs = funcs
}

func (m *defaultServeMux) splitPattern(pattern string) (string, string) {
	var re = regexp.MustCompile(`{[a-zA-Z0-9_]+}`)
	matches := re.FindStringSubmatch(pattern)
	if len(matches) > 0 {
		v := matches[0]
		ptrns := strings.Split(pattern, "/")

		idx := 0
		for i := range ptrns {
			if ptrns[i] == v {
				idx = i
				break
			}
		}
		mPtrn := strings.Join(ptrns[:idx], "/") + "/"
		sPtrn := strings.Join(ptrns[idx:], "/")
		return mPtrn, sPtrn
	}
	return pattern, "/"
}

func (m *defaultServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	pathReq := r.RequestURI
	switch r.Method {
	case http.MethodGet:
		for i := range m.methodFuncs {
			if strings.HasPrefix(pathReq, m.methodFuncs[i].path) {
				f := m.methodFuncs[i]
				sPtrn := strings.Replace(pathReq, f.path, "", 1)
				sps := strings.Split(sPtrn, "/")
				for j := range f.Get {
					ps := strings.Split(f.Get[j].path, "/")
					if len(ps) == len(sps) {
						f.Get[j].handler(w, r)
						return
					} else if sPtrn == "" && f.Get[j].path == "/" {
						f.Get[j].handler(w, r)
						return
					}
				}

			}
		}
	case http.MethodPost:
		for i := range m.methodFuncs {
			if strings.HasPrefix(pathReq, m.methodFuncs[i].path) {
				f := m.methodFuncs[i]
				sPtrn := strings.Replace(pathReq, f.path, "", 1)
				sps := strings.Split(sPtrn, "/")
				for j := range f.Post {
					ps := strings.Split(f.Post[j].path, "/")
					if len(ps) == len(sps) {
						f.Post[j].handler(w, r)
						return
					} else if sPtrn == "" && f.Post[j].path == "/" {
						f.Post[j].handler(w, r)
						return
					}
				}
			}
		}
	case http.MethodPut:
		for i := range m.methodFuncs {
			if strings.HasPrefix(pathReq, m.methodFuncs[i].path) {
				f := m.methodFuncs[i]
				sPtrn := strings.Replace(pathReq, f.path, "", 1)
				sps := strings.Split(sPtrn, "/")
				for j := range f.Put {
					ps := strings.Split(f.Put[j].path, "/")
					if len(ps) == len(sps) {
						f.Put[j].handler(w, r)
						return
					} else if sPtrn == "" && f.Put[j].path == "/" {
						f.Put[j].handler(w, r)
						return
					}
				}
			}
		}
	case http.MethodDelete:
		for i := range m.methodFuncs {
			if strings.HasPrefix(pathReq, m.methodFuncs[i].path) {
				f := m.methodFuncs[i]
				sPtrn := strings.Replace(pathReq, f.path, "", 1)
				sps := strings.Split(sPtrn, "/")
				for j := range f.Delete {
					ps := strings.Split(f.Delete[j].path, "/")
					if len(ps) == len(sps) {
						f.Delete[j].handler(w, r)
						return
					} else if sPtrn == "" && f.Delete[j].path == "/" {
						f.Delete[j].handler(w, r)
						return
					}
				}
			}
		}
	}
	noImplementHandler(w, r)
}

func noImplementHandler(w http.ResponseWriter, r *http.Request) {
	_ = ResponseJSON(w, http.StatusNotFound, "Page Not Found")
}
