function formateFileSize(byteNum: number) {
  const num = 1024;

  if (byteNum < num) return `${byteNum}B`;
  if (byteNum < num ** 2) return `${(byteNum / num).toFixed(2)}K`;
  if (byteNum < num ** 3) return `${(byteNum / num ** 2).toFixed(2)}M`;

  if (byteNum < num ** 4) return `${(byteNum / num ** 3).toFixed(2)}G`;

  return `${(byteNum / num ** 4).toFixed(2)}T`;
}

export default formateFileSize;
