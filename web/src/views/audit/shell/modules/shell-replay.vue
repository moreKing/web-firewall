<script setup lang="ts">
import '@xterm/xterm/css/xterm.css';
import { Terminal } from '@xterm/xterm';
import { FitAddon } from '@xterm/addon-fit';
import { onUnmounted, ref } from 'vue';
import CryptoJS from 'crypto-js';

import { $t } from '@/locales/index.js';
import { debounce } from '@/utils/debounce';
import { getAuditShellReplayToken } from '@/service/api';
import { getCurrentWs } from './utils';

const props = defineProps<{
  data: any;
}>();

const terminalBox = ref<any>(null);
let term: Terminal;
let socket: WebSocket | null;

// 字符串转base64
function getEncode64(str: string) {
  return btoa(
    encodeURIComponent(str).replace(/%([0-9A-F]{2})/g, function toSolidBytes(_match, p1) {
      return String.fromCharCode(Number.parseInt(`0x${p1}`, 16));
    })
  );
}

const showModal = defineModel<boolean>('show');

function afterData() {
  if (socket) {
    socket.close();
  }
  terminalBox.value.innerHTML = '';

  socket = null; // 清空socket
}

async function enterModal() {
  //
  const { data: tokenData } = await getAuditShellReplayToken();
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
  term.write('\x1B[1;3;31m 开始回放录像 \x1B[0m \r\n');

  socket = new WebSocket(`${getCurrentWs()}/api/v1/audit/shell-replay?token=${tokenData.token}&id=${props.data.id}`);
  // socket = new WebSocket('ws://162.14.109.53:8080/myssh')

  socket.binaryType = 'arraybuffer';

  // 打开socket监听事件的方法
  socket.onopen = () => {
    term.onData((data: any) => {
      socket?.send(JSON.stringify({ type: 'stdin', data: getEncode64(data) })); // getEncode64
    });

    fitAddon.fit();
  };
  socket.onclose = () => {
    // term.writeln(`\r\n${$t('common.logout')}`);
    term.write('\r\n\x1B[1;3;31m 结束播放 \x1B[0m \r\n');
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
        // zsentry.consume(recv.data);
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
            socket?.send(JSON.stringify({ type: 'pong' })); // getEncode64
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
    socket?.send(
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
}

onUnmounted(() => {
  if (socket) {
    socket.close();
  }
});
</script>

<template>
  <div>
    <NModal
      v-model:show="showModal"
      :mask-closable="false"
      preset="card"
      class="h-100vh w-full"
      :title="$t('page.audit.replay')"
      :bordered="false"
      :segmented="{
        content: false
      }"
      @after-leave="afterData"
      @after-enter="enterModal"
    >
      <div ref="terminalBox" class="min-height"></div>
    </NModal>
  </div>
</template>

<style scoped>
:deep(.terminal) {
  padding: 20px;
}

:deep(.xterm-viewport) {
  border-radius: 10px;
  overflow: hidden;
}

.min-height {
  height: calc(100vh - 110px);
  background-color: black;
  border-radius: 10px;
}
</style>
