/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 11:22:12
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-13 14:42:34
 * @Description: protocGenTool
 */
package config

type ProtocGenTool struct {
	IsFirstCreate bool
	IsAppend      bool
}

/**
 * @description: ProtocGenToolConf
 * @author: Jerry.Yang
 * @date: 2023-12-12 11:23:52
 * @return {*}
 */
var ProtocGenToolConf = &ProtocGenTool{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-12 11:26:39
 * @return {*}
 */
func (p *ProtocGenTool) SetConfig() error {

	// set isFirstCreate to conf
	// set isAppend to conf
	// ProtocGenToolConf.IsFirstCreate = *isFirstCreate
	// ProtocGenToolConf.IsAppend = *isAppend
	ProtocGenToolConf = p
	return nil
}
