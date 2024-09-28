import { requestV1 } from '../request';

//  emailTest, getEmailConf, putEmailConf

export function getFirewallPolicyList(t: number) {
  return requestV1<Api.Common.ListRes>({
    url: `/firewall/policy/${t}`,
    method: 'GET'
  });
}

export function addFirewallPolicy(data: any) {
  return requestV1({ url: `/firewall/policy/${data.chain}`, data, method: 'POST' });
}

export function updateFirewallPolicy(data: any) {
  return requestV1({ url: `/firewall/policy/${data.id}`, data, method: 'PUT' });
}

// 修改策略位置
export function changeFirewallPolicyPosition(data: any) {
  return requestV1({ url: `/firewall/policy/position/${data.id}`, data, method: 'PUT' });
}

export function deleteFirewallPolicy(id: number) {
  return requestV1({ url: `/firewall/policy/${id}`, method: 'DELETE' });
}
// 入站策略
export function getInputPolicyList() {
  return requestV1<Api.Common.ListRes>({
    url: `/policy/input`,
    method: 'GET'
  });
}

export function addInputPolicy(data: any) {
  return requestV1({ url: `/policy/input`, data, method: 'POST' });
}

export function updateInputPolicy(data: any) {
  return requestV1({ url: `/policy/input/${data.id}`, data, method: 'PUT' });
}

export function changeInputPolicyPosition(data: any) {
  return requestV1({ url: `/policy/input/position/${data.id}`, data, method: 'PUT' });
}
export function deleteInputPolicy(id: number) {
  return requestV1({ url: `/policy/input/${id}`, method: 'DELETE' });
}

// 入站流控
export function getInputLimitPolicyList() {
  return requestV1<Api.Common.ListRes>({
    url: `/policy/input-limit`,
    method: 'GET'
  });
}

export function addInputLimitPolicy(data: any) {
  return requestV1({ url: `/policy/input-limit`, data, method: 'POST' });
}

export function updateInputLimitPolicy(data: any) {
  return requestV1({ url: `/policy/input-limit/${data.id}`, data, method: 'PUT' });
}

export function changeInputLimitPolicyPosition(data: any) {
  return requestV1({ url: `/policy/input-limit/position/${data.id}`, data, method: 'PUT' });
}
export function deleteInputLimitPolicy(id: number) {
  return requestV1({ url: `/policy/input-limit/${id}`, method: 'DELETE' });
}

// 出站策略

export function getOutputPolicyList() {
  return requestV1<Api.Common.ListRes>({
    url: `/policy/output`,
    method: 'GET'
  });
}

export function addOutputPolicy(data: any) {
  return requestV1({ url: `/policy/output`, data, method: 'POST' });
}

export function updateOutputPolicy(data: any) {
  return requestV1({ url: `/policy/output/${data.id}`, data, method: 'PUT' });
}

export function changeOutputPolicyPosition(data: any) {
  return requestV1({ url: `/policy/output/position/${data.id}`, data, method: 'PUT' });
}
export function deleteOutputPolicy(id: number) {
  return requestV1({ url: `/policy/output/${id}`, method: 'DELETE' });
}

// 出站流控
export function getOutputLimitPolicyList() {
  return requestV1<Api.Common.ListRes>({
    url: `/policy/output-limit`,
    method: 'GET'
  });
}

export function addOutputLimitPolicy(data: any) {
  return requestV1({ url: `/policy/output-limit`, data, method: 'POST' });
}

export function updateOutputLimitPolicy(data: any) {
  return requestV1({ url: `/policy/output-limit/${data.id}`, data, method: 'PUT' });
}

export function changeOutputLimitPolicyPosition(data: any) {
  return requestV1({ url: `/policy/output-limit/position/${data.id}`, data, method: 'PUT' });
}
export function deleteOutputLimitPolicy(id: number) {
  return requestV1({ url: `/policy/output-limit/${id}`, method: 'DELETE' });
}
