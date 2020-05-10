/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"github.com/kazekim/UserData/pkg/udhttp"
	"net/http"
)

func (ep *defaultEndpoint) listUsers(w http.ResponseWriter, r *http.Request) {

	us, err := ep.userUseCase.ListUser()
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	_ = udhttp.ResponseJSON(w, http.StatusOK, *us)

}
