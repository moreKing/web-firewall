<script setup lang="ts">
import { computed, ref } from 'vue';
import { $t } from '@/locales';
import { updateFirewallPolicy } from '@/service/api';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { checkIpAddr, checkIpMask } from '@/utils/ip_check';

const { formRef, validate } = useNaiveForm();

const showModal = defineModel<boolean>('show');

interface Rule {
  id: number;
  comment: string;
  expr: any[];
}

const props = defineProps<{
  row: Rule;
  network: any;
}>();

const emit = defineEmits<{
  (e: 'close'): void;
}>();

const networkOptions = computed(() => {
  return props.network.map((item: any) => {
    return {
      label: item.name,
      value: item.name
    };
  });
});

const loading = ref(false);

const formValue = ref({
  sipAny: true,
  sip: '',
  dipAny: true,
  dip: '',
  oif: '',
  masquerade: true,
  snat: '',
  add: 1,
  position: 0,
  comment: ''
});

const rules = computed(() => {
  // inside computed to make locale reactive, if not apply i18n, you can define it without computed
  const { defaultRequiredRule } = useFormRules();

  return {
    oif: [defaultRequiredRule],
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
    snat: [
      defaultRequiredRule,
      {
        trigger: ['input', 'change'],
        validator(_rule: any, value: string) {
          if (!checkIpAddr(value)) {
            return new Error($t('page.firewallPolicy.ipValidationFailure'));
          }
          return true;
        }
      }
    ]
  };
});
function initData() {
  formValue.value = {
    sipAny: true,
    sip: '',
    dipAny: true,
    dip: '',
    oif: '',
    masquerade: true,
    snat: '',
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

  expr.push({
    type: 'match',
    protocol: 'oif',
    field: '',
    Value: formValue.value.oif
  });

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

  if (!formValue.value.masquerade) {
    expr.push({
      type: 'match',
      protocol: 'snat',
      field: 'ip to',
      Value: formValue.value.snat
    });
  } else {
    expr.push({
      type: 'match',
      protocol: 'masquerade',
      field: '',
      Value: ''
    });
  }

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
    if (item.protocol === 'ip' && item.field === 'saddr') {
      formValue.value.sipAny = false;
      formValue.value.sip = item.value;
      return;
    }

    if (item.protocol === 'ip' && item.field === 'daddr') {
      formValue.value.dipAny = false;
      formValue.value.dip = item.value;
      return;
    }

    if (item.protocol === 'oif') {
      formValue.value.oif = item.value;
      return;
    }

    if (item.protocol === 'masquerade') {
      formValue.value.masquerade = true;
    }

    if (item.protocol === 'snat') {
      formValue.value.snat = item.value;
      formValue.value.masquerade = false;
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
        <NFormItem :label="$t('page.firewallPolicy.destinationEthernet')" path="oif">
          <!-- <NInput v-model:value="formValue.protocol" /> -->
          <NSelect v-model:value="formValue.oif" :options="networkOptions" />
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

        <NFormItem :label="$t('page.firewallPolicy.nat')" path="masquerade">
          <NRadioGroup v-model:value="formValue.masquerade" name="radiogroup">
            <NSpace>
              <NRadio :value="true">
                {{ $t('page.firewallPolicy.dynamicIp') }}
              </NRadio>
              <NRadio :value="false">
                {{ $t('page.firewallPolicy.partialIp') }}
              </NRadio>
            </NSpace>
          </NRadioGroup>
        </NFormItem>

        <NFormItem v-if="!formValue.masquerade" label=" " path="snat">
          <NInput v-model:value="formValue.snat" />
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
