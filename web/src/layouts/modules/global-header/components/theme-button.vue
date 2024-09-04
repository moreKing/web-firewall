<script setup lang="ts">
import { useAppStore } from '@/store/modules/app';
import { $t } from '@/locales';
import MajesticonsColorSwatchLine from '~icons/majesticons/color-swatch-line';
defineOptions({
  name: 'ThemeButton'
});

const appStore = useAppStore();
</script>

<template>
  <NTooltip>
    <template #trigger>
      <NButton quaternary @click="appStore.openThemeDrawer">
        <div class="flex-center gap-8px">
          <NIcon size="16">
            <MajesticonsColorSwatchLine />
          </NIcon>
        </div>
      </NButton>
    </template>
    {{ $t('icon.themeConfig') }}
  </NTooltip>
</template>

<style scoped></style>
