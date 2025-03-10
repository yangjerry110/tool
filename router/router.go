package router

func InitHttpRouterEnginee() router {
	return SetRouterEnginee(&http{})
}

func RegisterHttp(registerName string, routerRegister routerRegister) RouterRegisterHttp {
	routerEnginee().register(registerName, routerRegister)
	return routerRegister.routerRegisterHttp()
}
