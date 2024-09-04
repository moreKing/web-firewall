import { requestBase, requestV1 } from '../request';

export function GetPasswordComplex() {
  return requestV1({ url: '/public/password-complex' });
}

interface PublicSetPasswordData {
  oldPassword: string;
  newPassword: string;
}

export function PublicSetPassword(data: PublicSetPasswordData) {
  return requestV1({
    url: '/public/set-password',
    method: 'put',
    data
  });
}

export function PublicSetPersonProfile(data: any) {
  return requestV1({
    url: '/public/profile',
    method: 'put',
    data
  });
}

export function GetOpenApi() {
  return requestBase({
    url: '/api.json',
    method: 'get'
  });
}

export function PublicLogout() {
  return requestV1({
    url: '/public/logout',
    method: 'post'
  });
}
