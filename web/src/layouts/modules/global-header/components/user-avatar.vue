<script setup lang="ts">
import { computed, h } from 'vue';
import type { VNode } from 'vue';
import { NIcon } from 'naive-ui';
import { useAuthStore } from '@/store/modules/auth';
import { useRouterPush } from '@/hooks/common/router';
// import { useSvgIcon } from '@/hooks/common/icon';
import { $t } from '@/locales';
import { PublicLogout } from '@/service/api';
import PhUserCircle from '~icons/ph/user-circle';
import PrimeSignOut from '~icons/prime/sign-out';

defineOptions({
  name: 'UserAvatar'
});

const authStore = useAuthStore();
const { routerPushByKey, toLogin } = useRouterPush();
// const { SvgIconVNode } = useSvgIcon();

function loginOrRegister() {
  toLogin();
}

type DropdownKey = 'user-center' | 'logout';

type DropdownOption =
  | {
      key: DropdownKey;
      label: string;
      icon?: () => VNode;
    }
  | {
      type: 'divider';
      key: string;
    };

const options = computed(() => {
  const opts: DropdownOption[] = [
    {
      label: $t('common.userCenter'),
      key: 'user-center',
      icon: () => h(NIcon, { size: 20 }, { default: () => h(PhUserCircle, {}) })
    },
    {
      type: 'divider',
      key: 'divider'
    },
    {
      label: $t('common.logout'),
      key: 'logout',
      icon: () => h(NIcon, { size: 20 }, { default: () => h(PrimeSignOut, {}) })
    }
  ];

  return opts;
});

function logout() {
  window.$dialog?.info({
    title: $t('common.tip'),
    content: $t('common.logoutConfirm'),
    positiveText: $t('common.confirm'),
    negativeText: $t('common.cancel'),
    onPositiveClick: async () => {
      await PublicLogout();
      authStore.resetStore();
    }
  });
}

function handleDropdown(key: DropdownKey) {
  if (key === 'logout') {
    logout();
  } else {
    // If your other options are jumps from other routes, they will be directly supported here
    routerPushByKey(key);
  }
}
</script>

<template>
  <NButton v-if="!authStore.isLogin" quaternary @click="loginOrRegister">
    {{ $t('page.login.common.loginOrRegister') }}
  </NButton>

  <NDropdown v-else placement="bottom" trigger="click" :options="options" @select="handleDropdown">
    <div>
      <ButtonIcon>
        <icon-ph:user-circle class="text-icon-large" />
        <span class="text-16px font-medium">{{ authStore.userInfo.user.username }}</span>
      </ButtonIcon>
    </div>
  </NDropdown>
</template>

<style scoped></style>
