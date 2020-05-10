/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package userdb

import (
	userdbentities "github.com/kazekim/UserData/internal/userdb/v1/entities"
	"time"
)

func (c defaultClient) List() (*[]userdbentities.User, error) {
	rows, _ := c.db.Query("SELECT " + userdbentities.TableUserFieldID + ", " + userdbentities.TableUserFieldEmail + ", " + userdbentities.TableUserFieldFullName +
		", " + userdbentities.TableUserFieldTelephone + ", " + userdbentities.TableUserFieldCreatedAt +
		", " + userdbentities.TableUserFieldUpdatedAt + " FROM " + userdbentities.TableUser)
	var id int64
	var email string
	var fullName string
	var telephone string
	var createdAt time.Time
	var updatedAt time.Time

	var us []userdbentities.User
	for rows.Next() {
		err := rows.Scan(&id, &email, &fullName, &telephone, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		u := userdbentities.User{
			id,
			email,
			fullName,
			telephone,
			createdAt,
			updatedAt,
		}
		us = append(us, u)
	}
	return &us, nil
}
