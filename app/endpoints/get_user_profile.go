/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"github.com/kazekim/UserData/app/entities"
	"github.com/kazekim/UserData/internal/udconst"
	"github.com/kazekim/UserData/pkg/udhttp"
	"net/http"
)

type GetUserProfileResponse struct {
	User entities.User `json:"user"`
}

func (ep *defaultEndpoint) getUserProfile(w http.ResponseWriter, r *http.Request) {

	id, err := ep.parseUserId(r)
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, udconst.ErrMessageIDIntegerOnly)
		return
	}
	u, err := ep.userUseCase.GetUser(id)
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var response GetUserProfileResponse
	response.User = *u

	_ = udhttp.ResponseJSON(w, http.StatusOK, response)
}
