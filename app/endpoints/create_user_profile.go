/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package endpoints

import (
	"fmt"
	"net/http"
)

func createUserProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wassap, %s!", r.URL.Path[1:])
}
