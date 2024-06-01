/*
 * ***********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ***********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ***********************************************************
 */

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"XiaoService/api/auth/v1"
)

type IAuthV1 interface {
	AuthRegister(ctx context.Context, req *v1.AuthRegisterReq) (res *v1.AuthRegisterRes, err error)
}
