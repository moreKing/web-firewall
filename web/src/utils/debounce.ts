/** 防抖函数实现 */

/**
 * wait 毫秒仅执行一次，若中间有新的请求发生，则重置计时wait
 *
 * @param {function} fn 执行函数
 * @param {number} wait 等待时间，毫秒
 * @param {boolean} immediate 是否立即执行
 */
export function debounce(fn: Function, wait = 500, immediate = false) {
  let timeout: any; // 局部全局变量
  return (...args: any[]) => {
    if (timeout) clearTimeout(timeout); // 清除计时器，但是timeout本身还在，只是不会在执行
    if (immediate) {
      // 第一次为true, 执行
      if (!timeout) {
        fn(args);
      }
      // 总是执行第一次操作
      // 多次操作，timeout初始化，多次触发只有当wait等待时间结束timeout才为空
      timeout = setTimeout(() => {
        timeout = null;
      }, wait);
    } else {
      // 总是执行最后一次操作
      timeout = setTimeout(() => {
        fn(args);
      }, wait);
    }
  };
}

/**
 * wait 毫秒仅执行一次
 *
 * @param {function} fn 执行函数
 * @param {number} wait 等待时间，毫秒
 * @param {boolean} immediate 是否立即执行
 */
export function throttle(fn: Function, wait = 500, immediate = false) {
  let timeout: any; // 局部全局变量
  return (...args: any[]) => {
    // if (timeout) clearTimeout(timeout); // 清除计时器，但是timeout本身还在，只是不会在执行
    if (timeout) {
      return;
    }
    if (immediate) {
      // 第一次为true, 执行
      if (!timeout) {
        fn(args);
      }
      // 总是执行第一次操作
      // 多次操作，timeout初始化，多次触发只有当wait等待时间结束timeout才为空
      timeout = setTimeout(() => {
        timeout = null;
      }, wait);
    } else {
      // 总是执行最后一次操作
      timeout = setTimeout(() => {
        fn(args);
        timeout = null;
      }, wait);
    }
  };
}
