<script setup lang="ts">
import type { Ref } from 'vue';
import { computed, ref } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import type { FormInst } from 'naive-ui';
import { $t } from '@/locales';
// import { AddOutline, RemoveOutline } from '@vicons/ionicons5';
import { getMessageConf, putMessageConf, testMessage } from '@/service/api';
import { PUT_MESSAGE_CONF, TEST_MESSAGE } from '@/utils/permissions_consts';

const message = useMessage();
const dialog = useDialog();

const formValue = ref({
  state: 0,
  url: '',
  method: 'POST',
  encType: 'json',
  parameters: [],
  content: '特权账号验证码：{code}, {validity}分钟内有效，请勿告知他人。'
});

const loading = ref(false);

const formRef = ref<FormInst | null>(null);

const onSubmit = async () => {
  await formRef.value?.validate();
  loading.value = true;
  const { error } = await putMessageConf(formValue.value);
  loading.value = false;
  if (error) return;
  message.success($t('common.updateSuccess'));
  initData();
};

async function initData() {
  loading.value = true;
  const { data, error } = await getMessageConf();
  loading.value = false;
  if (error) return;
  if (!data.parameters) {
    data.parameters = [];
  }
  formValue.value = data;
}

const rules: Ref<any> = computed(() => {
  return {
    url: [
      {
        type: 'url',
        required: true,
        message: $t('form.url.invalid'),
        trigger: 'input'
      }
    ],
    method: [
      {
        type: 'enum',
        enum: ['GET', 'POST', 'PUT', 'PATCH'],
        required: true,
        message: $t('form.range.invalid', { range: 'GET,POST,PUT,PATCH' }),
        trigger: 'input'
      }
    ],
    encType: [
      {
        type: 'enum',
        enum: ['json', 'form-data'],
        required: true,
        message: $t('form.range.invalid', { range: 'json,form-data' }),
        trigger: 'input'
      }
    ],

    content: [
      {
        required: true,
        message: $t('form.range.required'),
        trigger: 'input'
      }
    ],
    phone: [
      {
        required: true,
        message: $t('form.phone.invalid'),
        trigger: 'input',
        validator(_rule: any, value: string) {
          //   符合短信格式吗
          if (value === '') {
            return false;
          }
          if (!validatePhone(value)) {
            return false;
          }

          return true;
        }
      }
    ]
  };
});

const patternMobile = /^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$/;
function validatePhone(value: string) {
  if (!patternMobile.test(value)) {
    return false;
  }
  return true;
}

//   邮件测试
const showEmailTest = ref(false);
const testEmailFormValue = ref({ phone: '' });
const testEmailLoading = ref(false);
// 邮件测试
const onEmailTest = () => {
  testEmailFormValue.value = { phone: '' };
  testEmailLoading.value = false;
  showEmailTest.value = true;
};

async function testEmailClick() {
  await formRef.value?.validate();
  testEmailLoading.value = true;
  const { error } = await testMessage(testEmailFormValue.value.phone);
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
  <NSpace vertical :size="14" class="mb-30px ml-20px mt-10px font-size-14px text-truegray-400">
    <div>{{ $t('page.basic.parameterDescription') }}</div>
    <div>
      {{
        $t('page.basic.messageDescription', {
          mobile: ' {mobile}&nbsp;&nbsp;',
          content: ' {content}&nbsp;&nbsp;',
          validity: ' {validity}&nbsp;&nbsp;',
          code: ' {code}'
        })
      }}
    </div>
  </NSpace>

  <NGrid cols="1 s:1 m:2 l:2 xl:3 2xl:3" item-responsive responsive="screen">
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
          <NFormItem :label="$t('datatable.status')" path="state">
            <NRadioGroup v-model:value="formValue.state">
              <NSpace>
                <NRadio key="1" :value="0">{{ $t('page.manage.common.status.disable') }}</NRadio>
                <NRadio key="2" :value="1">{{ $t('page.basic.builtIn') }}</NRadio>
                <NRadio key="3" :value="2">Python</NRadio>
              </NSpace>
            </NRadioGroup>
          </NFormItem>
          <div v-if="formValue.state === 1">
            <NFormItem :label="$t('page.basic.requestUrl')" path="url">
              <NInput v-model:value="formValue.url" />
            </NFormItem>

            <NFormItem :label="$t('page.basic.requestMethod')" path="method">
              <NSelect
                v-model:value="formValue.method"
                :options="[
                  {
                    label: 'GET',
                    value: 'GET'
                  },
                  {
                    label: 'POST',
                    value: 'POST'
                  },
                  {
                    label: 'PUT',
                    value: 'PUT'
                  },
                  {
                    label: 'PATCH',
                    value: 'PATCH'
                  }
                ]"
              />
            </NFormItem>

            <NFormItem v-if="formValue.method !== 'GET'" :label="$t('page.basic.parameterType')" path="encType">
              <NRadioGroup v-model:value="formValue.encType">
                <NSpace>
                  <NRadio key="json" value="json">json</NRadio>
                  <NRadio key="form-data" value="form-data">form-data</NRadio>
                </NSpace>
              </NRadioGroup>
            </NFormItem>

            <NFormItem :label="$t('page.basic.requestParameters')" path="parameters">
              <!-- <n-input v-model:value="formValue.parameters" type="textarea" /> -->
              <NDynamicInput
                v-model:value="formValue.parameters"
                preset="pair"
                :key-placeholder="$t('form.key')"
                :value-placeholder="$t('form.value')"
              >
                <template #action="{ index, create, remove }">
                  <NSpace class="ml-20px">
                    <NButton strong secondary type="error" @click="() => remove(index)">
                      <NIcon>
                        <icon-carbon:subtract-large />
                      </NIcon>
                    </NButton>
                    <NButton strong secondary type="success" @click="() => create(index)">
                      <NIcon>
                        <icon-carbon:add-large />
                      </NIcon>
                    </NButton>
                  </NSpace>
                </template>
              </NDynamicInput>
            </NFormItem>

            <NFormItem :label="$t('page.basic.messageContent')" path="content">
              <NInput
                v-model:value="formValue.content"
                type="textarea"
                placeholder="特权账号验证码：{code}, {validity}分钟内有效，请勿告知他人。"
              />
            </NFormItem>
          </div>

          <NFormItem label-width="0px" class="mt-40px">
            <NSpace>
              <NButton v-permission="TEST_MESSAGE" @click="onEmailTest">
                {{ $t('page.basic.test') }}
              </NButton>
              <NButton v-permission="PUT_MESSAGE_CONF" type="primary" @click="onSubmit">
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
    :title="$t('page.basic.test')"
    class="w-600px"
    :mask-closable="false"
  >
    <NSpin :show="testEmailLoading">
      <NForm
        ref="formRef"
        :label-width="80"
        :rules="rules"
        :model="testEmailFormValue"
        label-placement="left"
        label-align="left"
      >
        <NFormItem :label="$t('datatable.phone')" path="phone">
          <NInput v-model:value="testEmailFormValue.phone" />
        </NFormItem>
      </NForm>
    </NSpin>

    <template #footer>
      <NSpace justify="end">
        <NButton @click="showEmailTest = false">{{ $t('common.cancel') }}</NButton>
        <NButton type="primary" :loading="testEmailLoading" @click="testEmailClick">
          {{ $t('page.userSetting.submit') }}
        </NButton>
      </NSpace>
    </template>
  </NModal>
</template>

<!--
 <style scoped>
  :deep(.n-input-number .el-input__inner) {
    text-align: left;
  }
</style>
-->
