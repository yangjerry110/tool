package config

import (
	"github.com/yangjerry110/tool/internal/errors"
	"google.golang.org/protobuf/compiler/protogen"
)

type ProtocService struct {
	ProtoService *protogen.Service
	ServiceName  string
}

/**
 * @description: ProtocServiceConf
 * @author: Jerry.Yang
 * @date: 2024-08-19 16:24:37
 * @return {*}
 */
var ProtocServiceConfs = []*ProtocService{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2024-08-19 16:25:12
 * @return {*}
 */
func (p *ProtocService) SetConfig() error {

	// Judge protoService
	if p.ProtoService == nil {
		return errors.ErrProtocGenToolServiceNoServices
	}

	// set ServiceName
	p.ServiceName = p.ProtoService.GoName
	// set ProtocServiceConfs
	ProtocServiceConfs = append(ProtocServiceConfs, p)
	return nil
}
