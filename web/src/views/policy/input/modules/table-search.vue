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
  protocol: string;
  port: string;
  icmp: string;
  ctState: string;
  ip: string;
  policy: string;
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
// async function reset() {
//   emit('reset');
// }

const search = debounce(() => {
  // model.value.page = 1;
  emit('search');
});
// async function search() {
//   emit('search');
// }
// 0.启用+有效期 1.启用 2.有效期 3.禁用 4.全部 中的一个" dc:"1.启用 2.有效期 3.禁用  0.启用+有效期  4.全部
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
          :label="$t('page.firewallPolicy.protocol')"
          path="protocol"
          class="pr-24px"
        >
          <NSelect
            v-model:value="model.protocol"
            clearable
            :options="[
              {
                label: 'tcp',
                value: 'tcp'
              },
              {
                label: 'udp',
                value: 'udp'
              },
              {
                label: 'icmp',
                value: 'icmp'
              },
              {
                label: 'ct state',
                value: 'ct'
              }
            ]"
            @update:value="search"
          />
        </NFormItemGi>

        <NFormItemGi span="24 s:24 m:12 l:8" :label="$t('page.firewallPolicy.sourceIp')" path="limitIp" class="pr-24px">
          <NInput v-model:value="model.ip" clearable @update:value="search" />
        </NFormItemGi>

        <NFormItemGi span="24 s:24 m:12 l:8" :label="$t('page.firewallPolicy.action')" path="policy" class="pr-24px">
          <NSelect
            v-model:value="model.policy"
            clearable
            :options="[
              {
                label: $t('page.firewallPolicy.accept'),
                value: 'accept'
              },
              {
                label: $t('page.firewallPolicy.reject'),
                value: 'drop'
              }
            ]"
            @update:value="search"
          />
        </NFormItemGi>

        <NFormItemGi
          v-if="model.protocol === 'tcp' || model.protocol === 'udp'"
          span="24 s:24 m:12 l:8"
          :label="$t('page.firewallPolicy.port')"
          path="port"
          class="pr-24px"
        >
          <NInput v-model:value="model.port" clearable :allow-input="onlyAllowNumber" @update:value="search" />
        </NFormItemGi>

        <NFormItemGi
          v-if="model.protocol === 'icmp'"
          span="24 s:24 m:12 l:8"
          :label="$t('page.firewallPolicy.option')"
          path="port"
          class="pr-24px"
        >
          <NSelect
            v-model:value="model.icmp"
            clearable
            :options="[
              {
                label: $t('page.firewallPolicy.pingReply'),
                value: 'echo-reply'
              },
              {
                label: $t('page.firewallPolicy.pingRequest'),
                value: 'echo-request'
              }
            ]"
            @update:value="search"
          />
        </NFormItemGi>

        <NFormItemGi
          v-if="model.protocol === 'ct'"
          span="24 s:24 m:12 l:8"
          :label="$t('page.firewallPolicy.option')"
          path="port"
          class="pr-24px"
        >
          <NSelect
            v-model:value="model.ctState"
            clearable
            :options="[
              {
                label: $t('page.firewallPolicy.newTcp'),
                value: 'new'
              },
              {
                label: $t('page.firewallPolicy.establishedTcp'),
                value: 'established'
              },
              {
                label: $t('page.firewallPolicy.relatedTcp'),
                value: 'related'
              },
              {
                label: $t('page.firewallPolicy.untrackedTcp'),
                value: 'untracked'
              },
              {
                label: $t('page.firewallPolicy.invalidTcp'),
                value: 'invalid'
              }
            ]"
            @update:value="search"
          />
        </NFormItemGi>
      </NGrid>
    </NForm>
  </NCard>
</template>

<style scoped></style>
