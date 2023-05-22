/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:49:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 16:00:45
 * @Description: baseVo
 */
package vo

/**
 * @description: CreateDemoInputVo
 * @param {...DemoInputVo} DemoInputVos
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:59:53
 * @return {*}
 */
func CreateDemoInputVo(DemoInputVos ...DemoInputVo) DemoInputVo {
	if len(DemoInputVos) == 0 {
		return &DemoInput{}
	}
	return DemoInputVos[0]
}

/**
 * @description: CreateDemoOutputVo
 * @param {...DemoOutputVo} DemoOutputVos
 * @author: Jerry.Yang
 * @date: 2023-05-18 16:00:51
 * @return {*}
 */
func CreateDemoOutputVo(DemoOutputVos ...DemoOutputVo) DemoOutputVo {
	if len(DemoOutputVos) == 0 {
		return &DemoOutput{}
	}
	return DemoOutputVos[0]
}

/**
 * @description: CreateNewVo
 * @param {...NewVo} NewVos
 * @author: Jerry.Yang
 * @date: 2023-05-08 15:11:16
 * @return {*}
 */
func CreateNewVo(NewVos ...NewVo) NewVo {
	if len(NewVos) == 0 {
		return &New{}
	}
	return NewVos[0]
}
