<script setup lang="ts">
import { computed, defineModel, ref } from 'vue';
import type { FormInst } from 'naive-ui';
import { $t } from '@/locales';

const showModal = defineModel<boolean>('show');
const props = defineProps<{
  issuer: string;
  secret: string;
  username: string;
}>();

const emit = defineEmits<{
  (e: 'finish'): void;
}>();
const model = defineModel<any>('model');

const text = computed(() => {
  // `otpauth://totp/${data.issuer}:${username}@${document.location.hostname}?secret=${data.token}&issuer=${data.issuer}`
  return `otpauth://totp/${props.issuer}:${props.username}@${document.location.hostname}?secret=${props.secret}&issuer=${props.issuer}`;
});

const rules = computed<any>(() => {
  return {
    code: [
      {
        type: 'string',
        required: true,
        message: $t('form.required'),
        trigger: ['input', 'blur']
      }
    ]
  };
});

const formRef = ref<FormInst | null>(null);
async function handleSubmit() {
  await formRef.value?.validate();
  // model.value.code = formValue.value.code;
  emit('finish');
}
</script>

<template>
  <NModal
    v-model:show="showModal"
    transform-origin="center"
    :close-on-esc="false"
    :mask-closable="false"
    preset="card"
    class="w-420px"
    :title="$t('datatable.code')"
  >
    <NTabs v-if="props.secret" type="segment" animated>
      <NTabPane name="scan" :tab="$t('page.login.bindCode.scanningBinding')">
        <div class="mt-10px h-310px text-center">
          <!-- <QrcodeVue :value="otpauth" :size="300" level="H" style="text-align: center; margin: auto" class="" /> -->
          <NQrCode :value="text" :size="280" :padding="0" />
        </div>
      </NTabPane>
      <NTabPane name="manua" :tab="$t('page.login.bindCode.manualInput')">
        <div>
          <div class="mt-20px h-300px flex text-center">
            <div class="w-full flex-self-center text-center">
              <div class="font-size-5 font-900">
                {{ props.secret }}
              </div>
              <div class="mt-10px">{{ $t('page.login.bindCode.tip') }}</div>
            </div>
          </div>
        </div>
      </NTabPane>
    </NTabs>

    <NForm ref="formRef" :label-width="80" label-placement="left" label-align="left" :rules="rules" :model="model">
      <NFormItem path="code">
        <NInput v-model:value="model.code" type="password" show-password-on="mousedown" @keyup.enter="handleSubmit" />
      </NFormItem>
    </NForm>

    <template #footer>
      <NSpace justify="end">
        <NButton type="info" @click="emit('finish')">
          {{ $t('page.userSetting.submit') }}
        </NButton>
      </NSpace>
    </template>
  </NModal>
</template>
