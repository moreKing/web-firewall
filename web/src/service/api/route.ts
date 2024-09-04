import type { LastLevelRouteKey } from '@elegant-router/types';
import { useAuthStore } from '@/store/modules/auth';
import { createStaticRoutes } from '@/router/routes';
import { request } from '../request';

/** get constant routes */
export function fetchGetConstantRoutes() {
  return request<Api.Route.MenuRoute[]>({ url: '/route/getConstantRoutes' });
}

/** get user routes */
export function fetchGetUserRoutes() {
  const authStore = useAuthStore();
  const staticRoute = createStaticRoutes();
  let routes: any = [];
  if (authStore.userInfo.user.roleId === 1) {
    routes = staticRoute.authRoutes;
  } else {
    routes = parseRoutePath(authStore.userInfo.routes, staticRoute.authRoutes);
  }

  return {
    data: { home: <LastLevelRouteKey>authStore.userInfo.home, routes },
    error: null
  };
}

/**
 * whether the route is exist
 *
 * @param routeName route name
 */
export function fetchIsRouteExist(routeName: string) {
  return request<boolean>({ url: '/route/isRouteExist', params: { routeName } });
}

//  返回用户拥有权限的路由表，包含public
export function parseRoutePath(strs: string[], routes: any[]): any[] {
  const newRoutes: any[] = [];
  // 遍历所有路由
  for (const route of routes) {
    // 判断路由是否为常量路由
    if (route.meta && route.meta?.public) {
      newRoutes.push(route);
      // eslint-disable-next-line no-continue
      continue;
    }

    // 判断路由是否包含子路由
    if (route?.children && route.children.length > 0) {
      // 如果包含子路由，则递归调用getRoutePath函数，传入strs和route的children
      // newRoutes.push(...);

      const ch = parseRoutePath(strs, route.children);
      if (ch.length > 0) {
        newRoutes.push({
          ...route,
          children: ch
        });
      }
      // eslint-disable-next-line no-continue
      continue;
    }

    // 判断路由的path是否包含strs中的任意一个
    if (strs.includes(route.name)) {
      // 如果包含，则递归调用getRoutePath函数，传入strs和route的children
      newRoutes.push(route);
    }
  }

  return newRoutes;
}
