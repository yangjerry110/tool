/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 16:47:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-12 16:47:56
 * @Description: template
 */
package template

type TemplateInterface interface {
	New() error
}

/**
 * @description: CreateTemplate
 * @param {TemplateInterface} TemplateInterface
 * @author: Jerry.Yang
 * @date: 2023-12-12 16:48:39
 * @return {*}
 */
func CreateTemplate(TemplateInterface TemplateInterface) TemplateInterface {
	return TemplateInterface
}
