<script setup lang="ts">
import { computed, ref } from 'vue';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useThemeStore } from '@/store/modules/theme';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { useAuthStore } from '@/store/modules/auth';
import BindTOTPVue from './modules/bind-totp.vue';
const authStore = useAuthStore();

const symbolId = computed(() => {
  const { VITE_ICON_LOCAL_PREFIX: prefix } = import.meta.env;
  const icon = 'login-left';
  return `#${prefix}-${icon}`;
});

const appStore = useAppStore();
const themeStore = useThemeStore();

const leftBgColor = computed(() => {
  return themeStore.darkMode ? '#121212' : '#e9edf7';
});
const rightBgColor = computed(() => {
  return themeStore.darkMode ? '#1C1C12' : '#f8f8f8';
});

const textColor = computed(() => {
  return themeStore.darkMode ? '#ffffffe5' : '#323d6f';
});

//  登录
const { formRef, validate } = useNaiveForm();

interface TotpOption {
  username: string;
  secret: any;
  issuer: any;
}

//  绑定令牌
const showBindVue = ref(false);
const code = ref('');
const totpOption = ref<TotpOption>({
  username: '',
  secret: '',
  issuer: ''
});

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

interface LoginRes {
  authentication: string;
  issuer?: string;
}

async function handleSubmit() {
  totpOption.value = {
    username: model.value.username,
    secret: '',
    issuer: ''
  };

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
        showBindVue.value = true;
        break;
      default:
        code.value = '';
        model.value.token = '';
        break;
    }
  }
}

function resetPassword() {
  window.$message?.error($t('route.404'));
}

function finishCode() {
  showBindVue.value = false;
  handleSubmit();
}
</script>

<template>
  <div class="relative size-full flex-center overflow-hidden">
    <NScrollbar>
      <NGrid cols="12" item-responsive class="h-100vh">
        <NGridItem span="0 1200:0 1300:4" :style="{ backgroundColor: leftBgColor }" class="h-full">
          <div class="pl-12px pr-12px pt-20%">
            <div class="w-full flex items-center justify-center font-size-32px text-primary font-700">
              <SystemLogo class="mr-10px text-64px text-primary lt-sm:text-48px" />
              <span>Web Firewall</span>
            </div>

            <div
              class="mb-20px mt-20px flex items-center justify-center font-size-12 font-900"
              :style="{ color: textColor }"
            >
              {{ $t('system.loginTitle') }}
            </div>

            <div class="mb-30px mt-30px flex items-center justify-center font-size-16px" :style="{ color: textColor }">
              {{ $t('system.loginContent') }}
            </div>

            <div class="mt-40px">
              <!-- <icon-local-login-left class="w-full" /> -->
              <!-- <SvgIcon local-icon="login-left" /> -->
              <template v-if="true">
                <svg aria-hidden="true" class="block h-410px w-full">
                  <use :xlink:href="symbolId" fill="currentColor" />
                </svg>
              </template>
            </div>
          </div>
        </NGridItem>
        <NGridItem span="12 1200:12 1300:8" :style="{ backgroundColor: rightBgColor }">
          <div class="h-full flex items-center justify-center p-20px">
            <NCard class="w-520px rd-20px pb-40px pl-75px pr-75px pt-40px shadow-xl">
              <div class="flex justify-center">
                <SystemLogo class="mr-10px text-64px text-primary lt-sm:text-48px" />
              </div>
              <h2 class="text-color mb-20px mt-28px flex justify-center text-22px font-500">
                {{ $t('system.welcome') }} Web Firewall ! 🎉
              </h2>

              <NForm ref="formRef" :model="model" :rules="rules" size="large" :show-label="false">
                <NFormItem path="username">
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
                <NSpace vertical>
                  <div class="flex-y-center justify-end">
                    <NButton quaternary type="info" @click="resetPassword">
                      {{ $t('page.login.pwdLogin.forgetPassword') }}
                    </NButton>
                  </div>

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
                </NSpace>
              </NForm>

              <NDivider dashed>{{ $t('icon.themeConfig') }}</NDivider>

              <NFlex justify="center">
                <ThemeSchemaSwitch
                  :theme-schema="themeStore.themeScheme"
                  :show-tooltip="false"
                  class="text-20px lt-sm:text-18px"
                  @switch="themeStore.toggleThemeScheme"
                />
                <LangSwitch
                  :lang="appStore.locale"
                  :lang-options="appStore.localeOptions"
                  :show-tooltip="false"
                  @change-lang="appStore.changeLocale"
                />
              </NFlex>
            </NCard>
          </div>
        </NGridItem>
      </NGrid>
    </NScrollbar>

    <!-- 二步认证 -->
    <BindTOTPVue
      v-model:show="showBindVue"
      v-model:model="model"
      :issuer="totpOption.issuer"
      :secret="totpOption.secret"
      :username="totpOption.username"
      @finish="finishCode"
    />
  </div>
</template>

<style scoped>
.shadow-xl {
  --tw-shadow: 0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1) !important;
  --tw-shadow-colored: 0 20px 25px -5px var(--tw-shadow-color), 0 8px 10px -6px var(--tw-shadow-color) !important;
  box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow) !important;
}
.text-color {
  color: var(--n-title-text-color);
}
</style>
