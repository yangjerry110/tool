/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 16:49:37
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 17:17:12
 * @Description:
 */
package template

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/errors"
	"github.com/yangjerry110/tool/conf"
)

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} name
 * @param {string} content
 * @param {interface{}} data
 * @param {...string} fileType
 * @author: Jerry.Yang
 * @date: 2023-12-12 16:52:57
 * @return {*}
 */
func SaveTemplate(path string, name string, content string, data interface{}, fileType ...string) error {

	/**
	 * @step
	 * @判断path
	 **/
	if path == "" {
		return errors.ErrTemplateNoPath
	}

	/**
	 * @step
	 * @获取template
	 **/
	template := getTemplate()

	/**
	 * @step
	 * @渲染模版
	 **/
	temp, err := template.Parse(content)
	if err != nil {
		fmt.Printf("\r\n parse %s/%s Err : %+v \r\n", path, name, err)
		return err
	}

	/**
	 * @step
	 * @open file
	 **/
	createFileName := fmt.Sprintf("%s/%s", path, name)
	fileObj, err := os.Create(createFileName)
	if err != nil {
		fmt.Printf("\r\n create %s/%s Err : %+v \r\n", path, name, err)
		return err
	}
	defer fileObj.Close()

	/**
	 * @step
	 * @传入需要渲染的数据和渲染之后的模版的出处
	 **/
	err = temp.Execute(fileObj, data)
	if err != nil {
		fmt.Printf("\r\n Execute %s/%s Err : %+v \r\n", path, name, err)
		return err
	}

	/**
	 * @step
	 * @格式化代码
	 **/
	thisFileType := "go"
	if len(fileType) > 0 {
		thisFileType = fileType[0]
	}

	if thisFileType == "go" {
		cmd := exec.Command("gofmt", "-w", createFileName)
		if err := cmd.Run(); err != nil {
			fmt.Printf("\r\n gofmt %s/%s Err : %+v \r\n", path, name, err)
			return err
		}
	}
	return nil
}

/**
 * @description: AppendTemplate
 * @param {string} path
 * @param {string} content
 * @param {interface{}} data
 * @param {...string} fileType
 * @author: Jerry.Yang
 * @date: 2023-12-12 17:28:55
 * @return {*}
 */
func AppendTemplate(path string, content string, data interface{}, fileType ...string) error {

	/**
	 * @step
	 * @判断path
	 **/
	if path == "" {
		return errors.ErrTemplateNoPath
	}

	/**
	 * @step
	 * @获取template
	 **/
	template := getTemplate(path)

	/**
	 * @step
	 * @渲染模版
	 **/
	temp, err := template.Parse(content)
	if err != nil {
		fmt.Printf("\r\n parse %s Err : %+v \r\n", path, err)
		return err
	}

	/**
	 * @step
	 * @open file
	 **/
	fileObj, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("\r\n create %s Err : %+v \r\n", path, err)
		return err
	}
	defer fileObj.Close()

	/**
	 * @step
	 * @传入需要渲染的数据和渲染之后的模版的出处
	 **/
	err = temp.Execute(fileObj, data)
	if err != nil {
		fmt.Printf("\r\n Execute %s Err : %+v \r\n", path, err)
		return err
	}

	/**
	 * @step
	 * @格式化代码
	 **/
	thisFileType := "go"
	if len(fileType) > 0 {
		thisFileType = fileType[0]
	}

	if thisFileType == "go" {
		cmd := exec.Command("gofmt", "-w", path)
		if err := cmd.Run(); err != nil {
			fmt.Printf("\r\n gofmt %s Err : %+v \r\n", path, err)
			return err
		}
	}
	return nil
}

/**
 * @description: FirstUpper
 * @param {string} s
 * @author: Jerry.Yang
 * @date: 2023-12-12 17:13:35
 * @return {*}
 */
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

/**
 * @description: GetFormatNowTime
 * @author: Jerry.Yang
 * @date: 2023-05-16 15:08:36
 * @return {*}
 */
func GetFormatNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

/**
 * @description: getTemplate
 * @param {...string} names
 * @author: Jerry.Yang
 * @date: 2023-12-12 17:01:28
 * @return {*}
 */
func getTemplate(names ...string) *template.Template {

	/**
	 * @step
	 * @name 定义
	 **/
	name := "gen-tool-template"

	/**
	 * @step
	 * @判断name
	 **/
	if len(names) > 0 {
		name = names[0]
	}

	// set conf
	if err := conf.CreateConf(&config.Template{Name: name}).SetConfig(); err != nil {
		return nil
	}

	// return conf
	return config.TemplateConfs[name]
}
