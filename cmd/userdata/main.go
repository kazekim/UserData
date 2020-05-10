/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/kazekim/UserData/app"
	"github.com/kazekim/UserData/internal/udconst"
	"github.com/kazekim/UserData/pkg/env"
	"github.com/kazekim/UserData/pkg/udhttp"
)

var (
	restfulPort = env.GetValue(udconst.EnvRestfulPort, "8080")
)

func main() {

	sv := udhttp.NewServer(restfulPort)
	cfg := app.NewConfig()

	s := app.NewCoreService(sv, cfg)
	s.RegisterService()
	s.StartHttpService()

	fmt.Println("Start HttpServer: http://localhost:" + restfulPort)
	sv.ListenAndServe()
}
