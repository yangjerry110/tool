/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 16:24:23
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-08 17:35:54
 * @Description: yaml conf
 */
package conf

import (
	"github.com/spf13/viper"
	"github.com/yangjerry110/tool/internal/errors"
)

type Yaml struct {
	FilePath string
	FileName string
	FileType string
	ConfData interface{}
}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-08 17:35:25
 * @return {*}
 */
func (y *Yaml) SetConfig() error {

	// judge filePath
	if y.FilePath == "" {
		return errors.ErrYamlConfIsNoFilePath
	}

	// judge fileName
	if y.FileName == "" {
		return errors.ErrYamlConfIsNoFileName
	}

	// judge confData
	if y.ConfData == nil {
		return errors.ErrYamlConfIsNoConfData
	}

	// hot reload yaml conf
	// (1) Use Signal to update, but Signal, only user connection and user behavior, is not sure whether it is feasible, thinking
	// (2) Use the last update time of the file, in this case, you need to load a conf file to add a goroutine, is the cost a little too big, thinking
	// (3) Using viper, hhhhh, look at the implementation of viper, the number of coroutines and so on are controlled, and the writing is OK
	viper.AddConfigPath(y.FilePath)
	viper.SetConfigName(y.FileName)
	viper.SetConfigType(y.FileType)

	// judge is have err
	// return err
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// watch config
	viper.WatchConfig()

	// set decode config
	viper.Unmarshal(&y.ConfData)
	return nil
}
