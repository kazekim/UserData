/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/kazekim/UserData/app"
	"github.com/kazekim/UserData/internal/migrations"
	"github.com/kazekim/UserData/internal/udconst"
	"github.com/kazekim/UserData/pkg/env"
	"github.com/kazekim/UserData/pkg/udhttp"
	"github.com/kazekim/UserData/pkg/udsqlite"
)

var (
	restfulPort = env.GetValue(udconst.EnvRestfulPort, "8080")
	dbName      = env.GetValue(udconst.EnvDBName, "user")
)

func main() {

	db, err := udsqlite.NewDB(dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations.AutoMigrateDB(db)

	sv := udhttp.NewServer(restfulPort)
	cfg := app.NewConfig(restfulPort, dbName)

	s := app.NewCoreService(sv, db, cfg)
	s.RegisterService()
	s.StartHttpService()

	fmt.Println("Start HttpServer: http://localhost:" + restfulPort)
	sv.ListenAndServe()
}
