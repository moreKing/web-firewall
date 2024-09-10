<script setup lang="ts">
import { $t } from '@/locales';
import { debounce } from '@/utils/debounce';

defineOptions({
  name: 'UserSearch'
});

interface Emits {
  (e: 'reset'): void;
  (e: 'search'): void;
}

const emit = defineEmits<Emits>();

interface SearchParams {
  sourceIp: string;
  dport: string;
  destIp: string;
}

const model = defineModel<SearchParams>('model', { required: true });

// const

//  防抖
const reset = debounce(
  () => {
    emit('reset');
  },
  300,
  true
);

const search = debounce(() => {
  emit('search');
});

const onlyAllowNumber = (value: string) => {
  if (value === '') return true;
  if (value || /^\d+$/.test(value)) {
    const tmp = Number.parseInt(value, 10);
    return tmp >= 0 && tmp < 65536;
  }

  return false;
};
</script>

<template>
  <NCard :title="$t('common.search')" :bordered="false" size="small" class="card-wrapper">
    <template #header-extra>
      <NButton quaternary @click="reset">
        <template #icon>
          <icon-ic-round-refresh class="text-icon" />
        </template>
        {{ $t('common.reset') }}
      </NButton>
    </template>

    <NForm :model="model" label-placement="left" :label-width="120">
      <NGrid responsive="screen" item-responsive>
        <NFormItemGi
          span="24 s:24 m:12 l:8"
          :label="$t('page.firewallPolicy.sourceIp')"
          path="sourceIp"
          class="pr-24px"
        >
          <NInput v-model:value="model.sourceIp" clearable @update:value="search" />
        </NFormItemGi>

        <NFormItemGi span="24 s:24 m:12 l:8" :label="$t('page.firewallPolicy.destIp')" path="destIp" class="pr-24px">
          <NInput v-model:value="model.destIp" clearable @update:value="search" />
        </NFormItemGi>

        <NFormItemGi span="24 s:24 m:12 l:8" :label="$t('page.firewallPolicy.port')" path="dport" class="pr-24px">
          <NInput v-model:value="model.dport" clearable :allow-input="onlyAllowNumber" @update:value="search" />
        </NFormItemGi>
      </NGrid>
    </NForm>
  </NCard>
</template>

<style scoped></style>
