<script lang="ts" setup>
import { computed, ref } from 'vue';
import { $t } from '@/locales';
import BasicSetting from './modules/BasicSetting.vue';
import SafetySetting from './modules/SafetySetting.vue';

const typeTabList = computed(() => [
  {
    name: $t('page.userSetting.basicSetting'),
    desc: $t('page.userSetting.userInfoSetting'),
    key: 1
  },
  {
    name: $t('page.userSetting.securitySetting'),
    desc: $t('page.userSetting.passwordSetting'),
    key: 2
  }
]);

const type = ref(1);
const typeTitle = ref(typeTabList.value[0].name);

function switchType(e: any) {
  type.value = e.key;
  typeTitle.value = e.name;
}
</script>

<template>
  <div>
    <NGrid :x-gap="24">
      <NGridItem span="6">
        <NCard :bordered="false" size="small" class="proCard">
          <NThing
            v-for="item in typeTabList"
            :key="item.key"
            class="thing-cell"
            :class="{ 'thing-cell-on': type === item.key }"
            @click="switchType(item)"
          >
            <template #header>{{ item.name }}</template>
            <template #description>{{ item.desc }}</template>
          </NThing>
        </NCard>
      </NGridItem>
      <NGridItem span="18">
        <BasicSetting v-if="type === 1" />
        <SafetySetting v-if="type === 2" />
      </NGridItem>
    </NGrid>
  </div>
</template>

<style lang="scss" scoped>
.thing-cell {
  margin: 0 -16px 10px;
  padding: 5px 16px;

  &:hover {
    background: rgb(var(--primary-50-color));
    cursor: pointer;
  }
}

.thing-cell-on {
  background: rgb(var(--primary-50-color));
  color: var(--n-color-target);

  ::v-deep(.n-thing-main .n-thing-header .n-thing-header__title) {
    color: var(--n-color-target);
  }

  &:hover {
    background: rgb(var(--primary-50-color));
  }
}
</style>
