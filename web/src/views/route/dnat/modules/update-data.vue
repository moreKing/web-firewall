<script setup lang="ts">
import { computed, ref } from 'vue';
import { $t } from '@/locales';
import { updateDnatPolicy } from '@/service/api';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { checkIpAddr, checkIpMask, checkPortString } from '@/utils/ip_check';

const { formRef, validate } = useNaiveForm();

const showModal = defineModel<boolean>('show');

const props = defineProps<{
  row: any;
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
  protocol: string;
  dipAny: boolean;
  dip: string;
  dnat: string;
  port: {
    pair: [string, string];
    protocol: string;
    key: string;
    value: string;
  }[];
  iif: string;
  add: number;
  position: number;
  comment: string;
}

const formValue = ref<FormValue>({
  protocol: '',
  dipAny: true,
  dip: '',
  dnat: '',
  port: [],
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

    port: [
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

    dnat: [
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
    protocol: '',
    dipAny: true,
    dip: '',
    dnat: '',
    port: [],
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

  const { error } = await updateDnatPolicy({
    id: props.row.id,
    ...formValue.value
  });
  loading.value = false;
  if (error) return;
  window.$message?.success($t('common.modifySuccess'));
  showModal.value = false;
}

async function enterModal() {
  formValue.value = props.row;
  formValue.value.dipAny = !formValue.value.dip;

  for (let index = 0; index < formValue.value.port.length; index += 1) {
    if (!formValue.value.port[index].protocol || formValue.value.port[index].protocol === '') {
      formValue.value.port[index].protocol = formValue.value.protocol;
    }
    formValue.value.port[index].pair = [formValue.value.port[index].key, formValue.value.port[index].value];
  }

  loading.value = false;
}

function portHandle(index: number, v: string[]) {
  formValue.value.port[index].key = v[0];
  formValue.value.port[index].value = v[1];
}
</script>

<template>
  <NModal
    v-model:show="showModal"
    :mask-closable="false"
    preset="card"
    class="w-800px"
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

        <NFormItem :label="$t('page.firewallPolicy.intranetIp')" path="dnat">
          <NInput v-model:value="formValue.dnat" />
        </NFormItem>

        <NFormItem :label="$t('page.firewallPolicy.natPort')" path="port">
          <NDynamicInput
            v-model:value="formValue.port"
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
              <NSpace :wrap="false">
                <NFormItem
                  :path="`port[${index}].protocol`"
                  ignore-path-change
                  :rule="{ required: true, trigger: ['input', 'change'], message: $t('form.required') }"
                >
                  <NSelect
                    v-model:value="formValue.port[index].protocol"
                    class="w-120px"
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
                        label: 'tcp+udp',
                        value: 'tcp+udp'
                      }
                    ]"
                  />
                </NFormItem>

                <NFormItem
                  ignore-path-change
                  :show-label="false"
                  :path="`port[${index}].key`"
                  :rule="{
                    type: 'array',
                    trigger: ['blur'],
                    validator() {
                      //  判断转换前的端口
                      if (!checkPortString(formValue.port[index].key)) {
                        return new Error($t('page.firewallPolicy.portValidationFailure'));
                      }
                      // 判断转换后的端口
                      if (!checkPortString(formValue.port[index].value)) {
                        return new Error($t('page.firewallPolicy.portValidationFailure'));
                      }
                      return true;
                    }
                  }"
                >
                  <NInput
                    v-model:value="formValue.port[index].pair"
                    pair
                    :placeholder="[$t('page.firewallPolicy.destPort'), $t('page.firewallPolicy.natPort')]"
                    separator="→"
                    @keydown.enter.prevent
                    @update:value="v => portHandle(index, v)"
                  />
                </NFormItem>
              </NSpace>
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
