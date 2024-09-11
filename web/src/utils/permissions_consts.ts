// 防火墙策略接口

export const GET_FIREWALL_POLICY = 'GET:/api/v1/firewall/policy/{chain}';
export const ADD_FIREWALL_POLICY = 'POST:/api/v1/firewall/policy/{chain}';
export const UPDATE_FIREWALL_POLICY = 'PUT:/api/v1/firewall/policy/{id}';
export const CHANGE_FIREWALL_POLICY_POSITION = 'PUT:/api/v1/firewall/policy/position/{id}';
export const DELETE_FIREWALL_POLICY = 'DELETE:/api/v1/firewall/policy/{id}';

//  基本设置
// 本地密码复杂度
export const SET_PASSWORD_COMPLEX = 'PUT:/api/v1/system/password-complex';
export const SET_WEB_CONF = 'PUT:/api/v1/system/session-config';
export const GET_WEB_CONF = 'GET:/api/v1/system/session-config';

// 邮件
export const EMAIL_TEST = 'POST:/api/v1/system/email';
export const GET_EMAIL_CONF = 'GET:/api/v1/system/email';
export const PUT_EMAIL_CONF = 'PUT:/api/v1/system/email';

// 短信
export const TEST_MESSAGE = 'POST:/api/v1/system/message';
export const GET_MESSAGE_CONF = 'GET:/api/v1/system/message';
export const PUT_MESSAGE_CONF = 'PUT:/api/v1/system/message';

// 认证设置超时
export const GET_AUTH_CONF = 'POST:/api/v1/system/auth-conf';
export const PUT_AUTH_CONF = 'PUT:/api/v1/system/auth-conf';

// 切断在线登录会话
export const CUT_ONLINE_LOGIN = 'POST:/api/v1/audit/cut-login/{uuid}';
export const GET_AUDIT_SHELL_REPLAY_TOKEN = 'POST:/api/v1/audit/shell-token';
