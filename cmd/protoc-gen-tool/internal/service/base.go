/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-12 17:33:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 17:33:56
 * @Description:
 */
package service

type ProtocGenToolService interface {
	Generate() error
}

/**
 * @description: CreateProtoGenToolService
 * @param {ProtocGenToolService} ProtocGenToolService
 * @author: Jerry.Yang
 * @date: 2023-12-12 11:36:28
 * @return {*}
 */
func CreateProtoGenToolService(ProtocGenToolService ProtocGenToolService) ProtocGenToolService {
	return ProtocGenToolService
}
