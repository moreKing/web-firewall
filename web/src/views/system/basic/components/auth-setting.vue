<script setup lang="ts">
import type { Ref } from 'vue';
import { computed, ref } from 'vue';
import { useMessage } from 'naive-ui';
import type { FormInst } from 'naive-ui';
import { getAuthConf, putAuthConf } from '@/service/api';
import { $t } from '@/locales';
import { PUT_AUTH_CONF } from '@/utils/permissions_consts';

const loading = ref(false);
const message = useMessage();
const formRef = ref<FormInst | null>(null);

const formValue = ref({
  totpOffset: undefined as any,
  totpIssuer: '',
  messageOffset: undefined as any,
  emailOffset: undefined as any
});

const rules: Ref<any> = computed(() => {
  return {
    totpOffset: [
      {
        required: true,
        type: 'number',
        min: 1,
        max: 30,
        message: $t('form.range.invalid', { range: '1-30' }),
        trigger: 'input'
      }
    ],
    totpIssuer: [
      {
        required: true,
        message: $t('form.required'),
        trigger: 'input'
      }
    ],
    messageOffset: [
      {
        required: true,
        type: 'number',
        min: 1,
        max: 30,
        message: $t('form.range.invalid', { range: '1-30' }),
        trigger: 'input'
      }
    ],
    emailOffset: [
      {
        required: true,
        type: 'number',
        min: 1,
        max: 30,
        message: $t('form.range.invalid', { range: '1-30' }),
        trigger: 'input'
      }
    ]
  };
});

async function onSubmit() {
  await formRef.value?.validate();
  loading.value = true;
  const { error } = await putAuthConf(formValue.value);
  loading.value = false;
  if (error) return;
  message.success($t('common.updateSuccess'));
}

async function initData() {
  loading.value = true;
  const { data, error } = await getAuthConf();
  loading.value = false;
  if (error) return;
  formValue.value = data;
}

initData();
</script>

<template>
  <NGrid cols="1 s:1 m:1 l:2 xl:2 2xl:3" item-responsive responsive="screen">
    <NGridItem>
      <NSpace vertical :size="14" class="mb-30px ml-20px mt-10px font-size-14px text-truegray-400">
        <span>{{ $t('page.basic.totpTip') }}</span>
        <!-- <span>请配置NTP服务或手工校准服务器时间</span> -->
        <span>{{ $t('page.basic.totpDesc') }}</span>
      </NSpace>

      <NSpin :show="loading">
        <NForm
          ref="formRef"
          :model="formValue"
          label-width="180px"
          label-placement="left"
          label-align="left"
          :rules="rules"
          class="ml-20px mt-20px"
        >
          <NFormItem :label="$t('page.basic.totpOffset')" path="totpOffset">
            <NInputNumber v-model:value="formValue.totpOffset" class="w-full" :min="1" :max="30" />
          </NFormItem>

          <NFormItem :label="$t('page.basic.totpIssue')" path="totpIssuer">
            <NInput v-model:value="formValue.totpIssuer" />
          </NFormItem>

          <NFormItem :label="$t('page.basic.emailCodeValidity')" path="emailOffset">
            <NInputNumber v-model:value="formValue.emailOffset" class="w-full" :min="1" :max="30" />
          </NFormItem>

          <NFormItem :label="$t('page.basic.messageCodeValidity')" path="messageOffset">
            <NInputNumber v-model:value="formValue.messageOffset" class="w-full" :min="1" :max="30" />
          </NFormItem>

          <NFormItem label-width="0px" class="mt-40px">
            <NButton v-permission="PUT_AUTH_CONF" type="primary" @click="onSubmit">
              {{ $t('page.userSetting.submit') }}
            </NButton>
          </NFormItem>
        </NForm>
      </NSpin>
    </NGridItem>
  </NGrid>
</template>
