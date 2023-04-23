/*
 * @Author: Jerry.Yang
 * @Date: 2023-02-09 14:49:25
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-02-09 15:20:33
 * @Description: base db
 */
package db

type BaseDbInterface interface{}

type BaseDb struct {
	Dsn string `yaml:"dsn"`
}

var DbConfig = map[string]*BaseDb{}
