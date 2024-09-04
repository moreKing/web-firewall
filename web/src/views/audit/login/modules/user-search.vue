<script setup lang="ts">
import { computed } from 'vue';
import { $t } from '@/locales';
import { debounce } from '@/utils/debounce';

defineOptions({
  name: 'UserSearch'
});

interface Emits {
  (e: 'reset'): void;
  (e: 'search'): void;
}

interface SearchParams {
  page: number;
  limit: number;
  total: number;
  username: string;
  loginname: string;
  timestamp: number;
  timeRange: [number, number] | null;
  state: number;
  online: number;
}

const emit = defineEmits<Emits>();

const model = defineModel<SearchParams>('model', { required: true });

const props = defineProps<{
  resetKey: number;
}>();

//  防抖
const reset = debounce(
  () => {
    emit('reset');
  },
  500,
  true
);
// async function reset() {
//   emit('reset');
// }

const search = debounce(() => {
  model.value.page = 1;
  emit('search');
});

// async function search() {
//   emit('search');
// }
// 0.启用+有效期 1.启用 2.有效期 3.禁用 4.全部 中的一个" dc:"1.启用 2.有效期 3.禁用  0.启用+有效期  4.全部
const statusOptions = computed(() => [
  {
    label: $t('datatable.all'),
    value: 0
  },
  {
    label: $t('common.dialog.success'),
    value: 1
  },
  {
    label: $t('common.dialog.failed'),
    value: 2
  }
]);

const onlineOptions = computed(() => [
  {
    label: $t('datatable.all'),
    value: 0
  },
  {
    label: $t('page.audit.online'),
    value: 1
  },
  {
    label: $t('page.audit.offline'),
    value: 2
  }
]);

// function dateDisabled(ts: number) {
//   return model.value.startTime ? ts <= model.value.startTime : false;
// }

function startDateDisabled(ts: number) {
  return ts > Date.now();
}
</script>

<template>
  <NCard :title="$t('common.search')" :bordered="false" size="small" class="mb-20px card-wrapper">
    <NForm :model="model" label-placement="left" :label-width="120">
      <NGrid responsive="screen" item-responsive>
        <NFormItemGi span="24 s:24 m:12 l:8" :label="$t('datatable.loginname')" path="loginname" class="pr-24px">
          <NInput
            v-model:value="model.loginname"
            :placeholder="$t('form.loginname.required')"
            clearable
            @update:value="search"
          />
        </NFormItemGi>
        <NFormItemGi span="24 s:24 m:12 l:8" :label="$t('datatable.username')" path="username" class="pr-24px">
          <NInput
            v-model:value="model.username"
            :placeholder="$t('page.manage.user.form.nickName')"
            clearable
            @update:value="search"
          />
        </NFormItemGi>

        <NFormItemGi span="24 s:24 m:12 l:8" :label="$t('page.audit.loginResult')" path="state" class="pr-24px">
          <NSelect v-model:value="model.state" :options="statusOptions" @update:value="search" />
        </NFormItemGi>

        <NFormItemGi span="24 s:24 m:12 l:8" :label="$t('page.audit.startTime')" path="username" class="pr-24px">
          <NDatePicker
            :key="props.resetKey"
            v-model:value="model.timeRange"
            type="daterange"
            class="w-full"
            :default-calendar-start-time="Date.now() - 1000 * 3600 * 24 * 30"
            clearable
            :is-date-disabled="startDateDisabled"
            @update:value="search"
          />
        </NFormItemGi>

        <NFormItemGi span="24 s:24 m:12 l:8" :label="$t('page.audit.loginStatus')" path="online" class="pr-24px">
          <NSelect v-model:value="model.online" :options="onlineOptions" @update:value="search" />
        </NFormItemGi>

        <!--
 <NFormItemGi span="24 s:24 m:12 l:8" :label="$t('page.audit.endTime')" path="username" class="pr-24px">
          <NDatePicker
            :key="props.resetKey"
            v-model:value="model.endTime"
            type="date"
            class="w-full"
            :is-date-disabled="dateDisabled"
            clearable
            @update:value="search"
          />
        </NFormItemGi>
-->

        <NFormItemGi span="24 s:24 m:12 l:8" class="pr-24px">
          <NSpace class="w-full" justify="end">
            <NButton @click="reset">
              <template #icon>
                <icon-ic-round-refresh class="text-icon" />
              </template>
              {{ $t('common.reset') }}
            </NButton>
            <NButton type="primary" ghost @click="search">
              <template #icon>
                <icon-ic-round-search class="text-icon" />
              </template>
              {{ $t('common.search') }}
            </NButton>
          </NSpace>
        </NFormItemGi>
      </NGrid>
    </NForm>
  </NCard>
</template>

<style scoped></style>
