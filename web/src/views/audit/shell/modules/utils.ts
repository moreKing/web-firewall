/**
 * 获取get请求提交的参数 url是请求完整的连接 key是键值对的键值
 *
 * 返回key对应的value值
 */

export function getUrl(url: string, key: string) {
  const tmp = url.match(`${key}=\\w+`);
  if (tmp && tmp.length > 0) {
    return tmp[0].split('=')[1];
  }
  return false;
}

export function getCurrentWs() {
  if (import.meta.env.MODE === 'prod') {
    return window.location.protocol === 'http:' ? `ws://${window.location.host}` : `wss://${window.location.host}`;
  }
  // console.log(import.meta.env.VITE_SERVICE_BASE_URL);
  const url = new URL(import.meta.env.VITE_SERVICE_BASE_URL);
  return url.protocol === 'http:' ? `ws://${url.host}` : `wss://${url.host}`;
}

export function formateFileSize(byteNum: number): string {
  const num = 1024;

  if (byteNum < num) return `${byteNum}B`;
  if (byteNum < num ** 2) return `${(byteNum / num).toFixed(2)}K`;
  if (byteNum < num ** 3) return `${(byteNum / num ** 2).toFixed(2)}M`;

  if (byteNum < num ** 4) return `${(byteNum / num ** 3).toFixed(2)}G`;

  return `${(byteNum / num ** 4).toFixed(2)}T`;
}
