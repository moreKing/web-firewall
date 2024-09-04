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
