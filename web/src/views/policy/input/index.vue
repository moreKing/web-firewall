<script setup lang="tsx">
import type { Ref } from 'vue';
import { computed, h, ref } from 'vue';
import { NSpace, NTag } from 'naive-ui';
import dayjs from 'dayjs';
import { deleteInputPolicy, getInputPolicyList } from '@/service/api';
import { $t } from '@/locales';
import CustomTableHeader from '@/components/table/custom-table-header.vue';
import { ADD_FIREWALL_POLICY } from '@/utils/permissions_consts';
import { checkIpAddr, checkIpInNet, checkPort } from '@/utils/ip_check';
import UserSearch from './modules/table-search.vue';
import { createAction } from './modules/utils';
import AddData from './modules/add-data.vue';
import UpdateData from './modules/update-data.vue';
import UpdatePosition from './modules/update-position.vue';

// const chain = 1;

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
        const { error } = await deleteInputPolicy(row.id);
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
    width: 64
  },

  {
    key: 'comment',
    title: $t('datatable.description'),
    align: 'center',
    minWidth: 200
  },

  {
    key: 'protocol',
    title: $t('page.firewallPolicy.protocol'),
    align: 'center',
    minWidth: 100
  },

  {
    title: $t('page.firewallPolicy.policy'),
    align: 'center',
    minWidth: 200,
    render(row: any) {
      let v = '';
      if (row.protocol === 'tcp' || row.protocol === 'udp') {
        v = row.port;
      } else if (row.protocol === 'icmp') {
        v = row.icmp;
      } else if (row.protocol === 'ct') {
        v = row.ct;
      }

      return h(
        NSpace,
        {
          justify: 'center',
          wrap: true
        },
        {
          default: () =>
            v.split(',').map((x: string) => {
              const tmp = x;
              return h(
                NTag,
                {
                  bordered: false,
                  type: row.policy === 'accept' ? 'success' : 'error'
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
            if (!row.ip || row.ip.trim() === '') {
              return h(
                NTag,
                {
                  bordered: false,
                  type: row.policy === 'accept' ? 'success' : 'error'
                },
                {
                  default: () => $t('page.firewallPolicy.allIp')
                }
              );
            }

            return row.ip.split(',').map((x: string) => {
              return h(
                NTag,
                {
                  bordered: false,
                  type: row.policy === 'accept' ? 'success' : 'error'
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
    title: $t('page.firewallPolicy.action'),
    align: 'center',
    minWidth: 100,
    render(row: any) {
      switch (row.policy) {
        case 'accept':
          return h(
            NTag,
            {
              type: 'success',
              bordered: false
            },
            {
              default: () => $t('page.firewallPolicy.accept')
            }
          );

        case 'drop':
          return h(
            NTag,
            {
              type: 'error',
              bordered: false
            },
            {
              default: () => $t('page.firewallPolicy.reject')
            }
          );

        default:
          return h(
            NTag,
            {
              type: 'error',
              bordered: false
            },
            {
              default: () => $t('page.firewallPolicy.drop')
            }
          );
      }
    }
  },

  {
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

const filterData = ref<any>([]);
const tableRef = ref<any>(null);

async function getData() {
  loading.value = true;
  const { data: res, error } = await getInputPolicyList();
  loading.value = false;
  if (error) return;

  if (!res.data) {
    data = [];
    pagination.value.itemCount = 0;
    filterData.value = [];
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

//  过滤
function filterDataFn() {
  // 翻页归1
  if (tableRef.value) tableRef.value.page(1);
  loading.value = true;
  //  只显示符合条件的数据
  // eslint-disable-next-line complexity
  filterData.value = data.filter((item: any) => {
    // 判断协议
    if (listParams.value.protocol && listParams.value.protocol !== '') {
      if (listParams.value.protocol !== item.protocol) {
        return false;
      }
      // 判断协议类型
      switch (listParams.value.protocol) {
        case 'tcp':
          // 判断端口
          if (listParams.value.port && listParams.value.port !== '') {
            const ports = item.port.split(',');
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
            const ports = item.port.split(',');
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

        case 'ct':
          // 判断 tcp状态
          if (listParams.value.ctState && listParams.value.ctState !== '') {
            const vs = item.ct.split(',');
            let tmpState = false;
            for (let i = 0; i < vs.length; i += 1) {
              // eslint-disable-next-line max-depth
              if (listParams.value.ctState.trim() === vs[i].trim()) {
                tmpState = true;
                break;
              }

              // eslint-disable-next-line max-depth
            }
            if (!tmpState) return false;
          }
          break;

        case 'icmp':
          // 判断 icmp状态

          if (listParams.value.icmp && listParams.value.icmp !== '') {
            const vs = item.icmp.split(',');
            let tmpState = false;
            for (let i = 0; i < vs.length; i += 1) {
              // eslint-disable-next-line max-depth
              if (listParams.value.icmp.trim() === vs[i].trim()) {
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
    if (listParams.value.ip && listParams.value.ip !== '' && item.ip && item.ip !== '') {
      // 判断ip 是否有效，有效才进行过滤
      if (checkIpAddr(listParams.value.ip)) {
        // 判断是否在范围&& checkIpInNet()
        const tmpIps = item.ip.split(',');

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
      if (listParams.value.policy !== item.policy) {
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
            <NTag :bordered="false" type="error">
              {{ $t('page.firewallPolicy.defaultPolicy') }}: {{ $t('page.firewallPolicy.reject') }}
            </NTag>
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
