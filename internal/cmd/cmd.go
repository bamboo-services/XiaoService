/*
 * ***********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ***********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ***********************************************************
 */

package cmd

import (
	"XiaoService/internal/config/middleware"
	"XiaoService/internal/controller/auth"
	"XiaoService/internal/controller/sms"
	"context"
	"github.com/bamboo-services/bamboo-utils/bmiddle"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(bmiddle.BambooMiddleHandler)
				group.Middleware(middleware.MiddleRequestHandler)
				group.Bind(
					auth.NewV1(),
					sms.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
