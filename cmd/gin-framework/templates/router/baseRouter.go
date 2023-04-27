/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 15:03:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 15:59:04
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
 * @description: CreateTestRouter
 * @param {...TestRouter} TestRouters
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:51:58
 * @return {*}
 */
func CreateTestRouter(TestRouters ...TestRouter) TestRouter {
	if len(TestRouters) == 0 {
		return &Test{}
	}
	return TestRouters[0]
}
