/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package userdb

import (
	"database/sql"
	"github.com/kazekim/UserData/internal/userdb/v1/entities"
)

func Prepare(db *sql.DB) {
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS " + userdbentities.TableUser + " (" + userdbentities.TableUserFieldID + " INTEGER PRIMARY KEY AUTOINCREMENT, CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, " + userdbentities.TableUserFieldEmail + " VARCHAR(50), " + userdbentities.TableUserFieldFullName + " VARCHAR(100), " + userdbentities.TableUserFieldTelephone + " VARCHAR(20))")
	_, _ = statement.Exec()
}
