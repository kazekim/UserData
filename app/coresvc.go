/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package app

import (
	"database/sql"
	"github.com/kazekim/UserData/app/endpoints"
	"github.com/kazekim/UserData/app/repositories"
	"github.com/kazekim/UserData/app/usecases"
	"github.com/kazekim/UserData/internal/userdb/v1"
	"github.com/kazekim/UserData/pkg/udhttp"
)

type Config struct {
	dbName string
}

func NewConfig() *Config {
	return &Config{}
}

type CoreService interface {
	RegisterService()
	StartHttpService()
}

type defaultCoreService struct {
	sv  *udhttp.Server
	cfg *Config
	ep  endpoints.Endpoint
}

func NewCoreService(sv *udhttp.Server, db *sql.DB, cfg *Config) CoreService {

	userDBClient := userdb.NewClient(db)

	userDBRepo := repositories.NewUserDBRepository(userDBClient)

	userUseCase := usecases.NewUserUseCase(userDBRepo)

	ep := endpoints.NewEndpoint(userUseCase)

	return &defaultCoreService{
		sv:  sv,
		cfg: cfg,
		ep:  ep,
	}
}

func (cs *defaultCoreService) RegisterService() {

	cs.ep.RegisterUserService()
}

func (cs *defaultCoreService) StartHttpService() {
	cs.sv.RegisterHandler(cs.ep.Handler())
}
