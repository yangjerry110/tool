/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:02:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 16:14:04
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
 * @description: CreateTestService
 * @param {...TestService} TestServices
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:14:10
 * @return {*}
 */
func CreateTestService(TestServices ...TestService) TestService {
	if len(TestServices) == 0 {
		return &Test{}
	}
	return TestServices[0]
}
