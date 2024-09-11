<script setup lang="ts">
import type { Ref } from 'vue';
import { computed, h, ref } from 'vue';
import { NButton, NTag } from 'naive-ui';
import dayjs from 'dayjs';
import CustomTable from '@/components/table/custom-table.vue';
import { $t } from '@/locales';
import { formateTimestamp } from '@/utils/time';
import { cutOnlineLogin, getAuditShell } from '@/service/api';
import { useAuth } from '@/hooks/business/auth';
import { CUT_ONLINE_LOGIN, GET_AUDIT_SHELL_REPLAY_TOKEN } from '@/utils/permissions_consts';
import UserSearchVue from './modules/user-search.vue';
import ShellReplay from './modules/shell-replay.vue';

const { hasAuth } = useAuth();

interface SearchParams {
  page: number;
  limit: number;
  total: number;
  username: string;
  loginname: string;
  timestamp: number;
  timeRange: [number, number] | null;
  clientIp: string;
  online: number;
}

const listParams: Ref<SearchParams> = ref({
  page: 1,
  limit: 10,
  total: 0,
  username: '',
  loginname: '',
  timestamp: 0,
  timeRange: null,
  startTime: undefined,
  endTime: undefined,
  clientIp: '',
  online: 0
});

const resetKey = ref(0);
function resetSearchParams() {
  listParams.value = {
    page: 1,
    limit: 10,
    total: 0,
    username: '',
    timeRange: null,
    loginname: '',
    timestamp: 0,
    clientIp: '',
    online: 0
  };
  resetKey.value += 1;
  getData();
}

const showReplay = ref(false);
const rowData = ref({});

const columns = computed<any>(() => [
  {
    key: 'loginname',
    title: $t('datatable.loginname'),
    align: 'center',
    minWidth: 200
  },
  {
    key: 'username',
    title: $t('datatable.username'),
    align: 'center',
    minWidth: 200
  },
  {
    key: 'clientIp',
    title: $t('page.basic.addr'),
    align: 'center',
    minWidth: 200
  },
  {
    title: $t('datatable.onlineTime'),
    align: 'center',
    ellipsis: true,
    minWidth: 200,
    render(row: any) {
      if (!row.online && row.success) {
        return h('span', null, { default: () => formateTimestamp(row.logoutAt - row.createdAt) });
      }

      return null;
    }
  },
  {
    title: $t('datatable.loginTime'),
    key: 'createdAt',
    align: 'center',
    minWidth: 200,
    render(row: any) {
      return h('span', null, {
        // format(row.createdAt, 'yyyy-MM-dd HH:mm:ss')
        default: () =>
          row.createdAt && row.createdAt > 0 ? dayjs.unix(row.createdAt).format('YYYY-MM-DD HH:mm:ss') : ''
      });
    }
  },
  {
    title: $t('datatable.logoutTime'),
    key: 'logoutAt',
    align: 'center',
    minWidth: 200,
    render(row: any) {
      // if (row.online && row.success) {
      //   return h('span', null, { default: () => formateTimestamp(row.logoutAt - row.loginAt) });
      // }

      if (row.online && row.success) {
        if (hasAuth(CUT_ONLINE_LOGIN)) {
          return h(
            NTag,
            {
              type: 'success',
              quaternary: true,
              onClick: () => {
                cutLine(row);
              }
            },
            { default: () => $t('page.audit.online') }
          );
        }
        return null;
      }

      return h('span', null, {
        // format(row.createdAt, 'yyyy-MM-dd HH:mm:ss')
        default: () => (row.logoutAt && row.logoutAt > 0 ? dayjs.unix(row.logoutAt).format('YYYY-MM-DD HH:mm:ss') : '')
      });
    }
  },
  {
    title: $t('datatable.action'),
    align: 'center',
    minWidth: 100,
    render(row: any) {
      if (hasAuth(GET_AUDIT_SHELL_REPLAY_TOKEN)) {
        return h(
          NButton,
          {
            type: 'primary',
            quaternary: true,
            onClick: () => {
              // window.open(`/audit/shell/replay/${row.uuid}`);
              rowData.value = row;
              showReplay.value = true;
            }
          },
          { default: () => $t('page.audit.replay') }
        );
      }
      return null;
    }
  }
]);

async function cutLine(row: any) {
  window.$dialog?.warning({
    title: $t('common.dialog.warning'),
    content: () =>
      h('p', null, [
        h('span', null, { default: () => $t('page.audit.logOffTip') }),
        h('p', { style: 'color: red;margin-top:10px' }, { default: () => row.loginname })
      ]),
    positiveText: $t('common.confirm'),
    negativeText: $t('common.cancel'),
    onPositiveClick: async () => {
      //   message.success('我就知道');
      const { error } = await cutOnlineLogin(row.uuid);
      if (error) return;
      getData();
      window.$message?.success($t('common.dialog.success'));
    }
  });
}

const loading = ref(false);
const data: Ref<any[]> = ref([]);
async function getData() {
  loading.value = true;

  const { data: res, error } = await getAuditShell({
    username: listParams.value.username,
    loginname: listParams.value.loginname,
    startTime: listParams.value.timeRange && listParams.value.timeRange[0] ? listParams.value.timeRange[0] / 1000 : 0,
    endTime:
      listParams.value.timeRange && listParams.value.timeRange[1]
        ? Math.floor(new Date(listParams.value.timeRange[1]).setHours(23, 59, 59, 999) / 1000)
        : 0,
    page: listParams.value.page - 1,
    limit: listParams.value.limit,
    clientIp: listParams.value.clientIp,
    online: listParams.value.online
  });
  loading.value = false;
  if (error) return;
  data.value = res.data;
  listParams.value.total = res.total;
  listParams.value.timestamp = res.timestamp;
}

getData();
</script>

<template>
  <NSpin :show="loading">
    <ShellReplay v-model:show="showReplay" :data="rowData" />
    <NGrid x-gap="12" y-gap="12" cols="24" item-responsive>
      <NGi span="24">
        <div class="min-h-500px">
          <UserSearchVue
            v-model:model="listParams"
            :reset-key="resetKey"
            @reset="resetSearchParams"
            @search="getData"
          />
          <NCard :title="$t('route.audit_shell')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
            <!-- <NScrollbar x-scrollable> -->
            <CustomTable v-model:paramas="listParams" :columns="columns" :data="data" @get-data="getData" />
            <!-- </NScrollbar> -->
          </NCard>
        </div>
      </NGi>
    </NGrid>
  </NSpin>
</template>

<style scoped></style>
