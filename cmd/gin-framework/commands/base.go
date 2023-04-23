/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 15:51:29
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-23 15:51:33
 * @Description: base
 */
package commands

/**
 * @description: CreateNewCommands
 * @param {...NewCommands} NewCommands
 * @author: Jerry.Yang
 * @date: 2023-04-23 15:52:27
 * @return {*}
 */
func CreateNewCommands(NewCommands ...NewCommands) NewCommands {
	if len(NewCommands) == 0 {
		return &New{}
	}
	return NewCommands[0]
}
