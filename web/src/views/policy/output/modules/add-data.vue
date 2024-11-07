<script setup lang="ts">
import { computed, ref } from 'vue';
import { $t } from '@/locales';
import { addOutputPolicy } from '@/service/api';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { checkIpMask } from '@/utils/ip_check';

const { formRef, validate } = useNaiveForm();

const showModal = defineModel<boolean>('show');

// const chain = 2;

const emit = defineEmits<{
  (e: 'close'): void;
}>();

const loading = ref(false);

const formValue = ref({
  protocol: 'tcp',
  port: '',
  pingOptions: [],
  ctStateOptions: [],
  limitIp: false,
  ip: '',
  add: 1,
  position: 0,
  policy: 'accept',
  comment: ''
});

const rules = computed(() => {
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
    policy: [defaultRequiredRule]
  };
});

function initData() {
  formValue.value = {
    protocol: 'tcp',
    port: '',
    pingOptions: [],
    ctStateOptions: [],
    limitIp: false,
    ip: '',
    policy: 'accept',
    add: 1,
    position: 0,
    comment: ''
  };
  emit('close');
}

async function onSubmit() {
  await validate();
  //  提交数据
  loading.value = true;

  const { error } = await addOutputPolicy({
    comment: formValue.value.comment,
    add: !(formValue.value.add === 1 || formValue.value.add === 3),
    position: formValue.value.add > 2 ? formValue.value.position : 0,
    ip: formValue.value.ip,
    policy: formValue.value.policy,
    port: formValue.value.port,
    protocol: formValue.value.protocol,
    ct: formValue.value.ctStateOptions.join(','),
    icmp: formValue.value.pingOptions.join(',')
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
                label: 'tcp',
                value: 'tcp'
              },
              {
                label: 'udp',
                value: 'udp'
              },
              {
                label: 'icmp',
                value: 'icmp type'
              },
              {
                label: 'ct state',
                value: 'ct state'
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

        <NFormItem
          v-if="formValue.protocol === 'ct state'"
          :label="$t('page.firewallPolicy.option')"
          path="ctStateOptions"
        >
          <NCheckboxGroup v-model:value="formValue.ctStateOptions">
            <NSpace>
              <NCheckbox value="new" :label="$t('page.firewallPolicy.newTcp')" />
              <NCheckbox value="established" :label="$t('page.firewallPolicy.establishedTcp')" />
              <NCheckbox value="related" :label="$t('page.firewallPolicy.relatedTcp')" />
              <NCheckbox value="untracked" :label="$t('page.firewallPolicy.untrackedTcp')" />
              <NCheckbox value="invalid" :label="$t('page.firewallPolicy.invalidTcp')" />
            </NSpace>
          </NCheckboxGroup>
        </NFormItem>

        <NFormItem
          v-if="formValue.protocol === 'icmp type'"
          :label="$t('page.firewallPolicy.option')"
          path="pingOptions"
        >
          <NCheckboxGroup v-model:value="formValue.pingOptions">
            <NSpace>
              <NCheckbox value="echo-reply" :label="$t('page.firewallPolicy.pingReply')" />
              <NCheckbox value="echo-request" :label="$t('page.firewallPolicy.pingRequest')" />
            </NSpace>
          </NCheckboxGroup>
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
            <NInput v-model:value="formValue.ip"    type="textarea"
              :autosize="{
                minRows: 1,
                maxRows: 5
              }"/>
            <span class="mb-30px mt-10px font-size-14px text-truegray-400">
              {{ $t('page.firewallPolicy.ipTip') }}
            </span>
          </NSpace>
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.action')" path="policy">
          <!-- <NInput v-model:value="formValue.policy" /> -->
          <NRadioGroup v-model:value="formValue.policy" name="policy">
            <NSpace>
              <NRadio value="accept">
                {{ $t('page.firewallPolicy.accept') }}
              </NRadio>
              <NRadio value="drop">
                {{ $t('page.firewallPolicy.reject') }}
              </NRadio>
            </NSpace>
          </NRadioGroup>
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
