<script lang="ts" setup>
import type { Component } from 'vue';
import { computed } from 'vue';
import { $t } from '@/locales';
import LineMdMenuFoldLeft from '~icons/line-md/menu-fold-left';
import LineMdMenuFoldRight from '~icons/line-md/menu-fold-right';
import PhCaretDoubleLeftBold from '~icons/ph/caret-double-left-bold';
import PhCaretDoubleRightBold from '~icons/ph/caret-double-right-bold';

defineOptions({ name: 'MenuToggler' });

interface Props {
  /** Show collapsed icon */
  collapsed?: boolean;
  /** Arrow style icon */
  arrowIcon?: boolean;
  zIndex?: number;
}

const props = withDefaults(defineProps<Props>(), {
  arrowIcon: false,
  zIndex: 98
});

type NumberBool = 0 | 1;

const icon = computed(() => {
  const icons: Record<NumberBool, Record<NumberBool, Component>> = {
    0: {
      0: LineMdMenuFoldLeft,
      1: LineMdMenuFoldRight
    },
    1: {
      0: PhCaretDoubleLeftBold,
      1: PhCaretDoubleRightBold
    }
  };

  const arrowIcon = Number(props.arrowIcon || false) as NumberBool;

  const collapsed = Number(props.collapsed || false) as NumberBool;

  return icons[arrowIcon][collapsed];
});
</script>

<template>
  <ButtonIcon
    :key="String(collapsed)"
    :tooltip-content="collapsed ? $t('icon.expand') : $t('icon.collapse')"
    tooltip-placement="bottom-start"
    :z-index="zIndex"
  >
    <!-- <SvgIcon :icon="icon" /> -->
    <component :is="icon" />
  </ButtonIcon>
</template>

<style scoped></style>
