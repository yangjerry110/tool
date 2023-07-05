/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-25 15:45:12
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 21:35:32
 * @Description: base
 */
package service

/**
 * @description: CreateCommonService
 * @param {...CommonService} CommonServices
 * @author: Jerry.Yang
 * @date: 2023-05-25 16:24:08
 * @return {*}
 */
func CreateCommonService(CommonServices ...CommonService) CommonService {
	if len(CommonServices) == 0 {
		return &Common{}
	}
	return CommonServices[0]
}

/**
 * @description: CreateParseService
 * @param {...ParseService} ParseServices
 * @author: Jerry.Yang
 * @date: 2023-05-25 21:35:38
 * @return {*}
 */
func CreateParseService(ParseServices ...ParseService) ParseService {
	if len(ParseServices) == 0 {
		return &Parse{}
	}
	return ParseServices[0]
}

/**
 * @description: CreateReviseInterfaceService
 * @param {...ReviseInterfaceService} ReviseInterfaceServices
 * @author: Jerry.Yang
 * @date: 2023-05-25 19:23:14
 * @return {*}
 */
func CreateReviseInterfaceService(ReviseInterfaceServices ...ReviseInterfaceService) ReviseInterfaceService {
	if len(ReviseInterfaceServices) == 0 {
		return &ReviseInterface{}
	}
	return ReviseInterfaceServices[0]
}
