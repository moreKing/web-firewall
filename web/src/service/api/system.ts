import { requestV1 } from '../request';

export function getSystemStatus() {
  return requestV1({ url: '/system/status' });
}

export function getSystemHome() {
  return requestV1({ url: '/system/home' });
}

export function getSystemKernel() {
  return requestV1({ url: '/system/kernel' });
}

export function setSystemKernel(data: any) {
  return requestV1({ url: '/system/kernel', data, method: 'put' });
}
