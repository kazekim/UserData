/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"github.com/kazekim/UserData/internal/udconst"
	"github.com/kazekim/UserData/pkg/udhttp"
	"github.com/kazekim/UserData/pkg/utils"
	"net/http"
)

func (ep *defaultEndpoint) getUserProfile(w http.ResponseWriter, r *http.Request) {

	userIdStr := udhttp.GetRequestURIParam(r, 1)
	id, err := utils.StringToInt64(userIdStr)
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, udconst.ErrMessageIDIntegerOnly)
		return
	}
	u, err := ep.userUseCase.GetUser(id)
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	_ = udhttp.ResponseJSON(w, http.StatusOK, *u)
}
