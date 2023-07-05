/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 16:40:07
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 15:56:41
 * @Description: baseController
 */
package controller

/**
 * @description: CreateBaseController
 * @param {...BaseController} BaseControllers
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:44:28
 * @return {*}
 */
func CreateBaseController(BaseControllers ...BaseController) BaseController {
	if len(BaseControllers) == 0 {
		return &Base{}
	}
	return BaseControllers[0]
}

/**
 * @description: CreateDemoController
 * @param {...DemoController} DemoControllers
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:56:46
 * @return {*}
 */
func CreateDemoController(DemoControllers ...DemoController) DemoController {
	if len(DemoControllers) == 0 {
		return &Demo{}
	}
	return DemoControllers[0]
}

/**
 * @description: CreateNewController
 * @param {...NewController} NewControllers
 * @author: Jerry.Yang
 * @date: 2023-05-06 11:29:58
 * @return {*}
 */
func CreateNewController(NewControllers ...NewController) NewController {
	if len(NewControllers) == 0 {
		return &New{}
	}
	return NewControllers[0]
}
