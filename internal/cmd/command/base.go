/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 15:50:41
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-11 17:44:47
 * @Description: base command
 */
package command

type CommandInterface interface {
	New() error
}

/**
 * @description: CreateCommand
 * @param {...CommandInterface} CommandInterfaces
 * @author: Jerry.Yang
 * @date: 2023-12-11 15:39:26
 * @return {*}
 */
func CreateCommand(CommandInterfaces ...CommandInterface) CommandInterface {
	if len(CommandInterfaces) == 0 {
		return &Command{}
	}
	return CommandInterfaces[0]
}
