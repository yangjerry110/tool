/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 14:45:34
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 22:47:13
 * @Description: common
 */
package templates

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/yangjerry110/tool/cmd/gen-tool/errors"
)

type CommonTemplate interface {
	GetTemplate(names ...string) *template.Template
	SaveTemplate(path string, name string, content string, data interface{}, fileType ...string) error
	AppendTemplate(path string, content string, data interface{}, fileType ...string) error
	FirstUpper(s string) string
	GetFormatNowTime() string
}

type Common struct{}

/**
 * @description: Template
 * @author: Jerry.Yang
 * @date: 2023-04-24 14:47:27
 * @return {*}
 */
var Template *template.Template

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 14:47:52
 * @return {*}
 */
func (c *Common) GetTemplate(names ...string) *template.Template {

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
	return template.New(name)

	/**
	 * @step
	 * @判断template是否已经初始化过了
	 **/
	if Template != nil {
		return Template
	}

	/**
	 * @step
	 * @假如没有初始化过，则初始化一下
	 **/
	Template = template.New(name)
	return Template
}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} name
 * @param {string} content
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:34:48
 * @return {*}
 */
func (c *Common) SaveTemplate(path string, name string, content string, data interface{}, fileType ...string) error {
	/**
	 * @step
	 * @判断path
	 **/
	if path == "" {
		return errors.ErrTemplateSavePathIsEmpty
	}

	/**
	 * @step
	 * @获取template
	 **/
	template := c.GetTemplate()

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
 * @date: 2023-05-11 15:12:26
 * @return {*}
 */
func (c *Common) AppendTemplate(path string, content string, data interface{}, fileType ...string) error {

	/**
	 * @step
	 * @判断path
	 **/
	if path == "" {
		return errors.ErrTemplateSavePathIsEmpty
	}

	/**
	 * @step
	 * @获取template
	 **/
	template := c.GetTemplate(path)

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
 * @date: 2023-05-09 16:54:10
 * @return {*}
 */
func (c *Common) FirstUpper(s string) string {
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
func (c *Common) GetFormatNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}