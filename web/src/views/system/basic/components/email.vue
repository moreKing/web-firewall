<script setup lang="ts">
import { computed, ref } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import type { FormInst } from 'naive-ui';
import { emailTest, getEmailConf, putEmailConf } from '@/service/api';
import { EMAIL_TEST, PUT_EMAIL_CONF } from '@/utils/permissions_consts';
import { $t } from '@/locales';

const message = useMessage();
const dialog = useDialog();

const formValue = ref({
  enable: false,
  smtp: '',
  port: 25,
  email: '',
  account: '',
  protocol: 1,
  password: ''
});

const loading = ref(false);

const showEditPass = ref(true);
const showEditPassCheckbox = ref(false);

const formRef = ref<FormInst | null>(null);

async function onSubmit() {
  await formRef.value?.validate();
  loading.value = true;
  const { error } = await putEmailConf(formValue.value);
  loading.value = false;
  if (error) return;
  message.success($t('common.updateSuccess'));
  initData();
}

function selectProtocolPort(v: number) {
  switch (v) {
    case 1:
      formValue.value.port = 25;
      break;
    case 2:
      formValue.value.port = 465;
      break;
    case 3:
      formValue.value.port = 587;
      break;

    default:
      break;
  }
}

async function initData() {
  loading.value = true;
  const { data, error } = await getEmailConf();
  loading.value = false;
  if (error) return;
  formValue.value = data;
  showEditPass.value = !data.enable;
  showEditPassCheckbox.value = data.enable;
}

const rules = computed<any>(() => {
  return {
    smtp: [
      {
        required: true,
        message: $t('form.required'),
        trigger: 'input'
      }
    ],
    port: [
      {
        type: 'number',
        min: 1,
        required: true,
        max: 65535,
        message: $t('form.webTimeout.invalid', { range: '1-65535' }),
        trigger: 'input'
      }
    ],
    email: [
      {
        type: 'email',
        required: true,
        message: $t('form.email.invalid'),
        trigger: 'input'
      }
    ],
    name: [
      {
        required: true,
        message: $t('form.required'),
        trigger: 'input'
      }
    ],
    protocol: [
      {
        type: 'number',
        min: 1,
        max: 3,
        message: $t('form.webTimeout.invalid', { range: '1,2,3' }),
        trigger: 'input'
      }
    ],
    account: [
      {
        required: true,
        message: $t('form.required'),
        trigger: 'input'
      }
    ],
    password: [
      {
        required: true,
        message: $t('form.required'),
        trigger: 'input'
      }
    ]
  };
});

//   邮件测试
const showEmailTest = ref(false);
const testEmailAddrStatus = ref(undefined);
const testEmailFormValue = ref({ email: '' });
const testEmailFormRef = ref<FormInst | null>(null);
const testEmailLoading = ref(false);
// 邮件测试
const onEmailTest = () => {
  testEmailFormValue.value = { email: '' };
  testEmailAddrStatus.value = undefined;
  testEmailLoading.value = false;
  showEmailTest.value = true;
};

async function testEmailClick() {
  await testEmailFormRef.value?.validate();
  testEmailLoading.value = true;
  const { error } = await emailTest(testEmailFormValue.value.email);
  testEmailLoading.value = false;
  if (error) {
    dialog.error({
      title: $t('common.dialog.error'),
      content: error.response?.data.message,
      positiveText: $t('common.confirm')
    });
    return;
  }
  dialog.success({
    title: $t('common.dialog.success'),
    content: $t('page.login.common.validateSuccess'),
    positiveText: $t('common.confirm')
  });
}

initData();
</script>

<template>
  <NGrid cols="1 s:1 m:1 l:2 xl:2 2xl:3" item-responsive responsive="screen">
    <NGridItem>
      <NSpin :show="loading">
        <NForm
          ref="formRef"
          :model="formValue"
          label-width="120px"
          label-placement="left"
          label-align="left"
          :rules="rules"
          require-asterisk-position="right"
          class="ml-20px mt-20px"
        >
          <NFormItem :label="$t('datatable.status')" path="enable">
            <NRadioGroup v-model:value="formValue.enable">
              <NSpace>
                <NRadio :key="0" :value="false">{{ $t('page.manage.common.status.disable') }}</NRadio>
                <NRadio :key="1" :value="true">{{ $t('page.manage.common.status.enable') }}</NRadio>
              </NSpace>
            </NRadioGroup>
          </NFormItem>

          <div v-if="formValue.enable">
            <NFormItem :label="$t('page.basic.emailServer')" path="smtp">
              <NInput v-model:value="formValue.smtp" placeholder="smtp.qq.com" />
            </NFormItem>

            <NFormItem :label="$t('page.basic.protocol')" path="protocol">
              <NRadioGroup v-model:value="formValue.protocol" @update:value="selectProtocolPort">
                <NRadio :key="1" :value="1">{{ $t('datatable.normal') }}</NRadio>
                <NRadio :key="2" :value="2">SSL/TLS</NRadio>
                <NRadio :key="3" :value="3">STARTTLS</NRadio>
              </NRadioGroup>
            </NFormItem>

            <NFormItem :label="$t('page.basic.emailPort')" path="port">
              <NInputNumber v-model:value="formValue.port" :min="1" :max="65535" class="w-full" />
            </NFormItem>

            <NFormItem :label="$t('page.basic.emailAccount')" path="account">
              <NInput v-model:value="formValue.account" />
            </NFormItem>

            <NFormItem :label="$t('page.basic.sendEmail')" path="email">
              <NInput v-model:value="formValue.email" placeholder="test@qq.com" />
            </NFormItem>

            <NFormItem v-if="showEditPassCheckbox" :label="$t('page.login.resetPwd.title')">
              <NCheckbox v-model:checked="showEditPass" />
            </NFormItem>
            <NFormItem v-if="showEditPass" :label="$t('page.basic.emailPassword')" path="password">
              <NInput v-model:value="formValue.password" type="password" show-password-on="mousedown" />
            </NFormItem>
          </div>

          <NFormItem label-width="0px" class="mt-40px">
            <NSpace>
              <NButton v-permission="EMAIL_TEST" @click="onEmailTest">
                {{ $t('page.basic.test') }}
              </NButton>
              <NButton v-permission="PUT_EMAIL_CONF" type="primary" @click="onSubmit">
                {{ $t('page.userSetting.submit') }}
              </NButton>
            </NSpace>
          </NFormItem>
        </NForm>
      </NSpin>
    </NGridItem>
  </NGrid>

  <NModal
    v-model:show="showEmailTest"
    preset="card"
    :bordered="false"
    :title="$t('page.basic.sendTest')"
    class="w-600px"
    :mask-closable="false"
  >
    <NSpin :show="testEmailLoading">
      <NForm
        ref="testEmailFormRef"
        :label-width="80"
        :model="testEmailFormValue"
        label-placement="left"
        label-align="left"
        :rules="rules"
      >
        <NFormItem :label="$t('page.basic.emailServer')" path="email">
          <NInput v-model:value="testEmailFormValue.email" />
        </NFormItem>
      </NForm>
    </NSpin>

    <template #footer>
      <NSpace justify="end">
        <NButton @click="showEmailTest = false">{{ $t('common.cancel') }}</NButton>
        <NButton type="primary" :loading="testEmailLoading" @click="testEmailClick">
          {{ $t('page.basic.test') }}
        </NButton>
      </NSpace>
    </template>
  </NModal>
</template>
