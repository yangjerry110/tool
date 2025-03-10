package router

var defaultRouter router

func SetRouterEnginee(routerEnginee router) router {
	defaultRouter = routerEnginee
	return defaultRouter
}

func routerEnginee() router {
	return defaultRouter
}
