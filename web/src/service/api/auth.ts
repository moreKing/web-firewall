import { request, requestV1 } from '../request';

/**
 * Login
 *
 * @param userName User name
 * @param password Password
 */
export function fetchLogin(data: any) {
  return requestV1<Api.Auth.LoginToken>({
    url: '/login',
    method: 'post',
    data
  });
}

/** Get user info */
export async function fetchGetUserInfo() {
  const req = await requestV1<Api.Auth.UserInfo>({ url: '/public/profile' });
  // 将用户路由权限还原

  return req;
}

/**
 * Refresh token
 *
 * @param refreshToken Refresh token
 */
export function fetchRefreshToken(refreshToken: string) {
  return request<Api.Auth.LoginToken>({
    url: '/auth/refreshToken',
    method: 'post',
    data: {
      refreshToken
    }
  });
}

/**
 * return custom backend error
 *
 * @param code error code
 * @param msg error message
 */
export function fetchCustomBackendError(code: string, msg: string) {
  return request({ url: '/auth/error', params: { code, msg } });
}
