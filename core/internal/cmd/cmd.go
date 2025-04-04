package cmd

import (
	"billionmail-core/internal/consts"
	"billionmail-core/internal/controller/dockerapi"
	"billionmail-core/internal/controller/domains"
	"billionmail-core/internal/controller/mail_boxes"
	"billionmail-core/internal/controller/overview"
	"billionmail-core/internal/controller/rbac"
	"billionmail-core/internal/service/database_initialization"
	docker "billionmail-core/internal/service/dockerapi"
	"billionmail-core/internal/service/middlewares"
	"billionmail-core/internal/service/phpfpm"
	rbac2 "billionmail-core/internal/service/rbac"
	"billionmail-core/internal/service/redis_initialization"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

var (
	Main = gcmd.Command{
		Name:  consts.DEFAULT_SERVER_NAME,
		Usage: consts.DEFAULT_SERVER_NAME,
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// Init Database
			err = database_initialization.InitDatabase()

			if err != nil {
				glog.Error(ctx, err)
				return err
			}

			// Init Redis
			err = redis_initialization.InitRedis()

			if err != nil {
				glog.Error(ctx, err)
				return err
			}

			// Connect to Docker
			dk, err := docker.NewDockerAPI()

			if err != nil {
				glog.Error(ctx, err)
				return err
			}

			defer dk.Close()

			// Create a new server instance
			s := g.Server(consts.DEFAULT_SERVER_NAME)

			s.Group("/api", func(group *ghttp.RouterGroup) {
				// Add CORS middleware
				group.Middleware(ghttp.MiddlewareCORS)

				// Add docker client middleware
				group.Middleware(func(r *ghttp.Request) {
					r.SetCtxVar(consts.DEFAULT_DOCKER_CLIENT_CTX_KEY, dk)
					r.Middleware.Next()
				})

				// Add JWT middleware
				group.Middleware(rbac2.JWT().JWTAuthMiddleware)

				// Add RBAC middleware
				// group.Middleware(middlewares.NewRBACMiddleware().PermissionCheck)

				// group.Middleware(ghttp.MiddlewareHandlerResponse)

				// Add response
				group.Middleware(middlewares.HandleApiResponse)

				group.Bind(
					rbac.NewV1(),
					domains.NewV1(),
					mail_boxes.NewV1(),
					overview.NewV1(),
					dockerapi.NewV1(),
				)
			})

			// Add PHP-FPM middleware
			s.BindMiddleware("/roundcube/*any", func(r *ghttp.Request) {
				if r.Method == "POST" && strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
					// Get and store the request body
					r.GetBody()
					r.Middleware.Next()
					return
				}
				r.Middleware.Next()
			})

			// Binding PHP-FPM handler
			s.BindHandler("/roundcube/*any", phpfpm.PHPFpmHandlerFactory(phpfpm.PHPFpmHandlerConfig{
				Network: "unix",
				Addr:    consts.PHP_FPM_SOCK_PATH,
				Root:    consts.ROUNDCUBE_ROOT_PATH_IN_CONTAINER,
				Static:  consts.ROUNDCUBE_ROOT_PATH,
			}))

			// Proxy unix socket for ACME challenge
			s.BindHandler("/.well-known/acme-challenge/*any", func(r *ghttp.Request) {
				// Set the backend URL to the Unix socket
				socketPath := "/tmp/acme-challenge.sock"
				backendURL, err := url.Parse(r.GetSchema() + "://unix" + socketPath)
				if err != nil {
					glog.Error(r.Context(), "Error parsing backend URL:", err)
					r.Response.WriteStatus(http.StatusInternalServerError)
					return
				}

				// Set up the dialer to connect to the Unix socket
				dialer := &net.Dialer{
					Timeout: 5 * time.Second,
				}

				proxy := httputil.NewSingleHostReverseProxy(backendURL)
				proxy.Transport = &http.Transport{
					// TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
						return dialer.Dial("unix", socketPath)
					},
				}

				// Save the original host header
				originalHost := r.Host
				r.Header.Set("X-Forwarded-Host", originalHost)

				// Forward the request to the backend
				proxy.ServeHTTP(r.Response.Writer, r.Request)
			})

			s.Run()
			return nil
		},
	}
)
