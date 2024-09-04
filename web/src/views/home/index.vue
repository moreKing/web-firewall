<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useAppStore } from '@/store/modules/app';
import pkg from '~/package.json';
import { useEcharts } from '@/hooks/common/echarts';
// import HeaderBanner from './modules/header-banner.vue';
// import CardData from './modules/card-data.vue';
// import LineChart from './modules/line-chart.vue';
// import PieChart from './modules/pie-chart.vue';
import { getSystemHome } from '@/service/api';
import { parseData, userLineChartOpts } from './modules/utils';

const appStore = useAppStore();

const gap = computed(() => (appStore.isMobile ? 0 : 16));

const column = computed(() => (appStore.isMobile ? 1 : 2));

interface HomeData {
  assets: number;
  users: number;
  online: number;
  accounts: number;
  yesterday: number[];
  today: number[];
  xAxis: string[];
  license: number;
}

const formData = ref<HomeData>({} as HomeData);

interface PkgJson {
  name: string;
  version: string;
  dependencies: PkgVersionInfo[];
  devDependencies: PkgVersionInfo[];
}

interface PkgVersionInfo {
  name: string;
  version: string;
}

const { name, version, dependencies, devDependencies } = pkg;

function transformVersionData(tuple: [string, string]): PkgVersionInfo {
  const [$name, $version] = tuple;
  return {
    name: $name,
    version: $version
  };
}

const pkgJson: PkgJson = {
  name,
  version,
  dependencies: Object.entries(dependencies).map(item => transformVersionData(item)),
  devDependencies: Object.entries(devDependencies).map(item => transformVersionData(item))
};

const latestBuildTime = BUILD_TIME;

const { domRef, updateOptions } = useEcharts(() => {
  return userLineChartOpts();
});

const loading = ref(false);

async function initData() {
  loading.value = true;
  const { data, error } = await getSystemHome();
  loading.value = false;
  if (error) return;
  formData.value = parseData(data);

  updateOptions(opts => {
    opts.xAxis.data = formData.value.xAxis;
    opts.series[0].data = formData.value.yesterday;
    opts.series[1].data = formData.value.today;
    return opts;
  });
}

// 在线趋势
onMounted(() => {
  initData();
});
</script>

<template>
  <NSpace vertical :size="16" class="mt-10px">
    <NCard :title="$t('page.about.projectInfo.title')" :bordered="false" size="small" segmented class="card-wrapper">
      <NDescriptions label-placement="left" bordered size="small" :column="column" class="m-10px">
        <NDescriptionsItem :label="$t('page.about.projectInfo.version')">
          <NTag type="primary">{{ pkgJson.version }}</NTag>
        </NDescriptionsItem>
        <NDescriptionsItem :label="$t('page.about.projectInfo.latestBuildTime')">
          <NTag type="primary">{{ latestBuildTime }}</NTag>
        </NDescriptionsItem>
        <NDescriptionsItem :label="$t('page.home.document_project')">
          <a class="text-primary" href="https://moujun.top/" target="_blank" rel="noopener noreferrer">
            {{ $t('page.home.document_project_link') }}
          </a>
        </NDescriptionsItem>
      </NDescriptions>
    </NCard>

    <!-- <CardData /> -->
    <NGrid :x-gap="gap" :y-gap="16" responsive="screen" item-responsive>
      <NGi span="24 s:24 m:24">
        <NCard
          :title="$t('page.home.onlineTrend')"
          :bordered="false"
          size="small"
          segmented
          class="h-480px card-wrapper"
        >
          <!-- <LineChart /> -->
          <div ref="domRef" class="h-360px w-full"></div>
        </NCard>
      </NGi>
    </NGrid>
  </NSpace>
</template>

<style scoped>
:deep(.n-card-header) {
  padding: 10px var(--n-padding-left);
}
</style>
