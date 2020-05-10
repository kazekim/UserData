/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package userdbentities

import "time"

const (
	TableUser               = "User"
	TableUserFieldID        = "ID"
	TableUserFieldEmail     = "Email"
	TableUserFieldFullName  = "Fullname"
	TableUserFieldTelephone = "Telephone"
	TableUserFieldCreatedAt = "CreatedAt"
	TableUserFieldUpdatedAt = "UpdatedAt"
)

type User struct {
	ID        int64
	Email     string
	Fullname  string
	Telephone string
	CreatedAt time.Time
	UpdatedAt time.Time
}
