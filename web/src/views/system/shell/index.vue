<script setup lang="ts">
import '@xterm/xterm/css/xterm.css';
import { Terminal } from '@xterm/xterm';
import { FitAddon } from '@xterm/addon-fit';
import { onMounted, onUnmounted, ref } from 'vue';
import './utils/zmodem.devel.js';
import CryptoJS from 'crypto-js';
import { localStg } from '@/utils/storage';
import { $t } from '@/locales/index.js';
import { debounce } from '@/utils/debounce';
import { formateFileSize, getCurrentWs } from './utils/utils';

const terminalBox = ref<any>(null);
let term: Terminal;
let socket: WebSocket;
let upZsession: any;

let uploadStartTime = 0;

// 字符串转base64
function getEncode64(str: string) {
  return btoa(
    encodeURIComponent(str).replace(/%([0-9A-F]{2})/g, function toSolidBytes(_match, p1) {
      return String.fromCharCode(Number.parseInt(`0x${p1}`, 16));
    })
  );
}

// 取消文件上传
function uploadFilecancel() {
  if (upZsession === 'undefined') {
    // window.$message?.error('会话不存在！');
    return;
  }
  try {
    // zsession 每 5s 发送一个 ZACK 包，5s 后会出现提示最后一个包是 ”ZACK“ 无法正常关闭
    // 这里直接设置 _last_header_name 为 ZRINIT，就可以强制关闭了
    // eslint-disable-next-line no-underscore-dangle
    upZsession._last_header_name = 'ZRINIT';
    upZsession.close();
  } catch (e) {
    window.console.log(e);
  }
}

let upfilestate = true;
// js出发文件选择框
function trigerUpload() {
  return new Promise(resolve => {
    const input = document.createElement('input');
    input.setAttribute('type', 'file');
    input.setAttribute('multiple', 'multiple');
    input.accept = '*/*';
    window.addEventListener(
      'focus',
      () => {
        setTimeout(() => {
          if (upfilestate) {
            uploadFilecancel();
          }
        }, 100);
      },
      { once: true }
    );
    input.addEventListener('change', (e: any) => {
      upfilestate = false;
      resolve(e.target.files);
    });

    input.click();
  });
}

// 保存文件
function saveFile(xfer: any, buffer: any) {
  return Zmodem.Browser.save_to_disk(buffer, xfer.get_details().name);
}
// 校验大小
function bytesHuman(bytes: any, precision: number = 1) {
  if (!/^([-+])?|(\.\d+)(\d+(\.\d+)?|(\d+\.)|Infinity)$/.test(bytes)) {
    return '-';
  }
  if (bytes === 0) return '0';
  // if (typeof precision === 'undefined') precision = 1;
  const units = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB', 'BB'];
  const num = Math.floor(Math.log(bytes) / Math.log(1024));
  const value = (bytes / 1024 ** Math.floor(num)).toFixed(precision);
  return `${value} ${units[num]}`;
}

// 实时显示进度
async function updateProgress(xfer: any, action = 'upload') {
  const detail = xfer.get_details();
  const name = detail.name;
  const total = detail.size;
  const offset = xfer.get_offset(); // 单位字节

  // 计算经过的时间
  const nowTime = Date.now();
  const progressTime = nowTime - uploadStartTime;

  let percent;
  if (total === 0 || total === offset) {
    percent = 100;
  } else {
    percent = Math.round((offset / total) * 100);
  }

  term.write(
    `\r${action}   ${name}   ${action === 'upload' ? $t('page.shell.upload') : $t('page.shell.download')}: ${bytesHuman(offset)}   ${$t('page.shell.size')}: ${bytesHuman(total)}   ${$t('page.shell.progress')}: ${
      percent
    }%    ${$t('page.shell.speed')}: ${formateFileSize(Math.ceil((offset / progressTime) * 1000))}/s   ${$t('page.shell.usedTime')}: ${Math.ceil(
      progressTime / 1000
    )}s`
  );
}

// 下载
function downloadFile(zsession: any) {
  uploadStartTime = Date.now();
  zsession.on('offer', (xfer: any) => {
    function on_form_submit() {
      // 跳过不能下载的文件
      if (xfer.get_details().size > 4000 * 1024 * 1024) {
        xfer.skip();
        // window.$message?.error(`sz命令不允许下载  ${xfer.get_details().name} ，文件超过 4 GB`);
        return;
      }

      const FILE_BUFFER: any[] = [];
      xfer.on('input', (payload: any) => {
        updateProgress(xfer, 'download');
        FILE_BUFFER.push(new Uint8Array(payload));
      });

      // 下载完成的回调函数
      xfer.accept().then(() => {
        saveFile(xfer, FILE_BUFFER);
        term.write('\r\n');
        // 一般做审计用
        socket.send(
          JSON.stringify({
            type: 'ignore',
            data: getEncode64(
              `${xfer.get_details().name}(${xfer.get_details().size}) ${$t('page.shell.downloadSuccess')}`
            )
          })
        );
        uploadStartTime = Date.now();
      }, window.console.error.bind(console));
    }
    on_form_submit();
  });
  const promise = new Promise<void>(res => {
    zsession.on('session_end', () => {
      res();
    });
  });
  zsession.start();
  return promise;
}

// 上传
function uploadFile(files: any) {
  new Promise<void>(res => {
    uploadStartTime = Date.now();
    // Zmodem.Browser.send_files(upZsession, files, {
    Zmodem.Browser.send_block_files(upZsession, files, {
      on_offer_response(obj: any, xfer: any) {
        if (xfer) {
          // term.write("\r\n");
        } else {
          term.write(`${obj.name} ${$t('page.shell.skip')}\r\n`);
          // window.$message?.warning(`${obj.name} 被跳过上传`);
        }
      },
      on_progress(_obj: any, xfer: any) {
        updateProgress(xfer, 'upload');
      },
      on_file_complete(obj: any) {
        term.write('\r\n');
        socket.send(
          JSON.stringify({
            type: 'ignore',
            data: getEncode64(`${obj.name}(${obj.size}) 文件上传成功！`)
          })
        );

        uploadStartTime = Date.now();
      }
    })
      .then(upZsession.close.bind(upZsession), window.console.error.bind(console))
      .then(() => {
        res();
      });
  })
    .catch(window.console.error.bind(console))
    .then(() => {});
}

// dom 挂载后 回调
onMounted(() => {
  // 创建一个客户端
  term = new Terminal({
    cursorBlink: true,
    cursorStyle: 'bar',
    fontFamily: 'Consolas, "Courier New", monospace',
    fontSize: 14,
    lineHeight: 1.2,
    fontWeight: 'normal',
    fontWeightBold: 'bold'
  });
  // term.write
  // 将客户端挂载到dom上
  const fitAddon = new FitAddon();
  term.loadAddon(fitAddon);
  term.open(terminalBox.value);
  // fitAddon.fit();

  // 创建socket连接
  term.write('Welcome Used \x1B[1;3;31m Web Firewall\x1B[0m !\r\n');
  // 获取当前请求的主机跟账号

  // let url = getCurrentWs()
  // if (window.location.protocol == 'http:') {
  //   url = 'ws://' + window.location.host
  // } else if (window.location.protocol == 'https:') {
  //   url = 'wss://' + window.location.host
  // }
  // url = 'ws://192.168.177.3:8080'

  socket = new WebSocket(
    // 'ws://127.0.0.1:8080/api/v1/ssh?id=' +
    `${getCurrentWs()}/api/v1/system/shell?Authorization=${localStg.get('token')}`
  );
  // socket = new WebSocket('ws://162.14.109.53:8080/myssh')

  // 实现rz sz
  const zsentry = new Zmodem.Sentry({
    to_terminal(_octets: any) {}, // i.e. send to the terminal
    on_detect(detection: any) {
      // 判断当前输入的命令是上传或者下载 ,当 Sentry 检测到新的 ZMODEM 时
      window.console.log('新的zmodem');

      const zsession = detection.confirm();
      let promise;
      if (zsession.type === 'receive') {
        promise = downloadFile(zsession);
        promise.catch(window.console.error.bind(console)).then(() => {});
      } else {
        upZsession = zsession;
        // 打开文件选择对话框
        // showModal.value = true
        // promise = uploadFile()
        trigerUpload().then((files: any) => {
          // console.log(files)
          if (upZsession === 'undefined') {
            // window.$message?.error('会话不存在！');
          }
          if (files.length <= 0) {
            // window.$message?.error('请先选择需要上传的文件！');
            return;
          }

          if (files.length > 10) {
            // window.$message?.error('上传的文件不能超过10个！');
            uploadFilecancel();
            return;
          }
          uploadFile(files);
        });
      }
    },
    on_retract() {},
    sender(octets: any) {
      socket.send(new Uint8Array(octets));
    } // 即发送到对等的ZMODEM
  });

  socket.binaryType = 'arraybuffer';

  // 打开socket监听事件的方法
  socket.onopen = () => {
    term.onData((data: any) => {
      // console.log('输入：', data)
      // socket.send(JSON.stringify({ type: "stdin", data: getEncode64(data) })) //getEncode64
      socket.send(JSON.stringify({ type: 'stdin', data: getEncode64(data) })); // getEncode64
    });

    window.$message?.success($t('page.login.common.loginSuccess'));
    fitAddon.fit();
  };
  socket.onclose = () => {
    term.writeln(`\r\n${$t('common.logout')}`);
  };
  socket.onerror = (err: any) => {
    // console.log(err);
    window.$dialog?.error({
      title: $t('common.error'),
      content: $t('page.shell.wsError'),
      positiveText: $t('common.confirm'),
      onPositiveClick: () => {
        window.location.reload();
      }
    });
    term.writeln('err :', err);
  };
  // 接收数据
  socket.onmessage = (recv: any) => {
    try {
      //   console.log(typeof recv.data)
      if (typeof recv.data === 'object') {
        // rz sz 文件传输
        zsentry.consume(recv.data);
      } else {
        // 命令结果回显
        const msg = JSON.parse(recv.data);
        switch (msg.type) {
          case 'stdin':
            term.write(CryptoJS.enc.Base64.parse(msg.data).toString(CryptoJS.enc.Utf8));
            break;
          case 'console':
            window.console.log(CryptoJS.enc.Base64.parse(msg.data).toString(CryptoJS.enc.Utf8));
            break;
          case 'error':
            // window.$dialog?.warning({
            //   title: '错误',
            //   content: CryptoJS.enc.Base64.parse(msg.data).toString(CryptoJS.enc.Utf8),
            //   positiveText: '确定'
            // });
            break;

          case 'ping':
            socket.send(JSON.stringify({ type: 'pong' })); // getEncode64
            break;

          default:
            break;
        }
      }
    } catch (e) {
      // console.log('unsupport data', recv.data)
      if (typeof recv.data === 'string') {
        const data = recv.data.split('"data":"');
        if (data.length > 0) {
          const datacon = data[data.length - 1].split('"');
          if (datacon.length > 0) {
            // console.log(datacon[0])
            // window.$message?.error('读取数据乱码！');
            term.write(atob(datacon[0]));
          }
        }
      }
      window.console.log(e);

      // term.write(CryptoJS.enc.Base64.parse(recv.data).toString(CryptoJS.enc.Utf8))
    }
  };

  // 客户端发送大小改变，后端也需要跟着改变
  term.onResize(({ cols, rows }: { cols: number; rows: number }) => {
    socket.send(
      JSON.stringify({
        type: 'resize',
        cols,
        rows
      })
    );
  });

  window.addEventListener(
    'resize',
    debounce(
      () => {
        fitAddon.fit();
      },
      100,
      false
    ),
    false
  );

  //

  onUnmounted(() => {
    if (socket) {
      socket.close();
    }
  });
});
</script>

<template>
  <div>
    <div ref="terminalBox" class="min-height"></div>
  </div>
</template>

<style scoped>
:deep(.terminal) {
  padding: 20px;
}

:deep(.xterm-viewport) {
  border-radius: 10px;
}

.min-height {
  height: calc(100vh - 180px);
}
</style>
