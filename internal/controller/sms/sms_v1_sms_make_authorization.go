/*
 * ***********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ***********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ***********************************************************
 */

package sms

import (
	"XiaoService/api/sms/v1"
	"XiaoService/internal/logic/sms"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// SmsMakeAuthorization
//
// # 生成短信验证码授权码
//
// 生成短信验证码授权码，用于发送短信验证码；根据用户的机器 IP 以及 UserAgent 生成授权码；
func (c *ControllerV1) SmsMakeAuthorization(
	ctx context.Context,
	req *v1.SmsMakeAuthorizationReq,
) (res *v1.SmsMakeAuthorizationRes, err error) {
	g.Log().Noticef(ctx, "[CONTROL] 生成短信验证码授权码 [SmsMakeAuthorization]")
	err = sms.New().CheckIfAuthorizationIsAvailable(ctx)
	if err != nil {
		return nil, err
	}
	// 从缓存获取授权码
	authorizationCode, err := sms.New().GetSmsAuthorization(ctx)
	if err != nil {
		return nil, err

	}
	return &v1.SmsMakeAuthorizationRes{
		AuthorizationCode: authorizationCode.SendingUUID,
	}, nil
}
