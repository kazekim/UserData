/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package userdbentities

const (
	TableUser               = "User"
	TableUserFieldID        = "ID"
	TableUserFieldEmail     = "Email"
	TableUserFieldFullName  = "Fullname"
	TableUserFieldTelephone = "Telephone"
)

type User struct {
	ID        int64
	Email     string
	Fullname  string
	Telephone string
}
