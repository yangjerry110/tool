/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-16 14:35:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 14:39:57
 * @Description: errors
 */
package errors

import "github.com/yangjerry110/tool/toolerrors"

var (
	ErrNoEndpoints          = toolerrors.New("no endpoints provided")
	ErrNoClientPools        = toolerrors.New("no client pools available")
	ErrNoHealthyClientPools = toolerrors.New("no healthy client pools available")
	ErrInitClientPoolFailed = toolerrors.New("failed to initialize client pool")
	ErrCheckHealthFailed    = toolerrors.New("health check failed")
)
