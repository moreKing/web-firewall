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

// 源地址转换
export function getSnatPolicyList() {
  return requestV1<Api.Common.ListRes>({
    url: `/route/snat`,
    method: 'GET'
  });
}

export function addSnatPolicy(data: any) {
  return requestV1({ url: `/route/snat`, data, method: 'POST' });
}

export function updateSnatPolicy(data: any) {
  return requestV1({ url: `/route/snat/${data.id}`, data, method: 'PUT' });
}

export function changeSnatPolicyPosition(data: any) {
  return requestV1({ url: `/route/snat/position/${data.id}`, data, method: 'PUT' });
}
export function deleteSnatPolicy(id: number) {
  return requestV1({ url: `/route/snat/${id}`, method: 'DELETE' });
}

// 目的地址转换
export function getDnatPolicyList() {
  return requestV1<Api.Common.ListRes>({
    url: `/route/dnat`,
    method: 'GET'
  });
}

export function addDnatPolicy(data: any) {
  return requestV1({ url: `/route/dnat`, data, method: 'POST' });
}

export function updateDnatPolicy(data: any) {
  return requestV1({ url: `/route/dnat/${data.id}`, data, method: 'PUT' });
}

export function changeDnatPolicyPosition(data: any) {
  return requestV1({ url: `/route/dnat/position/${data.id}`, data, method: 'PUT' });
}
export function deleteDnatPolicy(id: number) {
  return requestV1({ url: `/route/dnat/${id}`, method: 'DELETE' });
}

// 转发策略
export function getForwardPolicyList() {
  return requestV1<Api.Common.ListRes>({
    url: `/route/forward`,
    method: 'GET'
  });
}

export function addForwardPolicy(data: any) {
  return requestV1({ url: `/route/forward`, data, method: 'POST' });
}

export function updateForwardPolicy(data: any) {
  return requestV1({ url: `/route/forward/${data.id}`, data, method: 'PUT' });
}

export function changeForwardPolicyPosition(data: any) {
  return requestV1({ url: `/route/forward/position/${data.id}`, data, method: 'PUT' });
}
export function deleteForwardPolicy(id: number) {
  return requestV1({ url: `/route/forward/${id}`, method: 'DELETE' });
}

// 转发流控
export function getForwardLimitPolicyList() {
  return requestV1<Api.Common.ListRes>({
    url: `/route/limit`,
    method: 'GET'
  });
}

export function addForwardLimitPolicy(data: any) {
  return requestV1({ url: `/route/limit`, data, method: 'POST' });
}

export function updateForwardLimitPolicy(data: any) {
  return requestV1({ url: `/route/limit/${data.id}`, data, method: 'PUT' });
}

export function changeForwardLimitPolicyPosition(data: any) {
  return requestV1({ url: `/route/limit/position/${data.id}`, data, method: 'PUT' });
}
export function deleteForwardLimitPolicy(id: number) {
  return requestV1({ url: `/route/limit/${id}`, method: 'DELETE' });
}
