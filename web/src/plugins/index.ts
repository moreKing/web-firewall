import type { App } from 'vue';
import { permission } from './permission';
import throttle from './throttle';
import debounce from './debounce';

export * from './loading';
export * from './nprogress';
export * from './iconify';
export * from './dayjs';

export function setupDirectives(app: App) {
  // 权限控制指令（演示）
  app.directive('permission', permission);
  app.directive('throttle', throttle);
  app.directive('debounce', debounce);
  // 复制指令
}
