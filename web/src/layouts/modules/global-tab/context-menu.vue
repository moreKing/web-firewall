<script setup lang="ts">
import { computed, h } from 'vue';
import type { VNode } from 'vue';
import { $t } from '@/locales';
import { useTabStore } from '@/store/modules/tab';
import AntDesignCloseOutlined from '~icons/ant-design/close-outlined';
import AntDesignColumnWidthOutlined from '~icons/ant-design/column-width-outlined';
import MdiFormatHorizontalAlignLeft from '~icons/mdi/format-horizontal-align-left';
import MdiFormatHorizontalAlignRight from '~icons/mdi/format-horizontal-align-right';
import AntDesignLineOutlined from '~icons/ant-design/line-outlined';

defineOptions({
  name: 'ContextMenu'
});

interface Props {
  /** ClientX */
  x: number;
  /** ClientY */
  y: number;
  tabId: string;
  excludeKeys?: App.Global.DropdownKey[];
  disabledKeys?: App.Global.DropdownKey[];
}

const props = withDefaults(defineProps<Props>(), {
  excludeKeys: () => [],
  disabledKeys: () => []
});

const visible = defineModel<boolean>('visible');

const { removeTab, clearTabs, clearLeftTabs, clearRightTabs } = useTabStore();

type DropdownOption = {
  key: App.Global.DropdownKey;
  label: string;
  icon?: () => VNode;
  disabled?: boolean;
};

const options = computed(() => {
  const opts: DropdownOption[] = [
    {
      key: 'closeCurrent',
      label: $t('dropdown.closeCurrent'),
      icon: () => h(AntDesignCloseOutlined, {})
    },
    {
      key: 'closeOther',
      label: $t('dropdown.closeOther'),
      icon: () => h(AntDesignColumnWidthOutlined, {})
    },
    {
      key: 'closeLeft',
      label: $t('dropdown.closeLeft'),
      icon: () => h(MdiFormatHorizontalAlignLeft, {})
    },
    {
      key: 'closeRight',
      label: $t('dropdown.closeRight'),
      icon: () => h(MdiFormatHorizontalAlignRight, {})
    },
    {
      key: 'closeAll',
      label: $t('dropdown.closeAll'),
      icon: () => h(AntDesignLineOutlined, {})
    }
  ];
  const { excludeKeys, disabledKeys } = props;

  const result = opts.filter(opt => !excludeKeys.includes(opt.key));

  disabledKeys.forEach(key => {
    const opt = result.find(item => item.key === key);

    if (opt) {
      opt.disabled = true;
    }
  });

  return result;
});

function hideDropdown() {
  visible.value = false;
}

const dropdownAction: Record<App.Global.DropdownKey, () => void> = {
  closeCurrent() {
    removeTab(props.tabId);
  },
  closeOther() {
    clearTabs([props.tabId]);
  },
  closeLeft() {
    clearLeftTabs(props.tabId);
  },
  closeRight() {
    clearRightTabs(props.tabId);
  },
  closeAll() {
    clearTabs();
  }
};

function handleDropdown(optionKey: App.Global.DropdownKey) {
  dropdownAction[optionKey]?.();
  hideDropdown();
}
</script>

<template>
  <NDropdown
    :show="visible"
    placement="bottom-start"
    trigger="manual"
    :x="x"
    :y="y"
    :options="options"
    @clickoutside="hideDropdown"
    @select="handleDropdown"
  />
</template>

<style scoped></style>
