/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package app

import (
	"github.com/kazekim/UserData/app/endpoints"
	"github.com/kazekim/UserData/pkg/udhttp"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{
	}
}

type CoreService interface {
	RegisterService()
	StartHttpService()
}

type defaultCoreService struct {
	sv *udhttp.Server
	cfg *Config
	ep endpoints.Endpoint
}

func NewCoreService(sv *udhttp.Server, cfg *Config) CoreService {

	ep := endpoints.NewEndpoint()


	return &defaultCoreService{
		sv: sv,
		cfg: cfg,
		ep: ep,
	}
}

func (cs *defaultCoreService) RegisterService() {

	cs.ep.RegisterUserService()
}

func (cs *defaultCoreService) StartHttpService() {
	cs.sv.RegisterHandler(cs.ep.Handler())
}
