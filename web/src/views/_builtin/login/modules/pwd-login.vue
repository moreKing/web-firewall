<script setup lang="ts">
import { computed, ref } from 'vue';
import { $t } from '@/locales';
// import { loginModuleRecord } from '@/constants/app';
// import { useRouterPush } from '@/hooks/common/router';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { useAuthStore } from '@/store/modules/auth';
import BindTOTPVue from './bind-totp.vue';
import CodeInputVue from './code-input.vue';

defineOptions({
  name: 'PwdLogin'
});

const authStore = useAuthStore();
// const { toggleLoginModule } = useRouterPush();
const { formRef, validate } = useNaiveForm();

interface FormModel {
  username: string;
  password: string;
  token?: string;
  code?: string;
}

const model = ref<FormModel>({
  username: '',
  password: '',
  token: '',
  code: ''
});

const rules = computed(() => {
  // inside computed to make locale reactive, if not apply i18n, you can define it without computed
  const { formRules, createRequiredRule } = useFormRules();

  return {
    username: createRequiredRule($t('form.loginname.required')),
    password: formRules.loginPwd
  };
});

interface TotpOption {
  username: string;
  secret: any;
  issuer: any;
}

//  绑定令牌
const showBindVue = ref(false);
const totpOption = ref<TotpOption>({
  username: '',
  secret: '',
  issuer: ''
});

// 二步认证
const showInputCode = ref(false);
const code = ref('');

function finishBind() {
  showBindVue.value = false;
  showInputCode.value = true;
}

function finishCode() {
  showInputCode.value = false;
  handleSubmit();
}

interface LoginRes {
  authentication: string;
  issuer?: string;
}

async function handleSubmit() {
  await validate();
  const error = await authStore.login(model.value);
  model.value.code = '';
  model.value.token = '';
  if (error) {
    const res: LoginRes = error.response?.data?.data as LoginRes;
    switch (error?.response?.data.code) {
      case 1:
        //  手机令牌绑定

        totpOption.value = {
          username: model.value.username,
          secret: res.authentication,
          issuer: res.issuer
        };
        showBindVue.value = true;
        break;
      case 2:
        // 输入二步验证码
        // model.value.token = res.authentication;
        if (res && res.authentication) {
          model.value.token = res.authentication;
        }
        showInputCode.value = true;
        break;
      default:
        code.value = '';
        model.value.token = '';
        break;
    }
  }
}

// type AccountKey = 'super' | 'admin' | 'user';

// interface Account {
//   key: AccountKey;
//   label: string;
//   userName: string;
//   password: string;
// }
</script>

<template>
  <NForm ref="formRef" :model="model" :rules="rules" size="large" :show-label="false">
    <NFormItem path="userName">
      <NInput v-model:value="model.username" :placeholder="$t('page.login.common.userNamePlaceholder')" />
    </NFormItem>
    <NFormItem path="password">
      <NInput
        v-model:value="model.password"
        type="password"
        show-password-on="click"
        :placeholder="$t('page.login.common.passwordPlaceholder')"
        @keydown.enter="handleSubmit"
      />
    </NFormItem>
    <NSpace vertical :size="24">
      <!--
 <div class="flex-y-center justify-between">
        <NCheckbox>{{ $t('page.login.pwdLogin.rememberMe') }}</NCheckbox>
        <NButton quaternary @click="toggleLoginModule('reset-pwd')">
          {{ $t('page.login.pwdLogin.forgetPassword') }}
        </NButton>
      </div>
-->
      <NButton
        class="mb-20px mt-20px"
        type="primary"
        size="large"
        round
        block
        :loading="authStore.loginLoading"
        @click="handleSubmit"
      >
        {{ $t('common.confirm') }}
      </NButton>

      <!--
 <div class="flex-y-center justify-between gap-12px">
        <NButton class="flex-1" block @click="toggleLoginModule('code-login')">
          {{ $t(loginModuleRecord['code-login']) }}
        </NButton>
        <NButton class="flex-1" block @click="toggleLoginModule('register')">
          {{ $t(loginModuleRecord.register) }}
        </NButton>
      </div>
-->
      <!--
 <NDivider class="text-14px text-#666 !m-0">{{ $t('page.login.pwdLogin.otherAccountLogin') }}</NDivider>
      <div class="flex-center gap-12px">
        <NButton v-for="item in accounts" :key="item.key" type="primary" @click="handleAccountLogin(item)">
          {{ item.label }}
        </NButton>
      </div>
-->
    </NSpace>
    <BindTOTPVue
      v-model:show="showBindVue"
      :issuer="totpOption.issuer"
      :secret="totpOption.secret"
      :username="totpOption.username"
      @finish="finishBind"
    />

    <CodeInputVue v-model:model="model" v-model:show="showInputCode" @finish="finishCode" />
  </NForm>
</template>

<style scoped></style>
