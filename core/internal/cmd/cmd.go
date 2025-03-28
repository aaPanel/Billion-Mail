package cmd

import (
	"billionmail-core/internal/consts"
	"billionmail-core/internal/controller/dockerapi"
	"billionmail-core/internal/controller/domains"
	"billionmail-core/internal/controller/mail_boxes"
	"billionmail-core/internal/controller/overview"
	"billionmail-core/internal/service/database_initialization"
	docker "billionmail-core/internal/service/dockerapi"
	"billionmail-core/internal/service/middlewares"
	"billionmail-core/internal/service/phpfpm"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	Main = gcmd.Command{
		Name:  consts.DEFAULT_SERVER_NAME,
		Usage: consts.DEFAULT_SERVER_NAME,
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			err = database_initialization.InitDatabase()

			if err != nil {
				glog.Error(ctx, err)
				return err
			}

			dk, err := docker.NewDockerAPI()

			if err != nil {
				glog.Error(ctx, err)
				return err
			}

			defer dk.Close()

			s := g.Server(consts.DEFAULT_SERVER_NAME)

			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(func(r *ghttp.Request) {
					r.SetCtxVar(consts.DEFAULT_DOCKER_CLIENT_CTX_KEY, dk)
					r.Middleware.Next()
				})

				// group.Middleware(ghttp.MiddlewareHandlerResponse)

				group.Middleware(middlewares.HandleApiResponse)

				group.Bind(
					domains.NewV1(),
					mail_boxes.NewV1(),
					overview.NewV1(),
					dockerapi.NewV1(),
				)
			})

			// Binding PHP-FPM handler
			s.BindHandler("/roundcube/*any", phpfpm.PHPFpmHandlerFactory(phpfpm.PHPFpmHandlerConfig{
				Network: "unix",
				Addr:    consts.PHP_FPM_SOCK_PATH,
				Root:    consts.ROUNDCUBE_ROOT_PATH_IN_CONTAINER,
				Static:  consts.ROUNDCUBE_ROOT_PATH,
			}))

			s.Run()
			return nil
		},
	}
)
