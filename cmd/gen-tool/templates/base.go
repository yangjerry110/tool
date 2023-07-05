/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 14:45:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 15:47:43
 * @Description: base
 */
package templates

/**
 * @description: CreateCommonTemplate
 * @param {...CommonTemplate} CommonTemplates
 * @author: Jerry.Yang
 * @date: 2023-04-24 14:51:15
 * @return {*}
 */
func CreateCommonTemplate(CommonTemplates ...CommonTemplate) CommonTemplate {
	if len(CommonTemplates) == 0 {
		return &Common{}
	}
	return CommonTemplates[0]
}

/**
 * @description: CreateMainTemplate
 * @param {...MainTemplate} MainTemplates
 * @author: Jerry.Yang
 * @date: 2023-04-26 10:39:18
 * @return {*}
 */
func CreateMainTemplate(MainTemplates ...MainTemplate) MainTemplate {
	if len(MainTemplates) == 0 {
		return &Main{}
	}
	return MainTemplates[0]
}

/**
 * @description: CreateModTemplate
 * @param {...ModTemplate} ModTemplates
 * @author: Jerry.Yang
 * @date: 2023-04-26 10:40:40
 * @return {*}
 */
func CreateModTemplate(ModTemplates ...ModTemplate) ModTemplate {
	if len(ModTemplates) == 0 {
		return &Mod{}
	}
	return ModTemplates[0]
}
