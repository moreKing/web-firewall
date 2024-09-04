<script setup lang="ts">
// import type { DataTableColumn } from 'naive-ui';
import { $t } from '@/locales';

// const pagination: Pagination = reactive({
//   total: 0,
//   page: 1,
//   limit: 10
// });
const props = defineProps<{
  data: Array<any>;
  columns: Array<any>;
}>();

// const tableData = defineModel<any>('data');

const paramas = defineModel<any>('paramas', { required: true });

// const tableColumns = defineModel<DataTableColumn<any>[]>('columns', {
//   required: true
// });

// const emit = defineEmits(['getData']);
const emit = defineEmits<{
  (e: 'getData'): void;
}>();

function paginationUpdatePageSize() {
  paramas.value.page = 1;
  emit('getData');
}

// emit('getData');

// tableData.value =

function clickSelectDevs() {
  paramas.value.selected = !paramas.value.selected;
  paramas.value.page = 1;
  emit('getData');
}

// 清除选中
function clearCheckedRowKeys() {
  paramas.value.selected = false;
  paramas.value.targets = [];
  paramas.value.page = 1;
  emit('getData');
}
</script>

<template>
  <div>
    <NScrollbar x-scrollable trigger="none">
      <NDataTable
        v-model:checked-row-keys="paramas.targets"
        :data="props.data"
        :columns="columns"
        bordered
        :row-key="row => row.id"
      />
    </NScrollbar>

    <NSpace class="mb-20px mt-20px" justify="space-between">
      <div>
        <div v-if="paramas.targets && paramas.targets.length > 0">
          <NTag
            v-if="!paramas.selected"
            class="cursor-pointer"
            closable
            type="info"
            :color="{ color: '#fff', textColor: '#2080f0', borderColor: '#2080f0' }"
            @close="clearCheckedRowKeys"
            @click="clickSelectDevs"
          >
            {{ $t('page.task.pwdRule.selectedNum', { num: paramas.targets.length }) }}
          </NTag>

          <NTag
            v-if="paramas.selected"
            class="cursor-pointer"
            :bordered="false"
            closable
            type="info"
            @close="clearCheckedRowKeys"
            @click="clickSelectDevs"
          >
            {{ $t('page.task.pwdRule.selectedNum', { num: paramas.targets.length }) }}
          </NTag>
        </div>
      </div>

      <NSpace>
        <span class="lh-28px">{{ $t('datatable.itemCount', { total: paramas.total }) }}</span>
        <NPagination
          v-model:page="paramas.page"
          v-model:page-size="paramas.limit"
          :item-count="paramas.total"
          show-quick-jumper
          show-size-picker
          :page-slot="5"
          :display-order="['size-picker', 'pages', 'quick-jumper']"
          :page-sizes="[10, 20, 30, 50, 100]"
          @update:page="emit('getData')"
          @update:page-size="paginationUpdatePageSize"
        />
        <span class="lh-28px">{{ $t('datatable.page') }}</span>
      </NSpace>
    </NSpace>
  </div>
</template>
