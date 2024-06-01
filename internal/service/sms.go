/*
 * ***********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ***********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ***********************************************************
 */

// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"XiaoService/internal/model/rdo"
	"context"
)

type (
	ISms interface {
		// CheckIfAuthorizationIsAvailable
		//
		// # 检查授权是否可用
		//
		// 检查授权是否可用，用于检查授权是否可用，根据用户的机器 IP 以及 UserAgent 生成授权码；
		// 若返回的内容没有报错的内容则说明可以继续获取授权码；
		//
		// # 参数
		//   - ctx		上下文(context.Context)
		//
		// # 返回
		//   - err		错误(error)
		CheckIfAuthorizationIsAvailable(ctx context.Context) (err error)
		// GetSmsAuthorization
		//
		// # 获取短信验证码授权码
		//
		// 获取短信验证码授权码，用于获取短信验证码授权码；
		//
		// # 参数
		//   - ctx		上下文(context.Context)
		//
		// # 返回
		GetSmsAuthorization(ctx context.Context) (authorizationCode *rdo.RedisSmsAuthorization, err error)
	}
)

var (
	localSms ISms
)

func Sms() ISms {
	if localSms == nil {
		panic("implement not found for interface ISms, forgot register?")
	}
	return localSms
}

func RegisterSms(i ISms) {
	localSms = i
}
