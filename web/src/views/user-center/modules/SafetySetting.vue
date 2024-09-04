<script lang="ts" setup>
import { computed, reactive, ref } from 'vue';
import { GetPasswordComplex, PublicSetPassword } from '@/service/api';
import { $t } from '@/locales';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';

interface FormModel {
  oldPassword: string;
  newPassword: string;
  confirmPassword: string;
}

const formValue: FormModel = reactive({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
});

const passwordComplex = ref({
  complex: 1,
  length: 6
});

const rules = computed(() => {
  // inside computed to make locale reactive, if not apply i18n, you can define it without computed
  const { formRules, createUserPassword, createUserConfirmPassword } = useFormRules();

  return {
    oldPassword: formRules.loginPwd,
    newPassword: createUserPassword(passwordComplex.value.length, passwordComplex.value.complex),
    confirmPassword: createUserConfirmPassword(formValue.newPassword)
  };
});

const { formRef, validate } = useNaiveForm();

function initFormValue() {
  formValue.confirmPassword = '';
  formValue.newPassword = '';
  formValue.oldPassword = '';
}
const loading = ref(false);

// 提交修改
async function formSubmit() {
  await validate();
  loading.value = true;
  const { error } = await PublicSetPassword({
    oldPassword: formValue.oldPassword,
    newPassword: formValue.newPassword
  });
  loading.value = false;
  if (error) return;

  initFormValue();
  window.$message?.success('success');
}

async function initData() {
  loading.value = true;
  try {
    const { data } = await GetPasswordComplex();
    passwordComplex.value = data;
  } finally {
    loading.value = false;
  }
}

initData();
</script>

<template>
  <NCard :bordered="false" size="small" :title="$t('page.userSetting.securitySetting')" class="proCard">
    <NGrid cols="2 s:2 m:2 l:3 xl:3 2xl:3" responsive="screen">
      <NGridItem>
        <NSpin :show="loading">
          <NForm ref="formRef" :label-width="80" :model="formValue" :rules="rules">
            <NFormItem :label="$t('page.userSetting.oldPassword')" path="oldPassword">
              <NInput v-model:value="formValue.oldPassword" type="password" />
            </NFormItem>

            <NFormItem :label="$t('page.userSetting.newPassword')" path="newPassword">
              <NInput v-model:value="formValue.newPassword" type="password" />
            </NFormItem>

            <NFormItem :label="$t('page.userSetting.confirmPassword')" path="confirmPassword">
              <NInput v-model:value="formValue.confirmPassword" type="password" />
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
