import { requestV1 } from '../request';

//  emailTest, getEmailConf, putEmailConf

export function emailTest(addr: string) {
  return requestV1({ url: `/system/email`, data: { to: addr }, method: 'POST' });
}

export function getEmailConf() {
  return requestV1({ url: `/system/email` });
}

export function putEmailConf(data: any) {
  return requestV1({ url: `/system/email`, data, method: 'PUT' });
}

// getWebConf, setWebConf

export function getWebConf() {
  return requestV1({ url: `/system/session-config` });
}
export function setWebConf(data: any) {
  return requestV1({ url: `/system/session-config`, data, method: 'PUT' });
}

// setPasswordComplex
export function setPasswordComplex(data: any) {
  return requestV1({ url: `/system/password-complex`, data, method: 'PUT' });
}

//  getMessageConf, putMessageConf, testMessage

export function testMessage(addr: string) {
  return requestV1({ url: `/system/message`, data: { to: addr }, method: 'POST' });
}

export function getMessageConf() {
  return requestV1({ url: `/system/message` });
}

export function putMessageConf(data: any) {
  return requestV1({ url: `/system/message`, data, method: 'PUT' });
}

//  getAuthConf, putAuthConf
export function getAuthConf() {
  return requestV1({ url: `/system/auth-conf` });
}

export function putAuthConf(data: any) {
  return requestV1({ url: `/system/auth-conf`, data, method: 'PUT' });
}

// checkPort

export function checkPort(ip: string, port: number) {
  return requestV1({ url: `/system/check-port`, data: { ip, port }, method: 'POST' });
}
