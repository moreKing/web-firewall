<script setup lang="ts">
import { computed, ref } from 'vue';
import { getSystemKernel, setSystemKernel } from '@/service/api';
import { PUT_SYSTEM_KERNEL } from '@/utils/permissions_consts';
import { $t } from '@/locales';

const loading = ref(false);

const formValue = ref({
  forward: false
});

const formRef = ref();

const rules = computed<any>(() => {
  return {
    forward: [
      {
        type: 'bool',

        required: true,
        message: $t('form.required'),
        trigger: 'change'
      }
    ]
  };
});

async function initData() {
  loading.value = true;
  const { data, error } = await getSystemKernel();
  loading.value = false;
  if (error) return;

  formValue.value = data;
}

const onSubmit = async () => {
  await formRef.value?.validate();
  loading.value = true;
  const { error } = await setSystemKernel(formValue.value);
  loading.value = false;
  if (error) return;
  window.$message?.success($t('common.updateSuccess'));
  initData();
};

initData();
</script>

<template>
  <div>
    <NSpin :show="loading">
      <NCard>
        <NForm
          ref="formRef"
          label-placement="left"
          :rules="rules"
          label-width="100px"
          label-align="left"
          require-mark-placement="right-hanging"
        >
          <NFormItem :label="$t('page.kernel.forward')" path="forward">
            <NRadioGroup v-model:value="formValue.forward" name="kernel">
              <NSpace>
                <NRadio :value="true">{{ $t('page.manage.common.status.enable') }}</NRadio>
                <NRadio :value="false">{{ $t('page.manage.common.status.disable') }}</NRadio>
              </NSpace>
            </NRadioGroup>
          </NFormItem>
        </NForm>

        <NButton v-permission="PUT_SYSTEM_KERNEL" type="primary" @click="onSubmit">
          {{ $t('page.userSetting.submit') }}
        </NButton>
      </NCard>
    </NSpin>
  </div>
</template>

<style scoped></style>
