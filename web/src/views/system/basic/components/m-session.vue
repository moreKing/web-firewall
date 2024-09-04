<script setup lang="ts">
import { ref } from 'vue';
import type { FormInst } from 'naive-ui';
import { getWebConf, setWebConf } from '@/service/api';
import { $t } from '@/locales';
import { SET_WEB_CONF } from '@/utils/permissions_consts';

const loading = ref(false);

const formValue = ref({
  timeout: undefined
});

const rules: any = {
  timeout: [
    {
      type: 'number',
      min: 10,
      max: 1440,
      message: $t('form.webTimeout.invalid', { range: '10-1440' }),
      trigger: 'input'
    }
  ]
};
const server1Ref = ref<FormInst | null>(null);

async function onSubmit() {
  await server1Ref.value?.validate();
  loading.value = true;
  const { error } = await setWebConf(formValue.value);
  loading.value = false;
  if (error) return;
  window.$message?.success($t('common.updateSuccess'));
}

async function initData() {
  loading.value = true;
  const { data, error } = await getWebConf();
  loading.value = false;
  if (error) return;
  formValue.value = data;
}

initData();
</script>

<template>
  <NGrid cols="1 s:1 m:2 l:2 xl:3 2xl:3" item-responsive responsive="screen">
    <NGridItem>
      <NSpin :show="loading">
        <NForm
          ref="server1Ref"
          :model="formValue"
          label-width="140px"
          label-placement="left"
          label-align="left"
          :rules="rules"
          class="ml-20px mt-20px"
        >
          <NFormItem :label="$t('page.basic.webTimeout')" path="timeout">
            <NInputNumber v-model:value="formValue.timeout" :min="10" :max="1440" class="w-full" />
          </NFormItem>

          <NFormItem v-permission="SET_WEB_CONF" label-width="0px" class="mt-40px">
            <NButton type="primary" @click="onSubmit">
              {{ $t('page.userSetting.submit') }}
            </NButton>
          </NFormItem>
        </NForm>
      </NSpin>
    </NGridItem>
  </NGrid>
</template>
