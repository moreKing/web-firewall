<script setup lang="ts">
import type { Component } from 'vue';
import { computed } from 'vue';
import { $t } from '@/locales';
import MaterialSymbolsSunny from '~icons/material-symbols/sunny';
import SolarMoonSleepBold from '~icons/solar/moon-sleep-bold';
import MaterialSymbolsHdrAuto from '~icons/material-symbols/hdr-auto';

defineOptions({ name: 'ThemeSchemaSwitch' });

interface Props {
  /** Theme schema */
  themeSchema: UnionKey.ThemeScheme;
}

const props = withDefaults(defineProps<Props>(), {});

interface Emits {
  (e: 'switch'): void;
}

const emit = defineEmits<Emits>();

function handleSwitch() {
  emit('switch');
}

const icons: Record<UnionKey.ThemeScheme, Component> = {
  light: MaterialSymbolsSunny,
  dark: SolarMoonSleepBold,
  auto: MaterialSymbolsHdrAuto
};

const icon = computed(() => icons[props.themeSchema]);
</script>

<template>
  <NTooltip tooltip-placement="bottom">
    <template #trigger>
      <NButton quaternary @click="handleSwitch">
        <div class="flex-center gap-8px">
          <NIcon size="18">
            <component :is="icon" />
          </NIcon>
        </div>
      </NButton>
    </template>
    {{ $t('icon.themeSchema') }}
  </NTooltip>
</template>

<style scoped></style>
