<script setup lang="ts">
import type { Ref } from 'vue';
import { computed, h, ref } from 'vue';
import { NTag } from 'naive-ui';
import dayjs from 'dayjs';
import CustomTable from '@/components/table/custom-table.vue';
import { $t } from '@/locales';
import { getAuditSettings } from '@/service/api';
import UserSearchVue from './modules/user-search.vue';

interface SearchParams {
  page: number;
  limit: number;
  total: number;
  username: string;
  loginname: string;
  timestamp: number;
  timeRange: [number, number] | null;
  state: number;
  method: string;
  path: string;
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
  state: 0,
  method: '',
  path: ''
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
    state: 0,
    method: '',
    path: ''
  };
  resetKey.value += 1;
  getData();
}

const columns = computed<any>(() => [
  {
    title: $t('datatable.createTime'),
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
    key: 'name',
    title: $t('page.audit.opration'),
    align: 'center',
    minWidth: 120,
    render(row: any) {
      return h(
        NTag,
        { bordered: false, color: { color: '#0d948815', textColor: '#0d9488' } },
        { default: () => row.name }
      );
    }
  },

  {
    key: 'success',
    title: $t('datatable.status'),
    align: 'center',
    minWidth: 80,
    render(row: any) {
      if (row.success)
        return h(NTag, { bordered: false, type: 'success' }, { default: () => $t('common.dialog.success') });

      return h(NTag, { bordered: false, type: 'error' }, { default: () => $t('common.dialog.failed') });
    }
  },
  {
    key: 'requestMethod',
    title: $t('page.audit.method'),
    align: 'center',
    minWidth: 80,
    render(row: any) {
      let typeStr: any = 'info';

      switch (row.requestMethod) {
        case 'POST':
          typeStr = 'success';
          break;
        case 'PUT':
          typeStr = 'warning';
          break;
        case 'DELETE':
          typeStr = 'error';
          break;
        default:
          typeStr = 'info';
          break;
      }

      return h(NTag, { bordered: false, type: typeStr }, { default: () => row.requestMethod });
    }
  },
  {
    key: 'requestPath',
    title: $t('page.audit.path'),
    align: 'center',
    minWidth: 100,
    ellipsis: {
      tooltip: true
    }
  },
  {
    key: 'requestBody',
    title: $t('page.audit.requestBody'),
    align: 'center',
    minWidth: 100,
    ellipsis: {
      tooltip: true
    }
  },
  {
    key: 'responseBody',
    title: $t('page.audit.responseBody'),
    align: 'center',
    minWidth: 100,
    ellipsis: {
      tooltip: true
    }
  }
]);

const loading = ref(false);
const data: Ref<any[]> = ref([]);
async function getData() {
  loading.value = true;

  const { data: res, error } = await getAuditSettings({
    username: listParams.value.username,
    loginname: listParams.value.loginname,
    startTime: listParams.value.timeRange && listParams.value.timeRange[0] ? listParams.value.timeRange[0] / 1000 : 0,
    endTime:
      listParams.value.timeRange && listParams.value.timeRange[1]
        ? Math.floor(new Date(listParams.value.timeRange[1]).setHours(23, 59, 59, 999) / 1000)
        : 0,
    page: listParams.value.page - 1,
    limit: listParams.value.limit,
    state: listParams.value.state,
    method: listParams.value.method,
    path: listParams.value.path
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
    <NGrid x-gap="12" y-gap="12" cols="24" item-responsive>
      <NGi span="24">
        <div class="min-h-500px">
          <UserSearchVue
            v-model:model="listParams"
            :reset-key="resetKey"
            @reset="resetSearchParams"
            @search="getData"
          />
          <NCard
            :title="$t('route.audit_settings')"
            :bordered="false"
            size="small"
            class="sm:flex-1-hidden card-wrapper"
          >
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
