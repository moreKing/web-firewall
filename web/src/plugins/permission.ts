import type { ObjectDirective } from 'vue';
import { useAuth } from '@/hooks/business/auth';

export const permission: ObjectDirective = {
  mounted(el: HTMLButtonElement, binding) {
    if (!binding.value) return;
    // 获取传入参数
    const { hasAuth } = useAuth();
    if (!hasAuth(binding.value)) {
      el.remove();
    }
  }
};
