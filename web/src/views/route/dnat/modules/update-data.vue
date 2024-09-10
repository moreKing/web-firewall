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

interface FormValue {
  protocol: null | string;
  dipAny: boolean;
  dip: string;
  tip: string;
  ports: {
    key: string;
    value: string;
  }[];
  iif: string;
  add: number;
  position: number;
  comment: string;
}

const formValue = ref<FormValue>({
  protocol: null,
  dipAny: true,
  dip: '',
  tip: '',
  ports: [
    {
      key: '',
      value: ''
    }
  ],
  iif: '',
  add: 1,
  position: 0,
  comment: ''
});

const rules = computed<any>(() => {
  // inside computed to make locale reactive, if not apply i18n, you can define it without computed
  const { defaultRequiredRule } = useFormRules();

  return {
    iif: [defaultRequiredRule],
    protocol: [defaultRequiredRule],

    ports: [
      {
        trigger: ['input', 'change'],
        type: 'array',
        required: true,
        message: $t('form.required')
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

    tip: [
      {
        required: true,
        trigger: ['input', 'change'],
        validator(_rule: any, value: string) {
          if (!value || value === '') return new Error($t('form.required'));
          if (!checkIpAddr(value)) {
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
    protocol: null,
    dipAny: true,
    dip: '',
    tip: '',
    ports: [],
    iif: '',
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
    protocol: 'iif',
    field: '',
    Value: formValue.value.iif
  });

  if (!formValue.value.dipAny) {
    expr.push({
      type: 'match',
      protocol: 'ip',
      field: 'daddr',
      Value: formValue.value.dip
    });
  }

  const tps = formValue.value.ports.map((item: any) => {
    return `${item.key} : ${formValue.value.tip} . ${item.value}`;
  });
  expr.push({
    type: 'match',
    protocol: 'dnat',
    field: `ip to ${formValue.value.protocol} dport map`,
    Value: tps.join(', ')
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

    if (item.protocol === 'iif') {
      formValue.value.iif = item.value;
      return;
    }

    if (item.protocol === 'dnat') {
      if (item.field.includes(' tcp ')) {
        formValue.value.protocol = 'tcp';
      } else if (item.field.includes(' udp ')) {
        formValue.value.protocol = 'udp';
      } else {
        formValue.value.protocol = null;
      }
    }

    item.value.split(',').forEach((item2: any) => {
      const [key, value] = item2.split(':');
      formValue.value.ports.push({
        key: key.trim(),
        value: value.split(' . ')[1].trim()
      });
      formValue.value.tip = value.split(' . ')[0].trim();
    });

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
        <NFormItem :label="$t('page.firewallPolicy.sourceEthernet')" path="iif">
          <!-- <NInput v-model:value="formValue.protocol" /> -->
          <NSelect v-model:value="formValue.iif" :options="networkOptions" />
        </NFormItem>

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
              }
            ]"
          />
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

        <NFormItem :label="$t('page.firewallPolicy.nat')" path="tip">
          <NInput v-model:value="formValue.tip" />
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.natPort')" path="ports">
          <NDynamicInput
            v-model:value="formValue.ports"
            item-style="margin-bottom: 0;"
            :on-create="() => ({ key: '', value: '' })"
          >
            <template #action="{ index, create, remove }">
              <NSpace class="ml-20px w-100px">
                <NButton strong secondary type="success" @click="() => create(index)">
                  <NIcon>
                    <icon-carbon:add-large />
                  </NIcon>
                </NButton>

                <NButton v-if="index > 0" strong secondary type="error" @click="() => remove(index)">
                  <NIcon>
                    <icon-carbon:subtract-large />
                  </NIcon>
                </NButton>
              </NSpace>
            </template>

            <template #default="{ index }">
              <div class="flex">
                <NFormItem
                  ignore-path-change
                  :show-label="false"
                  :path="`ports[${index}].key`"
                  :rule="{
                    trigger: ['input', 'change'],
                    validator(_rule: any, value: string) {
                      if (!value || value.length === 0) return new Error($t('form.required'));
                      const pattern = /^\d+$/;
                      if (!pattern.test(value)) return new Error($t('page.firewallPolicy.portValidationFailure'));

                      const intItem = Number.parseInt(value, 10);
                      if (intItem < 0 || intItem > 65535) {
                        return new Error($t('page.firewallPolicy.portValidationFailure'));
                      }
                      return true;
                    }
                  }"
                >
                  <NInput
                    v-model:value="formValue.ports[index].key"
                    :placeholder="$t('page.firewallPolicy.destPort')"
                    @keydown.enter.prevent
                  />
                  <!-- 由于在 input 元素里按回车会导致 form 里面的 button 被点击，所以阻止了默认行为 -->
                </NFormItem>
                <!-- <div class="ml-8px mr-8px h-34px lh-34px">=</div> -->
                <icon-carbon:arrow-right class="ml-8px mr-8px h-34px lh-34px" />
                <NFormItem
                  ignore-path-change
                  :show-label="false"
                  :path="`ports[${index}].value`"
                  :rule="{
                    trigger: ['input', 'change'],
                    validator(_rule: any, value: string) {
                      if (!value || value.length === 0) return new Error($t('form.required'));
                      const pattern = /^\d+$/;
                      if (!pattern.test(value)) return new Error($t('page.firewallPolicy.portValidationFailure'));

                      const intItem = Number.parseInt(value, 10);
                      if (intItem < 0 || intItem > 65535) {
                        return new Error($t('page.firewallPolicy.portValidationFailure'));
                      }
                      return true;
                    }
                  }"
                >
                  <NInput
                    v-model:value="formValue.ports[index].value"
                    :placeholder="$t('page.firewallPolicy.natPort')"
                    @keydown.enter.prevent
                  />
                </NFormItem>
              </div>
            </template>
          </NDynamicInput>
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
