package global

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
	"slices"
)

//var routes = &[]model.Route{}

var routes []string

func init() {
	//getwd, err := os.Getwd()
	//if err != nil {
	//	g.Log().Fatal(context.TODO(), "getwd err", err)
	//	return
	//}
	//routesPath := path.Join(getwd, consts.RoutesPath)
	//g.Log().Debug(context.TODO(), "routesPath ", routesPath)
	//
	////	 读取routes
	//file, err := os.ReadFile(routesPath)
	//if err != nil {
	//	g.Log().Fatal(context.Background(), "读routesPath文件失败", err)
	//	return
	//}
	//err = json.Unmarshal(file, routes)
	//if err != nil {
	//	g.Log().Fatal(context.Background(), "读routesPath文件失败", err)
	//}
	//
	//routes = removeRoutesConstant(*routes)
	//
	//routes = sortSelect(*routes)

	conf := g.Cfg().MustGet(context.Background(), "server.menuList")
	routes = conf.Strings()

}

func removeRoutesConstant(routes []model.Route) *[]model.Route {
	var newRoutes = []model.Route{}
	for _, route := range routes {
		if route.Meta.Constant {
			continue
		}
		newRoute := model.Route{
			Name:      route.Name,
			Path:      route.Path,
			Component: route.Component,
			Meta:      route.Meta,
		}
		if route.Children != nil && len(*route.Children) > 0 {
			newRoute.Children = removeRoutesConstant(*route.Children)
		}
		newRoutes = append(newRoutes, newRoute)
	}

	return &newRoutes
}

func GetAllRoutes() []string {
	return routes
}

// 选择排序
func sortSelect(routes []model.Route) *[]model.Route {
	for i := 0; i < len(routes); i++ {
		tmp := i
		for j := i + 1; j < len(routes); j++ {
			if routes[tmp].Meta.Order > routes[j].Meta.Order {
				tmp = j
			}
		}
		if tmp != i {
			routes[tmp], routes[i] = routes[i], routes[tmp]
		}
		if routes[i].Children != nil && len(*routes[i].Children) > 0 {
			routes[i].Children = sortSelect(*routes[i].Children)
		}
	}

	return &routes
}

//func FilterRoutes(menus []string, r *[]model.Route) *[]model.Route {
//	var newRoutes []model.Route
//	for _, route := range *r {
//		if route.Children != nil && len(*route.Children) > 0 {
//			children := FilterRoutes(menus, route.Children)
//			if children != nil && len(*children) > 0 {
//				newRoutes = append(newRoutes, model.Route{
//					Name:      route.Name,
//					Path:      route.Path,
//					Component: route.Component,
//					Meta:      route.Meta,
//					Children:  children,
//				})
//			}
//		} else {
//			if slices.Contains(menus, route.Name) {
//				newRoutes = append(newRoutes, model.Route{
//					Name:      route.Name,
//					Path:      route.Path,
//					Component: route.Component,
//					Meta:      route.Meta,
//				})
//			}
//		}
//	}
//	return &newRoutes
//}

func FilterRoutes(menus []string) []string {
	var newRoutes []string
	for _, route := range routes {
		if slices.Contains(menus, route) {
			newRoutes = append(newRoutes, route)
		}
	}

	return newRoutes
}
