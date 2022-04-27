package system

import "association/service"

var (
	registerService  = service.ServiceGroupApp.SystemServiceGroup.RegisterService
	userLoginService = service.ServiceGroupApp.SystemServiceGroup.UserLoginService
	homePageService  = service.ServiceGroupApp.SystemServiceGroup.HomePageService
)
