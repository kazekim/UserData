/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package migrations

import (
	"database/sql"
	"github.com/kazekim/UserData/internal/userdb/v1"
)

func AutoMigrateDB(db *sql.DB) {
	userdb.Prepare(db)
}
