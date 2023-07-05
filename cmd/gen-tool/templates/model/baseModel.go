/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-18 11:34:00
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 11:38:15
 * @Description: baseModel
 */
package model

/**
 * @description: CreateBaseModel
 * @param {...BaseModel} BaseModels
 * @author: Jerry.Yang
 * @date: 2023-05-18 11:38:21
 * @return {*}
 */
func CreateBaseModel(BaseModels ...BaseModel) BaseModel {
	if len(BaseModels) == 0 {
		return &Base{}
	}
	return BaseModels[0]
}
