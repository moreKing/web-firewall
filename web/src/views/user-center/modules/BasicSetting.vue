<script lang="ts" setup>
import { computed, reactive, ref } from 'vue';
import { useMessage } from 'naive-ui';
import { $t } from '@/locales';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { PublicSetPersonProfile, fetchGetUserInfo } from '@/service/api';

const { formRef, validate } = useNaiveForm();

const message = useMessage();

const formValue = ref({
  username: '',
  mobile: '',
  email: '',
  bind: false,
  authenticateId: 0
});

const loading = ref(false);
const authenticate = reactive({
  email: false,
  message: false
});
async function initData() {
  loading.value = true;
  const { data, error } = await fetchGetUserInfo();
  loading.value = false;
  if (error) {
    return;
  }
  authenticate.email = data.email;
  authenticate.message = data.message;
  formValue.value = {
    username: data.user?.username,
    mobile: data.user?.mobile,
    email: data.user?.email,
    bind: data.user.totpState,
    authenticateId: data.user.authenticateId
  };
}

initData();

async function formSubmit() {
  await validate();
  loading.value = true;
  const { error } = await PublicSetPersonProfile(formValue.value);
  loading.value = false;
  if (error) {
    return;
  }
  // await setUserInfo(formValue.value);
  message.success($t('common.updateSuccess'));
}

const rules = computed(() => {
  // inside computed to make locale reactive, if not apply i18n, you can define it without computed
  const { formRules, createRequiredRule } = useFormRules();
  const emailReq: App.Global.FormRule[] = [...formRules.email];
  if (authenticate.email && formValue.value.authenticateId === 2) {
    emailReq.push(createRequiredRule($t('form.email.required')));
  }

  const phoneReq: App.Global.FormRule[] = [...formRules.phone];
  if (authenticate.message && formValue.value.authenticateId === 3) {
    phoneReq.push(createRequiredRule($t('form.phone.required')));
  }

  return {
    username: createRequiredRule($t('form.userName.required')),
    mobile: phoneReq,
    email: emailReq
  };
});

const secondOptions = computed(() => {
  const res = [
    { label: $t('page.basic.native'), value: 0 },
    { label: $t('page.basic.mobile'), value: 1 }
  ];

  if (authenticate.email) {
    res.push({ label: $t('page.basic.email'), value: 2 });
  }
  if (authenticate.message) {
    res.push({ label: $t('page.basic.message'), value: 3 });
  }
  // return [
  //   { label: $t('page.basic.mobile'), value: 3 },
  //   { label: $t('page.basic.email'), value: 4 },
  //   { label: $t('page.basic.message'), value: 5 }
  // ];
  return res;
});
</script>

<template>
  <NCard :bordered="false" size="small" :title="$t('page.userSetting.basicSetting')" class="proCard">
    <NGrid cols="2 s:2 m:2 l:3 xl:3 2xl:3" responsive="screen">
      <NGridItem>
        <NSpin :show="loading">
          <NForm ref="formRef" :label-width="80" :model="formValue" :rules="rules">
            <NFormItem :label="$t('page.userSetting.username')" path="username">
              <NInput v-model:value="formValue.username" />
            </NFormItem>

            <NFormItem :label="$t('page.userSetting.email')" path="email">
              <NInput v-model:value="formValue.email" :placeholder="$t('form.email.required')" />
            </NFormItem>

            <NFormItem :label="$t('page.userSetting.phone')" path="mobile">
              <NInput v-model:value="formValue.mobile" :placeholder="$t('form.phone.required')" />
            </NFormItem>

            <NFormItem :label="$t('page.basic.second')" path="authenticateId">
              <NSelect v-model:value="formValue.authenticateId" :options="secondOptions" />
            </NFormItem>

            <NFormItem v-if="formValue.authenticateId === 1">
              <NCheckbox v-model:checked="formValue.bind">{{ $t('datatable.bindTOTP') }}</NCheckbox>
            </NFormItem>

            <div>
              <NSpace>
                <NButton type="primary" @click="formSubmit">{{ $t('page.userSetting.submit') }}</NButton>
              </NSpace>
            </div>
          </NForm>
        </NSpin>
      </NGridItem>
    </NGrid>
  </NCard>
</template>
