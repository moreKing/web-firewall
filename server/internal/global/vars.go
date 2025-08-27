package global

import (
	"context"
	"regexp"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
)

var (
	HttpServer *ghttp.Server

	apis []ghttp.RouterItem

	paths []OpenApi

	regPath = regexp.MustCompile(`^{[A-Za-z\d-]+}$`)
)

type OpenApi struct {
	Name   string
	Path   string
	Method string
}

func init() {

	file, err := gcfg.NewAdapterFile("config.yaml")
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}

	err = file.Set("server.webTimeOut", 16)
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}
}

func GetApis() *[]ghttp.RouterItem {
	if len(apis) == 0 {

		LoginReg := regexp.MustCompile(`^/api/v\d+/login`)
		PublicReg := regexp.MustCompile(`^/api/v\d+/public/`)
		PrivateReg := regexp.MustCompile(`^/api/v\d+/`)
		items := HttpServer.GetRoutes()
		for _, item := range items {
			if !(LoginReg.MatchString(item.Route) || PublicReg.MatchString(item.Route)) && PrivateReg.MatchString(item.Route) {
				//g.Log().Debug(context.Background())
				apis = append(apis, item)
			}
		}
	}

	return &apis
}

func GetPaths() []OpenApi {
	if len(paths) == 0 {
		for s, path := range HttpServer.GetOpenApi().Paths {

			if path.Get != nil {
				paths = append(paths, OpenApi{
					Name:   path.Get.Summary,
					Path:   s,
					Method: "GET",
				})
			}

			if path.Put != nil {
				paths = append(paths, OpenApi{
					Name:   path.Put.Summary,
					Path:   s,
					Method: "PUT",
				})
			}

			if path.Post != nil {
				paths = append(paths, OpenApi{
					Name:   path.Post.Summary,
					Path:   s,
					Method: "POST",
				})
			}

			if path.Delete != nil {
				paths = append(paths, OpenApi{
					Name:   path.Delete.Summary,
					Path:   s,
					Method: "DELETE",
				})
			}

			if path.Patch != nil {
				paths = append(paths, OpenApi{
					Name:   path.Patch.Summary,
					Path:   s,
					Method: "PATCH",
				})
			}

		}

		g.Log().Debug(context.TODO(), "paths: ", paths)
	}

	return paths
}

func GetPathName(path, method string) string {
	// g.Log().Debug(context.TODO(), fmt.Sprintf("path: %s method: %s", path, method))
	for _, api := range GetPaths() {
		//g.Log().Debug(context.TODO(), api)
		if !strings.EqualFold(api.Method, method) {
			continue
		}
		tmpA := strings.Split(api.Path, "/")
		tmpB := strings.Split(path, "/")
		if len(tmpB) != len(tmpA) {
			continue
		}
		g.Log().Debug(context.TODO(), "api path: ", api.Path)
		for i := range tmpA {
			if tmpA[i] == tmpB[i] {
				continue
			}
			if strings.HasPrefix(tmpA[i], ":") {
				continue
			}

			goto LOOP
		}
		// 返回
		g.Log().Debug(context.TODO(), "regPath true: ", api.Name)
		return api.Name

	LOOP:
	}
	return ""
}
