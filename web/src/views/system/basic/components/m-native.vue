<script setup lang="ts">
import type { Ref } from 'vue';
import { computed, ref } from 'vue';
import { useMessage } from 'naive-ui';
import type { FormInst } from 'naive-ui';
import { GetPasswordComplex, setPasswordComplex } from '@/service/api';
import { SET_PASSWORD_COMPLEX } from '@/utils/permissions_consts';
import { $t } from '@/locales';

const loading = ref(false);
const message = useMessage();

const formValue = ref({
  length: undefined, // 密码长度
  complex: undefined, // 本地密码复杂度 1. 不限制  2.至少包含数字、字母 3.至少包含字母、数字、特殊字符 4.至少包含大小写字母、数字、特殊字符
  validity: undefined, // 密码有效期：1-不限；具体天数。本地认证密码过期时间,仅支持 1 30 60 90 180 365 选择一个值
  expire: undefined, // 密码过期处理：0-禁止登录；1-仅提醒'
  differTimes: undefined // 不能与前多少次密码一致 1-30
});

const rules: Ref<any> = computed(() => {
  return {
    length: [
      {
        type: 'number',
        min: 6,
        max: 200,
        required: true,
        message: $t('form.webTimeout.invalid', { range: '6-200' }),
        trigger: 'input'
      }
    ],
    complex: [
      {
        type: 'number',
        required: true,
        min: 1,
        max: 4,
        message: $t('form.webTimeout.invalid', { range: '1-4' }),
        trigger: 'input'
      }
    ],
    validity: [
      {
        type: 'enum',
        required: true,
        enum: [1, 30, 60, 90, 180, 365],
        message: $t('form.webTimeout.invalid', { range: '1,30,60,90,180,365' }),
        trigger: 'input'
      }
    ],
    differTimes: [
      {
        type: 'number',
        required: true,
        min: 1,
        max: 30,
        message: $t('form.webTimeout.invalid', { range: '1-30' }),
        trigger: 'input'
      }
    ],
    expire: [
      {
        type: 'enum',
        required: true,
        enum: [0, 1],
        message: $t('form.webTimeout.invalid', { range: '0,1' }),
        trigger: 'input'
      }
    ]
  };
});

const passwordValidTimeOptions = computed(() => [
  { label: $t('page.basic.noLimit'), value: 1 },
  { label: `30 ${$t('datatable.date.day')}`, value: 30 },
  { label: `60 ${$t('datatable.date.day')}`, value: 60 },
  { label: `90 ${$t('datatable.date.day')}`, value: 90 },
  { label: `180 ${$t('datatable.date.day')}`, value: 180 },
  { label: `1 ${$t('datatable.date.year')}`, value: 365 }
]);

// 1. 不限制  2.至少包含数字、字母 3.至少包含字母、数字、特殊字符 4.至少包含大小写字母、数字、特殊字符
const complexOptions = computed(() => [
  { label: $t('page.basic.noLimit'), value: 1 },
  { label: $t('form.userPwd2.invalid'), value: 2 },
  { label: $t('form.userPwd3.invalid'), value: 3 },
  { label: $t('form.userPwd4.invalid'), value: 4 }
]);

const expireOptions = computed(() => {
  return [
    { label: $t('page.basic.disableLogin'), value: 0 },
    { label: $t('page.basic.notify'), value: 1 }
  ];
});

const formValueRef = ref<FormInst | null>(null);
async function onSubmit() {
  await formValueRef.value?.validate();
  loading.value = true;
  const { error } = await setPasswordComplex(formValue.value);
  if (error) return;
  loading.value = false;
  message.success('操作成功');
  initData();
}

async function initData() {
  loading.value = true;
  const { data, error } = await GetPasswordComplex();
  loading.value = false;
  if (error) return;
  formValue.value = data;
}

initData();
</script>

<template>
  <NGrid cols="1 s:1 m:1 l:2 xl:2 2xl:3" item-responsive responsive="screen">
    <NGridItem>
      <NSpin :show="loading">
        <NForm
          ref="formValueRef"
          :model="formValue"
          label-width="160px"
          label-placement="left"
          label-align="left"
          :rules="rules"
          class="ml-20px mt-20px"
        >
          <NFormItem :label="$t('datatable.passwordLength')" path="length">
            <NInputNumber v-model:value="formValue.length" class="w-full" :min="6" :max="20" />
          </NFormItem>

          <NFormItem :label="$t('datatable.passwordComplexity')" path="complex">
            <NSelect v-model:value="formValue.complex" :options="complexOptions" />
          </NFormItem>

          <NFormItem :label="$t('datatable.passwordValid')" path="validity">
            <NSelect v-model:value="formValue.validity" :options="passwordValidTimeOptions" />
          </NFormItem>

          <NFormItem :label="$t('datatable.expiredAction')" path="expire">
            <NSelect v-model:value="formValue.expire" :options="expireOptions" />
          </NFormItem>

          <NFormItem :label="$t('datatable.differChecked')" path="differTimes">
            <NInputNumber v-model:value="formValue.differTimes" class="w-full" :min="1" :max="30" />
          </NFormItem>

          <NFormItem v-permission="SET_PASSWORD_COMPLEX" label-width="0px" class="mt-40px">
            <NButton type="primary" @click="onSubmit">{{ $t('page.userSetting.submit') }}</NButton>
          </NFormItem>
        </NForm>
      </NSpin>
    </NGridItem>
  </NGrid>
</template>
