<script setup lang="ts">
import { useAuth } from '@/hooks/business/auth';
// import { DELETE_ROLE, UPDATE_ROLE } from '@/utils/permissions_consts';
import { $t } from '@/locales';
import { GET_AUTH_CONF, GET_EMAIL_CONF, GET_MESSAGE_CONF, GET_WEB_CONF } from '@/utils/permissions_consts';
import EmailVue from './components/email.vue';
import SessionVue from './components/m-session.vue';
import NativeVue from './components/m-native.vue';
import Message from './components/message.vue';
import AuthSetting from './components/auth-setting.vue';

const { hasAuth } = useAuth();
</script>

<template>
  <div>
    <NCard shadow="never" class="min-height">
      <NTabs type="line" animated>
        <NTabPane :tab="$t('page.basic.nativePassword')" name="native">
          <NativeVue />
        </NTabPane>

        <NTabPane v-if="hasAuth(GET_WEB_CONF)" :tab="$t('page.basic.sessionSettings')" name="session">
          <SessionVue />
        </NTabPane>
        <NTabPane v-if="hasAuth(GET_EMAIL_CONF)" :tab="$t('page.basic.emailSettings')" name="mail">
          <EmailVue />
        </NTabPane>
        <NTabPane v-if="hasAuth(GET_MESSAGE_CONF)" :tab="$t('page.basic.messageSettings')" name="message">
          <Message />
        </NTabPane>

        <NTabPane v-if="hasAuth(GET_AUTH_CONF)" :tab="$t('page.basic.loginSetting')" name="authSetting">
          <AuthSetting />
        </NTabPane>
      </NTabs>
    </NCard>
  </div>
</template>

<style scoped>
.min-height {
  min-height: calc(100vh - 200px);
}
</style>
