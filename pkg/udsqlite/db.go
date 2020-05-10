/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package udsqlite

import "database/sql"

func NewDB(name string) (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "./"+name+".db")
	return database, err
}
