<script setup lang="ts">
import { computed, ref } from 'vue';
import { $t } from '@/locales';
import { addOutputLimitPolicy } from '@/service/api';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { checkIpMask } from '@/utils/ip_check';

const { formRef, validate } = useNaiveForm();

const showModal = defineModel<boolean>('show');

const emit = defineEmits<{
  (e: 'close'): void;
}>();

const loading = ref(false);

const formValue = ref({
  protocol: 'all',
  port: '',
  limitIp: false,
  ip: '',
  add: 1,
  position: 0,
  limit: null,
  speed: 'mb/s',
  comment: ''
});

const rules = computed<any>(() => {
  // inside computed to make locale reactive, if not apply i18n, you can define it without computed
  const { defaultRequiredRule } = useFormRules();

  return {
    protocol: [defaultRequiredRule],
    port: [
      defaultRequiredRule,
      {
        trigger: ['input', 'change'],
        pattern: /^\d[\d,-]*$/,
        message: $t('page.firewallPolicy.portValidationFailure')
      },
      {
        trigger: ['input', 'change'],
        validator(_rule: any, value: string) {
          const state = value.split(',').every((item: string) => {
            const intItem = Number.parseInt(item, 10);
            if (intItem < 0 || intItem > 65535) {
              return false;
            }
            return true;
          });
          if (!state) {
            return new Error($t('page.firewallPolicy.portValidationFailure'));
          }
          return true;
        }
      }
    ],
    ip: [
      defaultRequiredRule,
      {
        trigger: ['input', 'change'],
        validator(_rule: any, value: string) {
          const state = value.split(',').every((item: string) => {
            return checkIpMask(item);
          });
          if (!state) {
            return new Error($t('page.firewallPolicy.ipValidationFailure'));
          }
          return true;
        }
      }
    ],
    policy: [defaultRequiredRule],
    limit: {
      required: true,
      type: 'number',
      min: 1,
      message: $t('form.required'),
      trigger: ['change', 'input']
    }
  };
});

function initData() {
  formValue.value = {
    protocol: 'all',
    port: '',
    limitIp: false,
    ip: '',
    add: 1,
    position: 0,
    limit: null,
    speed: 'mb/s',
    comment: ''
  };
  emit('close');
}

async function onSubmit() {
  await validate();
  //  提交数据
  loading.value = true;

  const { error } = await addOutputLimitPolicy({
    ...formValue.value,
    add: !(formValue.value.add === 1 || formValue.value.add === 3),
    position: formValue.value.add > 2 ? formValue.value.position : 0
  });
  loading.value = false;
  if (error) return;
  window.$message?.success($t('common.addSuccess'));
  showModal.value = false;
}

async function enterModal() {
  loading.value = false;
}
</script>

<template>
  <NModal
    v-model:show="showModal"
    :mask-closable="false"
    preset="card"
    class="w-700px"
    :title="$t('common.add')"
    :bordered="false"
    :segmented="{
      content: true
    }"
    @after-leave="initData"
    @after-enter="enterModal"
  >
    <NSpin :show="loading">
      <NForm
        ref="formRef"
        :model="formValue"
        label-width="100px"
        label-placement="left"
        label-align="left"
        :rules="rules"
        class="ml-20px mr-30px"
      >
        <NFormItem :label="$t('page.firewallPolicy.protocol')" path="protocol">
          <!-- <NInput v-model:value="formValue.protocol" /> -->
          <NSelect
            v-model:value="formValue.protocol"
            :options="[
              {
                label: $t('page.firewallPolicy.all'),
                value: 'all'
              },
              {
                label: 'tcp',
                value: 'tcp'
              },
              {
                label: 'udp',
                value: 'udp'
              }
            ]"
          />
        </NFormItem>

        <NFormItem
          v-if="formValue.protocol === 'tcp' || formValue.protocol === 'udp'"
          :label="$t('page.firewallPolicy.port')"
          path="port"
        >
          <NSpace vertical :size="14" class="w-full">
            <NInput v-model:value="formValue.port" />
            <span class="mb-30px mt-10px font-size-14px text-truegray-400">
              {{ $t('page.firewallPolicy.portTip') }}
            </span>
          </NSpace>
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.position')" path="position">
          <NSpace :size="14" class="w-full">
            <NInputNumber v-if="formValue.add > 2" v-model:value="formValue.position" />
            <NSelect
              v-model:value="formValue.add"
              class="w-215px"
              :options="[
                {
                  label: $t('page.firewallPolicy.start'),
                  value: 1
                },
                {
                  label: $t('page.firewallPolicy.end'),
                  value: 2
                },
                {
                  label: $t('page.firewallPolicy.beforePosition'),
                  value: 3
                },
                {
                  label: $t('page.firewallPolicy.afterPosition'),
                  value: 4
                }
              ]"
            />
          </NSpace>
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.destIp')" path="limitIp">
          <NRadioGroup v-model:value="formValue.limitIp" name="radiogroup">
            <NSpace>
              <NRadio :value="false">
                {{ $t('page.firewallPolicy.allIp') }}
              </NRadio>
              <NRadio :value="true">
                {{ $t('page.firewallPolicy.partialIp') }}
              </NRadio>
            </NSpace>
          </NRadioGroup>
        </NFormItem>

        <NFormItem v-if="formValue.limitIp" :label="$t('page.firewallPolicy.partialIp')" path="ip">
          <NSpace vertical :size="14" class="w-full">
            <NInput v-model:value="formValue.ip" />
            <span class="mb-30px mt-10px font-size-14px text-truegray-400">
              {{ $t('page.firewallPolicy.ipTip') }}
            </span>
          </NSpace>
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.speed')" path="limit">
          <!-- <NInput v-model:value="formValue.policy" /> -->
          <NSpace>
            <NInputNumber v-model:value="formValue.limit" />

            <NSelect
              v-model:value="formValue.speed"
              class="w-215px"
              :options="[
                {
                  label: 'kb/s',
                  value: 'kb/s'
                },
                {
                  label: 'mb/s',
                  value: 'mb/s'
                },
                {
                  label: 'kb/m',
                  value: 'kb/m'
                },
                {
                  label: 'mb/m',
                  value: 'mb/m'
                }
              ]"
            />
          </NSpace>
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.comment')" path="comment">
          <NInput v-model:value="formValue.comment" type="textarea" />
        </NFormItem>
      </NForm>
    </NSpin>
    <template #footer>
      <NSpace justify="end">
        <NButton @click="showModal = false">{{ $t('common.cancel') }}</NButton>
        <NButton v-throttle="onSubmit" type="primary">{{ $t('common.confirm') }}</NButton>
      </NSpace>
    </template>
  </NModal>
</template>
