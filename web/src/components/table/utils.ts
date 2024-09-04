import { computed } from 'vue';

export function showColumns(columns: any[]) {
  return computed(() => {
    // 过滤掉不需要显示的列
    return columns.filter(column => !column.hide);
  });
}
