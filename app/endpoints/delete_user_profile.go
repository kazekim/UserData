/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"github.com/kazekim/UserData/internal/udconst"
	"github.com/kazekim/UserData/pkg/udhttp"
	"net/http"
)

func (ep *defaultEndpoint) deleteUserProfile(w http.ResponseWriter, r *http.Request) {

	id, err := ep.parseUserId(r)
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, udconst.ErrMessageIDIntegerOnly)
		return
	}
	err = ep.userUseCase.DeleteUser(id)
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	_ = udhttp.ResponseJSON(w, http.StatusOK, udconst.ErrMessageSuccess)
}
