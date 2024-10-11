/**
 * 查看此源码之前，你需要具备一些简单的位运算知识：
 *
 * & ：与运算 表示若两个二进制位都为1，则结果为1，否则为0。
 *
 * |： 或运算 表示若两个二进制位都为0，则结果为0，否则为1。
 *
 * ~ 取反 表示对一个二进制位取反。
 *
 * ^ 异或运算 表示若两个二进制位不相同，则结果为1，相同否则为0。
 *
 * << 左移运算 > > 右移运算 > > > 无符号右移运算
 *
 * 一些用法
 *
 * 判断x的n位是否为1： x&(1<<n)
 *
 * 判断x的n位是否为0： !(x&(1<<n))
 *
 * 本次使用以下两个关键的位运算
 *
 * 将x的n位移为1： x|(1<<n)
 *
 * 将x的n位移为0： x&(~(1<<n))
 *
 * x从右往左第一个1右边都变1： x|(x-1)
 *
 * x最右边变1： x|(x+1)
 */

/**
 * @param ipaddr
 * @returns
 */
export function checkIpAddr(ipaddr: string) {
  const tmpipaddr = ipaddr.trim();
  if (tmpipaddr === '') {
    return false;
  }
  const ss = tmpipaddr.split('.');
  if (ss.length !== 4) {
    return false;
  }
  let i = 0;
  for (i = 0; i < ss.length; i += 1) {
    if (!/^[\d]+$/.test(ss[i])) {
      return false;
    }
    if (ss[i].trim() === '') {
      return false;
    }
    if (Number.parseInt(ss[i], 10) < 0 || Number.parseInt(ss[i], 10) > 255) {
      return false;
    }
  }
  return true;
}

export function checkNetMask(ipaddr: string) {
  const tmpipaddr = ipaddr.trim();
  if (tmpipaddr === '') {
    return false;
  }

  if (tmpipaddr.length === 2 || tmpipaddr.length === 1) {
    if (Number.parseInt(tmpipaddr, 10) < 0 || Number.parseInt(tmpipaddr, 10) > 32) {
      return false;
    }
    return true;
  }

  const ss = tmpipaddr.split('.');
  if (ss.length !== 4) {
    return false;
  }
  let i = 0;
  for (i = 0; i < ss.length; i += 1) {
    if (Number.parseInt(ss[i], 10) < 0 || Number.parseInt(ss[i], 10) > 255) {
      return false;
    }
  }
  return true;
}

//
export function checkIpMask(ipaddr: string) {
  const tmpipaddr = ipaddr.trim();
  if (tmpipaddr === '') {
    return false;
  }

  if (tmpipaddr.includes('-')) {
    const ss = tmpipaddr.split('-');

    if (ss.length !== 2) {
      return false;
    }
    if (!checkIpAddr(ss[0])) {
      return false;
    }

    if (!checkIpAddr(ss[1])) {
      return false;
    }
    return true;
  }

  if (tmpipaddr.includes('/')) {
    const ss = tmpipaddr.split('/');
    if (ss.length !== 2) {
      return false;
    }
    if (!checkIpAddr(ss[0])) {
      return false;
    }
    if (!checkNetMask(ss[1])) {
      return false;
    }
    return true;
  }

  return checkIpAddr(ipaddr);
}

// 判断ip地址是否在网段内
export function checkIpInNet(ipaddr: string, netmask: string): boolean {
  // ipaddr = '192.168.1.1';
  // netmask = '192.168.1.0/24';

  if (netmask.includes('/')) {
    const ss = netmask.split('/');
    if (ss.length !== 2) {
      return false;
    }
    if (!checkIpAddr(ss[0])) {
      return false;
    }
    if (!checkNetMask(ss[1])) {
      return false;
    }
    // console.log(netmask, ' ip format success  0.0.0.0/0', ipaddr);
    // const range = getIPRange(ss[0], Number.parseInt(ss[1], 10)); // 网段地址
    // if (ipToNumber(range.startAddress) <= ipToNumber(ipaddr) && ipToNumber(range.endAddress) >= ipToNumber(ipaddr))
    //   return true;

    const { startAddress, endAddress } = parseIpRangeToNumber(ss[0], Number.parseInt(ss[1], 10));
    const currentIp = ipToNumber(ipaddr);
    if (startAddress <= currentIp && endAddress >= currentIp) return true;

    return false;
  }

  if (netmask.includes('-')) {
    const ss = netmask.split('-');
    if (ss.length !== 2) {
      return false;
    }
    if (!checkIpAddr(ss[0])) {
      return false;
    }
    if (!checkIpAddr(ss[1])) {
      return false;
    }
    if (ipToNumber(ss[0]) <= ipToNumber(ipaddr) && ipToNumber(ss[1]) >= ipToNumber(ipaddr)) return true;
    return false;
  }

  if (ipaddr.trim() === netmask.trim()) return true;
  return false;
}

function ipToNumber(ipaddr: string): number {
  const ss = ipaddr.trim().split('.');
  if (ss.length !== 4) {
    return 0;
  }

  return ss.reduce((res: number, item: string) => {
    return res * 256 + Number.parseInt(item, 10);
  }, 0);
}

function parseIpRangeToNumber(ip: string, maskBits: number) {
  // 坑爹的事情来了，ip地址是32位，而js位运算最高仅支持32位，1<<32 == 1 并且导致符号丢失正数变成负数 所以需要使用字符串的方式进行位运算
  // const ipN = ipToNumber(ip);
  // console.log(ipN, maskBits);
  // // eslint-disable-next-line no-bitwise
  // const bit = (1 << 32) | (1 << (32 - maskBits));
  // console.log(bit.toString(2));
  // // eslint-disable-next-line no-bitwise
  // const startAddress = ipN & (~bit & ~(1 << (32 - maskBits)));
  // // eslint-disable-next-line no-bitwise
  // const endAddress = ipN | ~(~bit | (bit - 1));

  const ip_bin_str = ip_to_binary_string(ip);
  const startAddress_str = ip_bin_str.substring(0, maskBits) + '0'.repeat(32 - maskBits);
  const endAddress_str = ip_bin_str.substring(0, maskBits) + '1'.repeat(32 - maskBits);

  let startAddress = 0;
  let endAddress = 0;
  for (let i = 0; i < 4; i += 1) {
    const curr_num = startAddress_str.substring(i * 8, (i + 1) * 8);
    startAddress = startAddress * 256 + Number.parseInt(curr_num, 2);

    const curr_num2 = endAddress_str.substring(i * 8, (i + 1) * 8);
    endAddress = endAddress * 256 + Number.parseInt(curr_num2, 2);
  }

  return {
    startAddress,
    endAddress
  };
}

function ip_to_binary_string(ip: string) {
  if (checkIpAddr(ip)) {
    let ip_str = '';
    const ip_arr = ip.split('.');
    for (let i = 0; i < 4; i += 1) {
      const curr_num = ip_arr[i];
      const number_bin = Number.parseInt(curr_num, 10);
      let number_bin_str = number_bin.toString(2);
      number_bin_str = `${'0'.repeat(8 - number_bin_str.length)}${number_bin_str}`;
      ip_str += number_bin_str;
    }
    return ip_str;
  }

  return '';
}

export function checkPort(port: string, portRange: string) {
  const portNum = Number.parseInt(port, 10);
  if (portRange.includes('-')) {
    const ss = portRange.split('-');
    if (ss.length !== 2) {
      return false;
    }

    return Number.parseInt(ss[0], 10) <= portNum && Number.parseInt(ss[1], 10) >= portNum;
  }
  return portNum === Number.parseInt(portRange, 10);
}

// 判断端口是否合法
export function checkPortString(port: string) {
  if (!port || port === '') return false;
  if (port.includes('-')) {
    const ss = port.split('-');
    if (ss.length !== 2) {
      return false;
    }

    const pattern = /^\d+$/;
    if (!pattern.test(ss[0])) return false;
    if (!pattern.test(ss[1])) return false;

    const start = Number.parseInt(ss[0], 10);
    const end = Number.parseInt(ss[1], 10);
    if (start <= 0 || end <= 0 || start > 65535 || end > 65535) return false;
    return start < end;
  }
  const pattern = /^\d+$/;
  if (!pattern.test(port)) return false;

  const start = Number.parseInt(port, 10);
  return !(start <= 0 || start > 65535);
}
