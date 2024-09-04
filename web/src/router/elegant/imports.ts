/* eslint-disable */
/* prettier-ignore */
// Generated by elegant-router
// Read more: https://github.com/soybeanjs/elegant-router

import type { RouteComponent } from "vue-router";
import type { LastLevelRouteKey, RouteLayout } from "@elegant-router/types";

import BaseLayout from "@/layouts/base-layout/index.vue";
import BlankLayout from "@/layouts/blank-layout/index.vue";

export const layouts: Record<RouteLayout, RouteComponent | (() => Promise<RouteComponent>)> = {
  base: BaseLayout,
  blank: BlankLayout,
};

export const views: Record<LastLevelRouteKey, RouteComponent | (() => Promise<RouteComponent>)> = {
  403: () => import("@/views/_builtin/403/index.vue"),
  404: () => import("@/views/_builtin/404/index.vue"),
  406: () => import("@/views/_builtin/406/index.vue"),
  500: () => import("@/views/_builtin/500/index.vue"),
  "iframe-page": () => import("@/views/_builtin/iframe-page/[url].vue"),
  login: () => import("@/views/_builtin/login/index.vue"),
  audit_login: () => import("@/views/audit/login/index.vue"),
  audit_settings: () => import("@/views/audit/settings/index.vue"),
  home: () => import("@/views/home/index.vue"),
  "policy_input-limit": () => import("@/views/policy/input-limit/index.vue"),
  policy_input: () => import("@/views/policy/input/index.vue"),
  "policy_output-limit": () => import("@/views/policy/output-limit/index.vue"),
  policy_output: () => import("@/views/policy/output/index.vue"),
  system_basic: () => import("@/views/system/basic/index.vue"),
  system_shell: () => import("@/views/system/shell/index.vue"),
  "user-center": () => import("@/views/user-center/index.vue"),
};