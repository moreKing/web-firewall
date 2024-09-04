<script setup lang="ts">
import * as monaco from 'monaco-editor';
// import { language } from 'monaco-editor/esm/vs/basic-languages/sql/sql.js';
import { language } from 'monaco-editor/esm/vs/basic-languages/python/python';
import { computed, onBeforeUnmount, onMounted, ref } from 'vue';
import EditorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';
import { $t } from '@/locales';
import { useThemeStore } from '@/store/modules/theme';

const { darkMode } = useThemeStore();

window.MonacoEnvironment = {
  getWorker() {
    return new EditorWorker();
  }
};

// const code = ref<string>('select * from table'); // 代码

const languageCode = ref('python');

// 获取 SQL 的关键字
const { keywords } = language;

let editor: any;

// 初始化 SQL 代码和表格数据
const tables: any = {};

// 编辑器的主题设置
const theme = computed(() => {
  return darkMode ? 'vs-dark' : 'vs-light';
});
const monacoEditor = ref<any | null>(null);

// 组件挂载后创建编辑器实例
onMounted(() => {
  initAutoCompletion();
  editor = monaco.editor.create(monacoEditor.value, {
    value: '',
    language: languageCode.value,
    readOnly: false,
    automaticLayout: true,
    colorDecorators: true, // 颜色装饰器
    theme: theme.value,
    minimap: {
      enabled: false
    },
    tabSize: 2,
    fontSize: 16
  });
});
// 组件卸载前销毁编辑器实例
onBeforeUnmount(() => {
  if (editor) {
    editor.dispose();
  }
});
// /** @description: 获取编辑器中填写的值 */
// function getValue() {
//   return editor.getValue();
// }

/** @description: 初始化自动补全 */
function initAutoCompletion() {
  monaco.languages.registerCompletionItemProvider(languageCode.value, {
    triggerCharacters: ['.', ' ', ...keywords],
    provideCompletionItems: (model, position) => {
      let suggestions: any[] = [];
      const { lineNumber, column } = position;
      const textBeforePointer = model.getValueInRange({
        startLineNumber: lineNumber,
        startColumn: 0,
        endLineNumber: lineNumber,
        endColumn: column
      });
      const words = textBeforePointer.trim().split(/\s+/);
      const lastWord = words[words.length - 1];

      if (lastWord.endsWith('.')) {
        const tableName = lastWord.slice(0, lastWord.length - 1);
        if (Object.keys(tables).includes(tableName)) {
          suggestions = [...getFieldsSuggest(tableName)];
        }
      } else if (lastWord === '.') {
        suggestions = [];
      } else {
        suggestions = [...getTableSuggest(), ...getKeywordsSuggest()];
      }

      return {
        suggestions
      };
    }
  });
}

/** @description: 获取关键字的补全列表 */
function getKeywordsSuggest() {
  return keywords.map((key: any) => ({
    label: key,
    kind: monaco.languages.CompletionItemKind.Keyword,
    insertText: key
  }));
}

/** @description: 获取表名的补全列表 */
function getTableSuggest() {
  return Object.keys(tables).map(key => ({
    label: key,
    kind: monaco.languages.CompletionItemKind.Variable,
    insertText: key
  }));
}

/**
 * @param {any} tableName
 * @description: 根据表名获取字段补全列表
 */
function getFieldsSuggest(tableName: string) {
  const fields = tables[tableName];
  if (!fields) {
    return [];
  }
  return fields.map((name: string) => ({
    label: name,
    kind: monaco.languages.CompletionItemKind.Field,
    insertText: name
  }));
}
</script>

<template>
  <div class="codemirror h-full">
    <div id="monacoEditor" ref="monacoEditor" class="container-code" />
    <NSpace justify="end">
      <NSpace class="mt-10px">
        <NButton>{{ $t('common.cancel') }}</NButton>
        <NButton>
          {{ $t('page.userSetting.submit') }}
        </NButton>
      </NSpace>
    </NSpace>
  </div>
</template>

<style scoped lang="scss">
:deep(.container-code) {
  height: calc(100% - 80px);
  border: 1px solid rgba(239, 239, 245, 1);
}
</style>
