/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 15:03:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 15:55:09
 * @Description: baseRouter
 */
package router

/**
 * @description: CreateBaseRouter
 * @param {...BaseRouter} BaseRouters
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:50:16
 * @return {*}
 */
func CreateBaseRouter(BaseRouters ...BaseRouter) BaseRouter {
	if len(BaseRouters) == 0 {
		return &Base{}
	}
	return BaseRouters[0]
}

/**
 * @description: CreateCommonRouter
 * @param {...CommonRouter} CommonRouters
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:50:54
 * @return {*}
 */
func CreateCommonRouter(CommonRouters ...CommonRouter) CommonRouter {
	if len(CommonRouters) == 0 {
		return &Common{}
	}
	return CommonRouters[0]
}

/**
 * @description: CreateDemoRouter
 * @param {...DemoRouter} DemoRouters
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:55:15
 * @return {*}
 */
func CreateDemoRouter(DemoRouters ...DemoRouter) DemoRouter {
	if len(DemoRouters) == 0 {
		return &Demo{}
	}
	return DemoRouters[0]
}

/**
 * @description: CreateNewRouter
 * @param {...NewRouter} NewRouters
 * @author: Jerry.Yang
 * @date: 2023-05-16 17:19:16
 * @return {*}
 */
func CreateNewRouter(NewRouters ...NewRouter) NewRouter {
	if len(NewRouters) == 0 {
		return &New{}
	}
	return NewRouters[0]
}
