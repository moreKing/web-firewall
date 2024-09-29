<script setup lang="ts">
import { computed, ref } from 'vue';
import { $t } from '@/locales';
import { changeInputPolicyPosition } from '@/service/api';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { checkIpMask } from '@/utils/ip_check';

const { formRef, validate } = useNaiveForm();

const showModal = defineModel<boolean>('show');

const emit = defineEmits<{
  (e: 'close'): void;
}>();

const props = defineProps<{
  row: any;
}>();

const loading = ref(false);

const formValue = ref<any>({
  add: 3,
  position: null,
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
    add: 3,
    position: null,
    comment: ''
  };
  emit('close');
}

async function onSubmit() {
  await validate();
  //  提交数据
  loading.value = true;

  const { error } = await changeInputPolicyPosition({
    id: props.row.id,
    add: !(formValue.value.add === 1 || formValue.value.add === 3),
    position: formValue.value.add > 2 ? formValue.value.position : 0
  });
  loading.value = false;
  if (error) return;
  window.$message?.success($t('common.updateSuccess'));
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
    :title="$t('page.firewallPolicy.position')"
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
          <NInput :value="props.row.comment" type="textarea" readonly />
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
