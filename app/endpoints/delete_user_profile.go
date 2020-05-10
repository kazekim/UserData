/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"github.com/kazekim/UserData/pkg/udhttp"
	"net/http"
)

func deleteUserProfile(w http.ResponseWriter, r *http.Request) {

	userId := udhttp.GetRequestURIParam(r, 1)
	_ = udhttp.ResponseJSON(w, http.StatusOK, userId)
}
