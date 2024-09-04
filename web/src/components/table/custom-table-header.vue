<script setup lang="ts">
// import { VueDraggable } from 'vue-draggable-plus';
import { $t } from '@/locales';
import { debounce } from '@/utils/debounce';
defineOptions({
  name: 'TableHeaderOperation'
});

interface Props {
  itemAlign?: NaiveUI.Align;
  // disabledDelete?: boolean;
  loading?: boolean;
  addPermission: String;
}

defineProps<Props>();

interface Emits {
  (e: 'add'): void;
  // (e: 'delete'): void;
  (e: 'refresh'): void;
}

const emit = defineEmits<Emits>();

// const columns = defineModel<any[]>('columns', {
//   required: true
// });

function add() {
  emit('add');
}

const refresh = debounce(
  () => {
    emit('refresh');
  },
  300,
  true
);
</script>

<template>
  <NSpace :align="itemAlign" wrap justify="end" class="lt-sm:w-200px">
    <slot name="prefix"></slot>
    <slot name="default">
      <NButton v-permission="addPermission" size="small" ghost type="primary" @click="add">
        <template #icon>
          <icon-ic-round-plus class="text-icon" />
        </template>
        {{ $t('common.add') }}
      </NButton>
      <!--
 <NPopconfirm @positive-click="batchDelete">
        <template #trigger>
          <NButton size="small" ghost type="error" :disabled="disabledDelete">
            <template #icon>
              <icon-ic-round-delete class="text-icon" />
            </template>
            {{ $t('common.batchDelete') }}
          </NButton>
        </template>
        {{ $t('common.confirmDelete') }}
      </NPopconfirm>
-->
    </slot>
    <NButton size="small" type="success" ghost @click="refresh">
      <template #icon>
        <icon-mdi-refresh class="text-icon" :class="{ 'animate-spin': loading }" />
      </template>
      {{ $t('common.refresh') }}
    </NButton>
    <!-- <TableColumnSetting v-model:columns="columns" /> -->
    <!--
 <NPopover placement="bottom-end" trigger="click" class="max-h-300px" scrollable>
      <template #trigger>
        <NButton size="small">
          <template #icon>
            <icon-ant-design-setting-outlined class="text-icon" />
          </template>
          {{ $t('common.columnSetting') }}
        </NButton>
      </template>
      <VueDraggable v-model="columns">
        <div
          v-for="item in columns"
          :key="item.key"
          class="h-36px flex-y-center rd-4px hover:(bg-primary bg-opacity-20)"
        >
          <icon-mdi-drag class="mr-8px cursor-move text-icon" />
          <NCheckbox v-model:checked="item.show">
            {{ item.title }}
          </NCheckbox>
        </div>
      </VueDraggable>
    </NPopover>
-->
    <slot name="suffix"></slot>
  </NSpace>
</template>
