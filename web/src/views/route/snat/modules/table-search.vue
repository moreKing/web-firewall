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

const props = defineProps<{
  network: any;
}>();
const networkOptions = computed(() => {
  return props.network.map((item: any) => {
    return {
      label: item.name,
      value: item.name
    };
  });
});
const emit = defineEmits<Emits>();

interface SearchParams {
  network: string;
  sourceIp: string;
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

//   emit('search');
// }
// 0.启用+有效期 1.启用 2.有效期 3.禁用 4.全部 中的一个" dc:"1.启用 2.有效期 3.禁用  0.启用+有效期  4.全部
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

        <NFormItemGi
          span="24 s:24 m:12 l:8"
          :label="$t('page.firewallPolicy.destinationEthernet')"
          path="policy"
          class="pr-24px"
        >
          <NSelect v-model:value="model.network" clearable :options="networkOptions" @update:value="search" />
        </NFormItemGi>
      </NGrid>
    </NForm>
  </NCard>
</template>

<style scoped></style>
