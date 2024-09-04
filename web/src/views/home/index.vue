<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useThemeVars } from 'naive-ui';
import { useAppStore } from '@/store/modules/app';
import pkg from '~/package.json';
import { useEcharts } from '@/hooks/common/echarts';
// import HeaderBanner from './modules/header-banner.vue';
// import CardData from './modules/card-data.vue';
// import LineChart from './modules/line-chart.vue';
// import PieChart from './modules/pie-chart.vue';
import { getSystemHome } from '@/service/api';
import { licenseOptions, parseData, userLineChartOpts } from './modules/utils';

const themeVars = useThemeVars();

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
const licenseColor = () => {
  const v = Math.round((formData.value.assets / formData.value.license) * 100);

  if (v > 20) return themeVars.value.infoColorSuppl;
  if (v > 75 && v <= 85) return themeVars.value.warningColorSuppl;
  if (v > 85) return themeVars.value.errorColorSuppl;

  return themeVars.value.successColorSuppl;
};

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

const { domRef: domLicenseRef, updateOptions: updateLicenseOptions } = useEcharts(() => {
  return licenseOptions();
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

  updateLicenseOptions(opts => {
    opts.title.text = `{value|${Math.round((formData.value.assets / formData.value.license) * 100)}} {unit|%}`;
    opts.series[0].data[0].label.formatter = (formData.value.license - formData.value.assets).toString();
    opts.series[0].data[1].label.formatter = formData.value.assets.toString();
    opts.series[0].data[0].value = formData.value.license - formData.value.assets;
    opts.series[0].data[1].value = formData.value.assets;
    opts.series[0].data[0].itemStyle.color = `${licenseColor()}22`;
    opts.series[0].data[1].itemStyle.color = licenseColor();

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
    <NGrid cols="s:1 m:2 l:4" responsive="screen" :x-gap="16" :y-gap="16">
      <NGi>
        <NCard :title="$t('page.home.userCount')" :bordered="false" size="small" segmented class="card-wrapper">
          <NSpace align="center" size="large" class="mb-10px mt-10px">
            <SvgIcon
              local-icon="home-users"
              :style="{
                color: themeVars.successColor
              }"
              class="text-size-40px"
            />
            <NSkeleton v-if="loading" :width="146" :sharp="false" size="medium" />
            <span v-else class="text-size-14px">{{ formData.users }}</span>
          </NSpace>
        </NCard>
      </NGi>

      <NGi>
        <NCard :title="$t('page.home.assetCount')" :bordered="false" size="small" segmented class="card-wrapper">
          <NSpace align="center" size="large" class="mb-10px mt-10px">
            <SvgIcon
              local-icon="home-assets"
              :style="{
                color: themeVars.infoColor
              }"
              class="text-size-40px"
            />
            <NSkeleton v-if="loading" :width="146" :sharp="false" size="medium" />
            <span v-else class="text-size-14px">{{ formData.assets }}</span>
          </NSpace>
        </NCard>
      </NGi>

      <NGi>
        <NCard :title="$t('page.home.accountCount')" :bordered="false" size="small" segmented class="card-wrapper">
          <NSpace align="center" size="large" class="mb-10px mt-10px">
            <SvgIcon
              local-icon="home-accounts"
              :style="{
                color: themeVars.warningColor
              }"
              class="text-size-40px"
            />
            <NSkeleton v-if="loading" :width="146" :sharp="false" size="medium" />
            <span v-else class="text-size-14px">{{ formData.accounts }}</span>
          </NSpace>
        </NCard>
      </NGi>

      <NGi>
        <NCard :title="$t('page.home.online')" :bordered="false" size="small" segmented class="card-wrapper">
          <NSpace align="center" size="large" class="mb-10px mt-10px">
            <SvgIcon
              local-icon="home-online"
              :style="{
                color: themeVars.errorColor
              }"
              class="text-size-40px"
            />
            <NSkeleton v-if="loading" :width="146" :sharp="false" size="medium" />
            <span v-else class="text-size-14px">{{ formData.online }}</span>
          </NSpace>
        </NCard>
      </NGi>
    </NGrid>

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
      <NGi span="24 s:24 m:14">
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
      <NGi span="24 s:24 m:10">
        <NCard
          :title="$t('page.home.licenseUsageRate')"
          :bordered="false"
          size="small"
          segmented
          class="h-480px card-wrapper"
        >
          <template #header-extra>
            <NTag type="info">{{ $t('page.home.licenseNum') }}: {{ formData.license }}</NTag>
          </template>
          <!--
 <div class="h-full flex justify-center" style="align-items: center">
            <div class="w-50%">
              <NProgress
                type="circle"
                :color="licenseColor"
                :fill-border-radius="0"
                :stroke-width="10"
                style="width: 100%"
                :percentage="Math.round((formData.assets / formData.license) * 100)"
              />
              <NSpace justify="space-between" class="mt-10px">
                <div>{{ $t('page.home.licenseNum') }} : {{ formData.license }}</div>
                <div>{{ $t('page.home.assetCount') }} : {{ formData.assets }}</div>
              </NSpace>
            </div>
          </div>
-->

          <div ref="domLicenseRef" class="h-360px w-full"></div>
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
