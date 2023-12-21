package config

import (
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/internal/errors"
)

type Dao struct {
	CliContext *cli.Context
	DaoName    string
}

/**
 * @description: DaoConf
 * @author: Jerry.Yang
 * @date: 2023-12-21 15:28:37
 * @return {*}
 */
var DaoConf = &Dao{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-21 15:29:38
 * @return {*}
 */
func (d *Dao) SetConfig() error {

	// If CliContext == nil
	if d.CliContext == nil {
		return errors.ErrCmdCommandNoCliContext
	}

	// get first args
	// this first arg is daoName
	daoName := d.CliContext.Args().First()

	// judge protobufName
	// if is empty return err
	if daoName == "" {
		return errors.ErrCmdConfNoProtobufName
	}

	// set daoName
	DaoConf.DaoName = daoName
	return nil

}
