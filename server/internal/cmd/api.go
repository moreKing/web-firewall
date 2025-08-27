package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"server/internal/consts"
	"server/internal/controller/audit"
	"server/internal/controller/home"
	"server/internal/controller/login"
	"server/internal/controller/policy"
	"server/internal/controller/public"
	"server/internal/controller/route"
	"server/internal/controller/system"
	"server/internal/global"
	"server/internal/service"
	"server/internal/vars"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func register() func(ctx context.Context, parser *gcmd.Parser) (err error) {
	// 不考虑使用命令行参数，仅提供http服务
	return func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := g.Server()

		// 注册api文档
		s.BindHandler("/api/docs.css", func(r *ghttp.Request) {
			r.Response.Header().Set("content-type", "text/css")
			file, err := os.ReadFile(g.Cfg().MustGet(ctx, "server.swaggerCss").String())
			if err != nil {
				g.Log().Error(ctx, err)
				return
			}
			r.Response.Write(file)
		})
		s.BindHandler("/api/docs.js", func(r *ghttp.Request) {
			r.Response.Header().Set("content-type", "application/javascript")
			file, err := os.ReadFile(g.Cfg().MustGet(ctx, "server.swaggerJs").String())
			if err != nil {
				g.Log().Error(ctx, err)
				return
			}
			r.Response.Write(file)

		})
		s.Group("/api", func(group *ghttp.RouterGroup) {
			// 注册统一响应编码
			group.Middleware(ghttp.MiddlewareHandlerResponse)

			// v1 版本 虽然应该不会有v2 版本
			group.Group("/v1", func(v1 *ghttp.RouterGroup) {
				v1.Bind(
					login.NewV1(),
				)

				v1.Group("/", func(auth *ghttp.RouterGroup) {
					auth.Middleware(service.Middleware().Auth) // 登录校验 没有登录不允许请求接口

					//注册登录后的公共接口
					auth.Group("/public", func(pub *ghttp.RouterGroup) {
						pub.Bind(
							public.NewV1(),
						)
					})

					// 暂时不使用rbac管理

					auth.Group("/", func(private *ghttp.RouterGroup) {
						private.Middleware(service.Middleware().ConfigLog)

						// firewall 策略管理
						private.Group("/policy", func(u *ghttp.RouterGroup) {
							u.Bind(
								policy.NewV1(),
							)
						})

						private.Group("/route", func(u *ghttp.RouterGroup) {
							u.Middleware(service.Middleware().IsForward)
							u.Bind(
								route.NewV1(),
							)
						})

						// 操作日志
						private.Group("/audit", func(u *ghttp.RouterGroup) {
							u.Bind(
								audit.NewV1(),
							)
						})

						// 系统管理
						private.Group("/system", func(sys *ghttp.RouterGroup) {
							sys.Bind(
								system.NewV1(),
								home.NewV1(),
							)
						})
					})

				})
			})

		})
		s.SetSessionIdName("SessionID")
		s.SetSwaggerUITemplate(consts.MySwaggerUITemplate)

		// 想知道注册了哪些接口
		//g.Log().Debug(ctx, s.GetRoutes())
		global.HttpServer = s
		//s.SetAddr("0.0.0.0")

		redirectPort := g.Cfg().MustGet(ctx, "server.https", false).Bool()
		// 开启https
		if redirectPort {
			//go redirectToHttps()  //代理80跳转，为了避免跟NGINX等服务发生冲突，此处选择不进行代理
			s.EnableHTTPS(vars.GetCrtFile(), vars.GetKeyFile())
		}

		s.Run()
		return nil
	}
}

/*
*
http 80 重定向到https 配置文件中的https端口
*/
func redirectToHttps() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 进行301永久重定向
		addr := g.Cfg().MustGet(context.Background(), "server.address", ":443").String()
		re := regexp.MustCompile(`(\d+)\s*$`)
		portByte := re.Find([]byte(addr))
		if portByte == nil {
			return
		}
		http.Redirect(w, r, fmt.Sprintf("https://%s:%s", r.Host, string(portByte)), http.StatusMovedPermanently)
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
