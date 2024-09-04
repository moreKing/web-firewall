import { requestV1 } from '../request';

export function getSystemStatus() {
  return requestV1({ url: '/system/status' });
}

export function getSystemHome() {
  return requestV1({ url: '/system/home' });
}
