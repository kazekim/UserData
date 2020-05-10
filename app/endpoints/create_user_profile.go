/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"github.com/kazekim/UserData/internal/udconst"
	"github.com/kazekim/UserData/pkg/udhttp"
	"net/http"
)

type CreateUserRequest struct {
	Email     string `json:"email"`
	Fullname  string `json:"full_name"`
	Telephone string `json:"telephone"`
}

func (ep *defaultEndpoint) createUserProfile(w http.ResponseWriter, r *http.Request) {

	var request CreateUserRequest
	err := ep.parseJSONBodyParam(r, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ep.userUseCase.CreateUser(request.Email, request.Fullname, request.Telephone)
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	_ = udhttp.ResponseJSON(w, http.StatusOK, udconst.ErrMessageSuccess)

}
