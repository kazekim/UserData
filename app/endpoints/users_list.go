/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"github.com/kazekim/UserData/app/entities"
	"github.com/kazekim/UserData/pkg/udhttp"
	"net/http"
)

type ListUsersResponse struct {
	Users []entities.User `json:"users"`
}

func (ep *defaultEndpoint) listUsers(w http.ResponseWriter, r *http.Request) {

	us, err := ep.userUseCase.ListUser()
	if err != nil {
		_ = udhttp.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var response ListUsersResponse
	response.Users = *us

	_ = udhttp.ResponseJSON(w, http.StatusOK, response)

}
