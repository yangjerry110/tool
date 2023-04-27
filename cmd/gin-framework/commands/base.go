/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 15:51:29
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-26 10:48:33
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

/**
 * @description: CreateInitCommands
 * @param {...InitCommands} InitCommands
 * @author: Jerry.Yang
 * @date: 2023-04-23 16:53:04
 * @return {*}
 */
func CreateInitCommands(InitCommands ...InitCommands) InitCommands {
	if len(InitCommands) == 0 {
		return &Init{}
	}
	return InitCommands[0]
}

/**
 * @description: CreateNewConfigCommands
 * @param {...NewConfigCommands} NewConfigCommands
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:05:28
 * @return {*}
 */
func CreateNewConfigCommands(NewConfigCommands ...NewConfigCommands) NewConfigCommands {
	if len(NewConfigCommands) == 0 {
		return &NewConfig{}
	}
	return NewConfigCommands[0]
}

/**
 * @description: CreateNewControllerCommands
 * @param {...NewControllerCommands} NewControllerCommands
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:50:13
 * @return {*}
 */
func CreateNewControllerCommands(NewControllerCommands ...NewControllerCommands) NewControllerCommands {
	if len(NewControllerCommands) == 0 {
		return &NewController{}
	}
	return NewControllerCommands[0]
}

/**
 * @description: CreateNewDaoCommands
 * @param {...NewDaoCommands} NewDaoCommands
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:13:06
 * @return {*}
 */
func CreateNewDaoCommands(NewDaoCommands ...NewDaoCommands) NewDaoCommands {
	if len(NewDaoCommands) == 0 {
		return &NewDao{}
	}
	return NewDaoCommands[0]
}

/**
 * @description: CreateNewErrorCommands
 * @param {...NewErrorCommands} NewErrorCommands
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:25:25
 * @return {*}
 */
func CreateNewErrorCommands(NewErrorCommands ...NewErrorCommands) NewErrorCommands {
	if len(NewErrorCommands) == 0 {
		return &NewError{}
	}
	return NewErrorCommands[0]
}

/**
 * @description: CreateNewLoggerCommands
 * @param {...NewLoggerCommands} NewLoggerCommands
 * @author: Jerry.Yang
 * @date: 2023-04-25 10:42:41
 * @return {*}
 */
func CreateNewLoggerCommands(NewLoggerCommands ...NewLoggerCommands) NewLoggerCommands {
	if len(NewLoggerCommands) == 0 {
		return &NewLogger{}
	}
	return NewLoggerCommands[0]
}

/**
 * @description: CreateNewModelCommands
 * @param {...NewModelCommands} NewModelCommands
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:01:43
 * @return {*}
 */
func CreateNewModelCommands(NewModelCommands ...NewModelCommands) NewModelCommands {
	if len(NewModelCommands) == 0 {
		return &NewModel{}
	}
	return NewModelCommands[0]
}

/**
 * @description: CreateNewRouterCommands
 * @param {...NewRouterCommands} NewRouterCommands
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:01:02
 * @return {*}
 */
func CreateNewRouterCommands(NewRouterCommands ...NewRouterCommands) NewRouterCommands {
	if len(NewRouterCommands) == 0 {
		return &NewRouter{}
	}
	return NewRouterCommands[0]
}

/**
 * @description: CreateNewServiceCommands
 * @param {...NewServiceCommands} NewServiceCommands
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:19:09
 * @return {*}
 */
func CreateNewServiceCommands(NewServiceCommands ...NewServiceCommands) NewServiceCommands {
	if len(NewServiceCommands) == 0 {
		return &NewService{}
	}
	return NewServiceCommands[0]
}

/**
 * @description: CreateNewVoCommands
 * @param {...NewVoCommands} NewVoCommands
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:33:49
 * @return {*}
 */
func CreateNewVoCommands(NewVoCommands ...NewVoCommands) NewVoCommands {
	if len(NewVoCommands) == 0 {
		return &NewVo{}
	}
	return NewVoCommands[0]
}

/**
 * @description: CreateNewTemplateCommands
 * @param {...NewTemplateCommands} NewTemplateCommands
 * @author: Jerry.Yang
 * @date: 2023-04-26 10:48:45
 * @return {*}
 */
func CreateNewTemplateCommands(NewTemplateCommands ...NewTemplateCommands) NewTemplateCommands {
	if len(NewTemplateCommands) == 0 {
		return &NewTemplate{}
	}
	return NewTemplateCommands[0]
}
