/*
 * @Author: Jerry.Yang
 * @Date: 2024-12-06 16:29:28
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-12-06 16:32:39
 * @Description: protoc service
 */
package config

import (
	"github.com/yangjerry110/protoc-gen-go/compiler/protogen"
	"github.com/yangjerry110/tool/internal/errors"
)

type ProtocService struct {
	ProtocService *protogen.Service
}

/**
 * @description: ProtocServiceConfs
 * @author: Jerry.Yang
 * @date: 2024-12-06 16:32:55
 * @return {*}
 */
var ProtocServiceConfs = []*ProtocService{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2024-12-06 16:32:48
 * @return {*}
 */
func (p *ProtocService) SetConfig() error {

	/**
	 * @step
	 * @判断protocService
	 **/
	if p.ProtocService == nil {
		return errors.ErrProtocGenToolServiceNoService
	}

	/**
	 * @step
	 * @set conf
	 **/
	ProtocServiceConfs = append(ProtocServiceConfs, p)
	return nil
}
