<script setup lang="tsx">
import type { Ref } from 'vue';
import { computed, h, ref } from 'vue';
import { NSpace, NTag } from 'naive-ui';
import dayjs from 'dayjs';
import { deleteFirewallPolicy, getFirewallPolicyList } from '@/service/api';
import { $t } from '@/locales';
import CustomTableHeader from '@/components/table/custom-table-header.vue';
import { ADD_FIREWALL_POLICY } from '@/utils/permissions_consts';
import { checkIpAddr, checkIpInNet, checkPort } from '@/utils/ip_check';
import UserSearch from './modules/table-search.vue';
import { createAction } from './modules/utils';
import AddData from './modules/add-data.vue';
import UpdateData from './modules/update-data.vue';
import UpdatePosition from './modules/update-position.vue';

const chain = 5;
interface SearchParams {
  protocol: string;
  port: string;
  icmp: string;
  ctState: string;
  ip: string;
  policy: string;
}

const listParams: Ref<SearchParams> = ref({
  protocol: '',
  port: '',
  icmp: '',
  ctState: '',
  ip: '',
  policy: ''
});

function resetSearchParams() {
  listParams.value = {
    protocol: '',
    port: '',
    icmp: '',
    ctState: '',
    ip: '',
    policy: ''
  };
  getData();
}
let data: any[] = [];
const loading = ref(false);
const showUpdateData = ref(false);
const updateData = ref<any>({});
const showUpdateDataPosition = ref(false);

function updateModalClick(row: any, method: string) {
  if (method === 'update') {
    updateData.value = row;
    showUpdateData.value = true;
  } else if (method === 'delete') {
    window.$dialog?.warning({
      title: $t('common.dialog.warning'),
      content: () =>
        h('p', null, [
          h('span', null, { default: () => $t('common.confirmDelete') })
          // h('p', { style: 'color: red;margin-top:10px' }, { default: () => row.loginname })
        ]),
      positiveText: $t('common.confirm'),
      negativeText: $t('common.cancel'),
      onPositiveClick: async () => {
        const { error } = await deleteFirewallPolicy(row.id);
        if (error) return;
        getData();
        window.$message?.success($t('common.deleteSuccess'));
      }
    });
  } else if (method === 'position') {
    updateData.value = row;
    showUpdateDataPosition.value = true;
  }
}

const columns = computed<any>(() => [
  {
    key: 'id',
    title: 'id',
    align: 'center',
    show: true,
    width: 64
  },

  {
    show: true,
    key: 'comment',
    title: $t('datatable.description'),
    align: 'center',
    minWidth: 200
  },

  {
    show: true,
    key: 'expr[1].protocol',
    title: $t('page.firewallPolicy.protocol'),
    align: 'center',
    minWidth: 100,
    render(row: any) {
      return h('span', null, {
        default: () =>
          row.expr[1].value && row.expr[1].value.trim() !== '' ? row.expr[1].protocol : $t('page.firewallPolicy.all')
      });
    }
  },

  {
    show: true,
    key: 'expr[1].value',
    title: $t('page.firewallPolicy.port'),
    align: 'center',
    minWidth: 200,
    render(row: any) {
      return h(
        NSpace,
        {
          justify: 'center',
          wrap: true
        },
        {
          default: () =>
            row.expr[1].value.split(',').map((x: string) => {
              let tmp = x;
              if (row.expr[1].field === 'vmap') {
                tmp = x.split(':')[0];
              }
              if (tmp.trim() === '') return null;
              return h(
                NTag,
                {
                  bordered: false,
                  type: row.expr[2].policy === 'accept' ? 'success' : 'error'
                },
                {
                  default: () => tmp
                }
              );
            })
        }
      );
    }
  },

  {
    show: true,
    key: 'expr[0].value',
    align: 'center',
    title: $t('page.firewallPolicy.sourceIp'),
    minWidth: 200,
    render(row: any) {
      return h(
        NSpace,
        {
          justify: 'center',
          wrap: true
        },
        {
          default: () => {
            if (!row.expr[0].value || row.expr[0].value.trim() === '') {
              return h(
                NTag,
                {
                  bordered: false,
                  type: row.expr[2].policy === 'accept' ? 'success' : 'error'
                },
                {
                  default: () => $t('page.firewallPolicy.allIp')
                }
              );
            }

            return row.expr[0].value.split(',').map((x: string) => {
              return h(
                NTag,
                {
                  bordered: false,
                  type: row.expr[2].policy === 'accept' ? 'success' : 'error'
                },
                {
                  default: () => x
                }
              );
            });
          }
        }
      );
    }
  },

  {
    show: true,
    key: 'expr[2].value',
    title: $t('page.firewallPolicy.speed'),
    align: 'center',
    minWidth: 100
  },

  {
    show: true,
    title: $t('datatable.createTime'),
    key: 'createdAt',
    align: 'center',
    width: 200,
    render(row: any) {
      return h('span', null, {
        // format(row.createdAt, 'yyyy-MM-dd HH:mm:ss')
        default: () =>
          row.createdAt && row.createdAt > 0 ? dayjs.unix(row.createdAt).format('YYYY-MM-DD HH:mm:ss') : ''
      });
    }
  },

  {
    show: true,
    title: $t('datatable.action'),
    key: 'actions',
    align: 'center',
    width: 200,
    render: (row: any) => createAction(row, updateModalClick)
  }
]);

const pagination = ref<any>({
  page: 1,
  pageSize: 20,
  showSizePicker: true,
  displayOrder: ['size-picker', 'pages', 'quick-jumper'],
  itemCount: 0,
  showQuickJumper: true,
  pageSizes: [10, 20, 50, 100],
  onChange: (page: number) => {
    pagination.value.page = page;
  },
  onUpdatePageSize: (pageSize: number) => {
    pagination.value.pageSize = pageSize;
    pagination.value.page = 1;
  },
  prefix({ itemCount }: { itemCount: number }) {
    return $t('datatable.itemCount', { total: itemCount });
  }
});

async function getData() {
  loading.value = true;
  const { data: res, error } = await getFirewallPolicyList(chain);
  loading.value = false;
  if (error) return;

  if (!res.data) {
    data = [];
    pagination.value.itemCount = 0;
    return;
  }
  data = res.data;
  filterDataFn();
  // pagination.value.itemCount = res.total;
}

getData();

const showAddData = ref(false);
//
function handleAdd() {
  showAddData.value = true;
}

const filterData = ref<any>([]);
const tableRef = ref<any>(null);

//  过滤
function filterDataFn() {
  // 翻页归1
  if (tableRef.value) tableRef.value.page(1);
  loading.value = true;
  //  只显示符合条件的数据
  // eslint-disable-next-line complexity
  filterData.value = data.filter((item: any) => {
    // 判断协议
    if (listParams.value.protocol && listParams.value.protocol !== '' && item.expr[1].value !== '') {
      if (listParams.value.protocol !== item.expr[1].protocol) {
        return false;
      }
      // 判断协议类型
      switch (listParams.value.protocol) {
        case 'tcp':
          // 判断端口
          if (listParams.value.port && listParams.value.port !== '') {
            const ports = item.expr[1].value.split(',');
            let tmpState = false;
            for (let i = 0; i < ports.length; i += 1) {
              // eslint-disable-next-line max-depth
              if (checkPort(listParams.value.port, ports[i])) {
                tmpState = true;
                break;
              }
            }

            // eslint-disable-next-line max-depth
            if (!tmpState) return false;
          }
          break;

        case 'udp':
          // 判断端口
          if (listParams.value.port && listParams.value.port !== '') {
            const ports = item.expr[1].value.split(',');
            let tmpState = false;
            for (let i = 0; i < ports.length; i += 1) {
              // eslint-disable-next-line max-depth
              if (checkPort(listParams.value.port, ports[i])) {
                tmpState = true;
                break;
              }
            }

            // eslint-disable-next-line max-depth
            if (!tmpState) return false;
          }
          break;

        case 'ct state':
          // 判断 tcp状态
          if (listParams.value.ctState && listParams.value.ctState !== '') {
            const vs = item.expr[1].value.split(',');
            let tmpState = false;
            for (let i = 0; i < vs.length; i += 1) {
              // eslint-disable-next-line max-depth
              if (listParams.value.ctState.trim() === vs[i].split(':')[0].trim()) {
                tmpState = true;
                break;
              }

              // eslint-disable-next-line max-depth
            }
            if (!tmpState) return false;
          }
          break;

        case 'icmp type':
          // 判断 icmp状态

          if (listParams.value.icmp && listParams.value.icmp !== '') {
            const vs = item.expr[1].value.split(',');
            let tmpState = false;
            for (let i = 0; i < vs.length; i += 1) {
              // eslint-disable-next-line max-depth
              if (listParams.value.icmp.trim() === vs[i].split(':')[0].trim()) {
                tmpState = true;
                break;
              }

              // eslint-disable-next-line max-depth
            }
            if (!tmpState) return false;
          }

          break;

        default:
          break;
      }
    }

    // 判断源地址
    if (listParams.value.ip && listParams.value.ip !== '' && item.expr[0].value && item.expr[0].value !== '') {
      // 判断ip 是否有效，有效才进行过滤
      if (checkIpAddr(listParams.value.ip)) {
        // 判断是否在范围&& checkIpInNet()
        const tmpIps = item.expr[0].value.split(',');

        let tmpIpState = false;
        for (let i = 0; i < tmpIps.length; i += 1) {
          // console.log(listParams.value.ip, tmpIps[i], checkIpInNet(listParams.value.ip, tmpIps[i]));
          if (tmpIps[i] && checkIpInNet(listParams.value.ip, tmpIps[i])) {
            tmpIpState = true;
            break;
          }
        }
        if (!tmpIpState) return false;
      }
    }

    // 判断策略
    if (listParams.value.policy && listParams.value.policy !== '') {
      if (listParams.value.policy !== item.expr[2].policy) {
        return false;
      }
    }
    return true;
  });

  loading.value = false;
}
</script>

<template>
  <NSpin :show="loading">
    <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
      <UserSearch v-model:model="listParams" @reset="resetSearchParams" @search="filterDataFn" />
      <NCard :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
        <template #header-extra>
          <CustomTableHeader
            v-model:columns="columns"
            :add-permission="ADD_FIREWALL_POLICY"
            @add="handleAdd"
            @refresh="getData"
          />
        </template>

        <template #header>
          <div class="mb-5px mt-5px">
            {{ $t('page.firewallPolicy.list') }}
          </div>
        </template>
        <NDataTable
          ref="tableRef"
          pagination-behavior-on-filter="first"
          :columns="columns"
          :data="filterData"
          :pagination="pagination"
        />
      </NCard>
    </div>

    <AddData v-model:show="showAddData" @close="getData" />
    <UpdateData v-model:show="showUpdateData" :row="updateData" @close="getData" />
    <UpdatePosition v-model:show="showUpdateDataPosition" :row="updateData" @close="getData" />
  </NSpin>
</template>

<style scoped lang="scss">
:deep(.n-pagination) {
  .n-pagination-prefix {
    margin-right: 10px;
  }
  .n-select {
    margin-right: 10px;
  }
  .n-pagination-quick-jumper {
    margin-left: 20px;
  }
}
</style>
