/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:00:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 10:56:25
 * @Description: base
 */
package dao

/**
 * @description: CreateBaseDao
 * @param {...BaseDao} BaseDaos
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:09:15
 * @return {*}
 */
func CreateBaseDao(BaseDaos ...BaseDao) BaseDao {
	if len(BaseDaos) == 0 {
		return &Base{}
	}
	return BaseDaos[0]
}

/**
 * @description: CreateCommonDao
 * @param {...CommonDao} CommonDaos
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:09:54
 * @return {*}
 */
func CreateCommonDao(CommonDaos ...CommonDao) CommonDao {
	if len(CommonDaos) == 0 {
		return &Common{}
	}
	return CommonDaos[0]
}

/**
 * @description: CreateNewDao
 * @param {...NewDao} NewDaos
 * @author: Jerry.Yang
 * @date: 2023-05-11 10:56:31
 * @return {*}
 */
func CreateNewDao(NewDaos ...NewDao) NewDao {
	if len(NewDaos) == 0 {
		return &New{}
	}
	return NewDaos[0]
}
