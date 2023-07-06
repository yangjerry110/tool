/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:02:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 15:58:02
 * @Description: base service
 */
package service

/**
 * @description: CreateBaseService
 * @param {...BaseService} BaseServices
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:12:41
 * @return {*}
 */
func CreateBaseService(BaseServices ...BaseService) BaseService {
	if len(BaseServices) == 0 {
		return &Base{}
	}
	return BaseServices[0]
}

/**
 * @description: CreateBeforStartService
 * @param {...BeforStartService} BeforStartServices
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:13:28
 * @return {*}
 */
func CreateBeforStartService(BeforStartServices ...BeforStartService) BeforStartService {
	if len(BeforStartServices) == 0 {
		return &BeforStart{}
	}
	return BeforStartServices[0]
}

/**
 * @description: CreateDemoService
 * @param {...DemoService} DemoServices
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:58:09
 * @return {*}
 */
func CreateDemoService(DemoServices ...DemoService) DemoService {
	if len(DemoServices) == 0 {
		return &Demo{}
	}
	return DemoServices[0]
}

/**
 * @description: CreateNewService
 * @param {...NewService} NewServices
 * @author: Jerry.Yang
 * @date: 2023-05-08 11:38:07
 * @return {*}
 */
func CreateNewService(NewServices ...NewService) NewService {
	if len(NewServices) == 0 {
		return &New{}
	}
	return NewServices[0]
}

/**
 * @description: CreateCommonService
 * @param {...CommonService} CommonServices
 * @author: Jerry.Yang
 * @date: 2023-05-17 18:30:11
 * @return {*}
 */
func CreateCommonService(CommonServices ...CommonService) CommonService {
	if len(CommonServices) == 0 {
		return &Common{}
	}
	return CommonServices[0]
}