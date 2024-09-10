<script setup lang="ts">
import { computed, ref } from 'vue';
import { $t } from '@/locales';
import { updateFirewallPolicy } from '@/service/api';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { checkIpMask } from '@/utils/ip_check';

const { formRef, validate } = useNaiveForm();

const showModal = defineModel<boolean>('show');

interface Rule {
  id: number;
  comment: string;
  expr: any[];
}

const props = defineProps<{
  row: Rule;
}>();

const emit = defineEmits<{
  (e: 'close'): void;
}>();

const loading = ref(false);

interface FormValue {
  protocol: string;
  dipAny: boolean;
  dip: string;
  portType: boolean;
  port: string;
  sipAny: boolean;
  sip: string;
  limit: number | null;
  limitSpeed: string;
  add: number;
  position: number;
  comment: string;
}

const formValue = ref<FormValue>({
  protocol: 'any',

  dipAny: true,

  dip: '',
  portType: false,
  port: '',

  sipAny: true,
  sip: '',

  limit: null,
  limitSpeed: 'mbytes/second',

  add: 1,
  position: 0,
  comment: ''
});

const rules = computed<any>(() => {
  // inside computed to make locale reactive, if not apply i18n, you can define it without computed
  const { defaultRequiredRule } = useFormRules();

  return {
    protocol: [defaultRequiredRule],

    sip: [
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

    dip: [
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
    port: [
      defaultRequiredRule,
      {
        trigger: ['input', 'change'],
        pattern: /^\d[\d,-]*\d$/,
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

    policy: [defaultRequiredRule]
  };
});

function initData() {
  formValue.value = {
    protocol: 'any',

    dipAny: true,

    dip: '',
    portType: false,
    port: '',

    sipAny: true,
    sip: '',

    limit: null,
    limitSpeed: 'mbytes/second',

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
  const expr: any = [];

  if (!formValue.value.sipAny) {
    expr.push({
      type: 'match',
      protocol: 'ip',
      field: 'saddr',
      Value: formValue.value.sip
    });
  }

  if (!formValue.value.dipAny) {
    expr.push({
      type: 'match',
      protocol: 'ip',
      field: 'daddr',
      Value: formValue.value.dip
    });
  }

  if (formValue.value.protocol !== 'any') {
    expr.push({
      type: 'match',
      protocol: formValue.value.protocol,
      field: formValue.value.portType ? 'sport' : 'dport',
      Value: formValue.value.port
    });
  }

  expr.push({
    type: 'match',
    protocol: 'limit',
    field: 'rate over',
    Value: `${formValue.value.limit} ${formValue.value.limitSpeed}`
  });

  expr.push({
    type: 'policy',
    policy: 'drop'
  });

  const { error } = await updateFirewallPolicy({
    id: props.row.id,
    comment: formValue.value.comment,
    expr
  });
  loading.value = false;
  if (error) return;
  window.$message?.success($t('common.modifySuccess'));
  showModal.value = false;
}

async function enterModal() {
  props.row.expr.forEach((item: any) => {
    if (item.protocol === 'ip' && item.field === 'daddr') {
      formValue.value.dipAny = false;
      formValue.value.dip = item.value;
      return;
    }

    if (item.protocol === 'ip' && item.field === 'saddr') {
      formValue.value.sipAny = false;
      formValue.value.sip = item.value;
    }

    if (item.protocol === 'tcp' || item.protocol === 'udp') {
      formValue.value.portType = item.field === 'sport';
      formValue.value.protocol = item.protocol;
      formValue.value.port = item.value;
    }

    if (item.protocol === 'limit') {
      const tmpLimit = item.value.trim().split(/\s+/);
      formValue.value.limit = Number.parseInt(tmpLimit[0].trim(), 10);
      formValue.value.limitSpeed = tmpLimit[1].trim();
    }

    // if (item.protocol === 'snat')
  });

  loading.value = false;
}
</script>

<template>
  <NModal
    v-model:show="showModal"
    :mask-closable="false"
    preset="card"
    class="w-700px"
    :title="$t('common.edit')"
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
                value: 'any'
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

        <NFormItem :label="$t('page.firewallPolicy.sourceIp')" path="sipAny">
          <NRadioGroup v-model:value="formValue.sipAny" name="radiogroup">
            <NSpace>
              <NRadio :value="true">
                {{ $t('page.firewallPolicy.allIp') }}
              </NRadio>
              <NRadio :value="false">
                {{ $t('page.firewallPolicy.partialIp') }}
              </NRadio>
            </NSpace>
          </NRadioGroup>
        </NFormItem>

        <NFormItem v-if="!formValue.sipAny" label=" " path="sip">
          <NSpace vertical :size="14" class="w-full">
            <NInput v-model:value="formValue.sip" />
            <span class="mb-30px mt-10px font-size-14px text-truegray-400">
              {{ $t('page.firewallPolicy.ipTip') }}
            </span>
          </NSpace>
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.destIp')" path="dipAny">
          <NRadioGroup v-model:value="formValue.dipAny" name="radiogroup">
            <NSpace>
              <NRadio :value="true">
                {{ $t('page.firewallPolicy.allIp') }}
              </NRadio>
              <NRadio :value="false">
                {{ $t('page.firewallPolicy.partialIp') }}
              </NRadio>
            </NSpace>
          </NRadioGroup>
        </NFormItem>

        <NFormItem v-if="!formValue.dipAny" label=" " path="dip">
          <NSpace vertical :size="14" class="w-full">
            <NInput v-model:value="formValue.dip" />
            <span class="mb-30px mt-10px font-size-14px text-truegray-400">
              {{ $t('page.firewallPolicy.ipTip') }}
            </span>
          </NSpace>
        </NFormItem>

        <div v-if="formValue.protocol !== 'any'">
          <NFormItem :label="$t('page.firewallPolicy.port')" path="portType">
            <NRadioGroup v-model:value="formValue.portType" name="radiogroup">
              <NSpace>
                <NRadio :value="true">
                  {{ $t('page.firewallPolicy.sourcePort') }}
                </NRadio>
                <NRadio :value="false">
                  {{ $t('page.firewallPolicy.destPort') }}
                </NRadio>
              </NSpace>
            </NRadioGroup>
          </NFormItem>

          <NFormItem label=" " path="port">
            <NSpace vertical :size="14" class="w-full">
              <NInput v-model:value="formValue.port" />
              <span class="mb-30px mt-10px font-size-14px text-truegray-400">
                {{ $t('page.firewallPolicy.portTip') }}
              </span>
            </NSpace>
          </NFormItem>
        </div>

        <NFormItem :label="$t('page.firewallPolicy.speed')" path="limit">
          <!-- <NInput v-model:value="formValue.policy" /> -->
          <NSpace>
            <NInputNumber v-model:value="formValue.limit" />

            <NSelect
              v-model:value="formValue.limitSpeed"
              class="w-215px"
              :options="[
                {
                  label: 'kb/s',
                  value: 'kbytes/second'
                },
                {
                  label: 'mb/s',
                  value: 'mbytes/second'
                },
                {
                  label: 'kb/m',
                  value: 'kbytes/minute'
                },
                {
                  label: 'mb/m',
                  value: 'mbytes/minute'
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
