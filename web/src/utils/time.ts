// 传入Timestamp 秒级单位时间戳，返回一个hh:mm:ss 格式的string
export function formateTimestamp(Timestamp: number) {
  if (!Timestamp || Timestamp <= 0) {
    return '00:00:00';
  }
  let h: number;
  let m;
  let tmp;
  let res = '';

  if (Timestamp > 3600) {
    h = Math.floor(Timestamp / 3600);
    if (h < 10) {
      res = `0${h}`;
    } else {
      res = `${h}`;
    }
    tmp = Timestamp % 3600;
  } else {
    res = '00';
    tmp = Timestamp;
  }

  if (tmp > 60) {
    m = Math.floor(tmp / 60);
    if (m < 10) {
      res += `:0${m}`;
    } else {
      res += `:${m}`;
    }
    tmp %= 60;
  } else {
    res += ':00';
  }

  if (tmp < 10) {
    res += `:0${tmp}`;
  } else {
    res += `:${tmp}`;
  }

  return res;
}
