<script setup lang="ts">
import { useRouterPush } from '@/hooks/common/router';
import { $t } from '@/locales';
const { routerPushByKey } = useRouterPush();
</script>

<template>
  <div>
    <div class="mt-20px text-center text-24px text-primary">{{ $t('route.406') }}</div>
    <div class="flex-col-center overflow-hidden">
      <div class="flex text-600px text-primary">
        <SvgIcon local-icon="no-license" />
      </div>

      <NButton type="primary" @click="routerPushByKey('system_license')">{{ $t('route.system_license') }}</NButton>
    </div>
  </div>
</template>

<style scoped></style>
