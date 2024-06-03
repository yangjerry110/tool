/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 16:24:23
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-05-31 14:41:02
 * @Description: yaml conf
 */
package conf

import (
	"fmt"
	"io/ioutil"

	"github.com/yangjerry110/tool/internal/errors"
	"github.com/yangjerry110/tool/internal/toolErrors"
	"gopkg.in/yaml.v3"
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
		return toolErrors.NewError(errors.ErrYamlConfIsNoFilePath)
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
	// Using viper，viper is not ok;
	// Use fastNotify monitor file

	// Read file content
	fileContent, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", y.FilePath, y.FileName))
	if err != nil {
		return err
	}

	// Yaml Unmarshal
	if err := yaml.Unmarshal(fileContent, y.ConfData); err != nil {
		return err
	}

	// Watch file
	if err := CreateConf(&Watch{WatchFile: &WatchFile{FilePath: y.FilePath, FileName: y.FileName, FileType: y.FileType, ConfData: y.ConfData}}).SetConfig(); err != nil {
		return err
	}
	return nil
}
