import { requestV1 } from '../request';

//  emailTest, getEmailConf, putEmailConf

export function getAuditLogin(data: any) {
  const params: string[] = [];

  for (const [key, value] of Object.entries(data)) {
    if (value === undefined || value === null || Number.isNaN(value)) {
      // eslint-disable-next-line no-continue
      continue;
    }
    let tmp = value;

    if (typeof value === 'string') {
      tmp = value.trim();
      if (tmp === '') {
        // eslint-disable-next-line no-continue
        continue;
      }
    }
    params.push(`${key}=${tmp}`);
  }

  const param = params.join('&');
  return requestV1<Api.Common.ListRes>({ url: `/audit/login?${param}` });
}

export function cutOnlineLogin(uuid: string) {
  return requestV1({ url: `/audit/cut-login/${uuid}`, method: 'POST' });
}

export function getAuditSettings(data: any) {
  const params: string[] = [];

  for (const [key, value] of Object.entries(data)) {
    if (value === undefined || value === null || Number.isNaN(value)) {
      // eslint-disable-next-line no-continue
      continue;
    }
    let tmp = value;

    if (typeof value === 'string') {
      tmp = value.trim();
      if (tmp === '') {
        // eslint-disable-next-line no-continue
        continue;
      }
    }
    params.push(`${key}=${tmp}`);
  }

  const param = params.join('&');
  return requestV1<Api.Common.ListRes>({ url: `/audit/settings?${param}` });
}

export function getAuditShell(data: any) {
  const params: string[] = [];

  for (const [key, value] of Object.entries(data)) {
    if (value === undefined || value === null || Number.isNaN(value)) {
      // eslint-disable-next-line no-continue
      continue;
    }
    let tmp = value;

    if (typeof value === 'string') {
      tmp = value.trim();
      if (tmp === '') {
        // eslint-disable-next-line no-continue
        continue;
      }
    }
    params.push(`${key}=${tmp}`);
  }

  const param = params.join('&');
  return requestV1<Api.Common.ListRes>({ url: `/audit/shell?${param}` });
}

export function getAuditShellReplayToken() {
  return requestV1<any>({ url: `/audit/shell-token` });
}
