/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 15:40:55
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:41:27
 * @Description: floder
 */
package folder

type FloderInterface interface {
	New() error
}

/**
 * @description: CreateFlod
 * @param {FloderInterface} FloderInterface
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:41:27
 * @return {*}
 */
func CreateFlod(FloderInterface FloderInterface) FloderInterface {
	return FloderInterface
}
