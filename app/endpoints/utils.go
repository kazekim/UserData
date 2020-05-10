/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"encoding/json"
	"github.com/kazekim/UserData/pkg/udhttp"
	"github.com/kazekim/UserData/pkg/utils"
	"net/http"
)

func (ep *defaultEndpoint) parseUserId(r *http.Request) (int64, error) {
	userIdStr := udhttp.GetRequestURIParam(r, 1)
	id, err := utils.StringToInt64(userIdStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ep *defaultEndpoint) parseJSONBodyParam(r *http.Request, req interface{}) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return err
	}
	return nil
}
