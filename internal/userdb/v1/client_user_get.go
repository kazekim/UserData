/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package userdb

import (
	"github.com/kazekim/UserData/internal/userdb/v1/entities"
	"github.com/kazekim/UserData/pkg/utils"
	"time"
)

func (c *defaultClient) GetUser(id int64) (*userdbentities.User, error) {

	rows, _ := c.db.Query("SELECT " + userdbentities.TableUserFieldEmail + ", " + userdbentities.TableUserFieldFullName +
		", " + userdbentities.TableUserFieldTelephone + ", " + userdbentities.TableUserFieldCreatedAt +
		", " + userdbentities.TableUserFieldUpdatedAt + " FROM " + userdbentities.TableUser + " WHERE " + userdbentities.TableUserFieldID + "=" + utils.Int64ToString(id) + "")
	var email string
	var fullName string
	var telephone string
	var createdAt time.Time
	var updatedAt time.Time
	rows.Next()
	err := rows.Scan(&email, &fullName, &telephone, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	user := userdbentities.User{
		id,
		email,
		fullName,
		telephone,
		createdAt,
		updatedAt,
	}
	return &user, nil
}
