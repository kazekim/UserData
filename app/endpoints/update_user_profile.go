/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"fmt"
	"github.com/kazekim/UserData/internal/udconst"
	"github.com/kazekim/UserData/pkg/udhttp"
	"net/http"
)

type UpdateUserRequest struct {
	Email     string `json:"email"`
	Fullname  string `json:"full_name"`
	Telephone string `json:"telephone"`
}

func (ep *defaultEndpoint) updateUserProfile(w http.ResponseWriter, r *http.Request) {

	id, err := ep.parseUserId(r)
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, udconst.ErrMessageIDIntegerOnly)
		return
	}

	var request UpdateUserRequest
	err = ep.parseJSONBodyParam(r, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(request, " ", id)

	err = ep.userUseCase.UpdateUser(id, request.Email, request.Fullname, request.Telephone)
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	_ = udhttp.ResponseJSON(w, http.StatusOK, udconst.ErrMessageSuccess)
}
