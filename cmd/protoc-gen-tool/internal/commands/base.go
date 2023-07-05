/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-23 16:03:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-24 14:32:11
 * @Description: base
 */
package commands

/**
 * @description: CreateGenerateCommands
 * @param {...GenerateCommands} GenerateCommands
 * @author: Jerry.Yang
 * @date: 2023-05-23 16:04:32
 * @return {*}
 */
func CreateGenerateCommands(GenerateCommands ...GenerateCommands) GenerateCommands {
	if len(GenerateCommands) == 0 {
		return &Generate{}
	}
	return GenerateCommands[0]
}

/**
 * @description: CreateFileCommands
 * @param {...FileCommands} FileCommands
 * @author: Jerry.Yang
 * @date: 2023-05-23 17:31:48
 * @return {*}
 */
func CreateFileCommands(FileCommands ...FileCommands) FileCommands {
	if len(FileCommands) == 0 {
		return &File{}
	}
	return FileCommands[0]
}

/**
 * @description: CreateRouterCommands
 * @param {...RouterCommands} RouterCommands
 * @author: Jerry.Yang
 * @date: 2023-05-23 19:10:52
 * @return {*}
 */
func CreateRouterCommands(RouterCommands ...RouterCommands) RouterCommands {
	if len(RouterCommands) == 0 {
		return &Router{}
	}
	return RouterCommands[0]
}

/**
 * @description: CreateServiceCommands
 * @param {...ServiceCommands} ServiceCommands
 * @author: Jerry.Yang
 * @date: 2023-05-24 14:20:29
 * @return {*}
 */
func CreateServiceCommands(ServiceCommands ...ServiceCommands) ServiceCommands {
	if len(ServiceCommands) == 0 {
		return &Service{}
	}
	return ServiceCommands[0]
}

/**
 * @description: CreateHttpCommands
 * @param {...HttpCommands} HttpCommands
 * @author: Jerry.Yang
 * @date: 2023-05-24 14:32:16
 * @return {*}
 */
func CreateHttpCommands(HttpCommands ...HttpCommands) HttpCommands {
	if len(HttpCommands) == 0 {
		return &Http{}
	}
	return HttpCommands[0]
}
