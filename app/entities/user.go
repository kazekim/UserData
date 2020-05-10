/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package entities

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"full_name"`
	Telephone string `json:"telephone"`
}
