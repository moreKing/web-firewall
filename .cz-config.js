'use strict';
module.exports = {
    "types": [
        {
            "name": "✨ 新特性",
            "value": "feat"
        },
        {
            "name": "🐞 bug修复",
            "value": "fix"
        },
        {
            "name": "⚡︎ 性能优化",
            "value": "perf"
        },
        {
            "name": "♻ 重构代码",
            "value": "refactor"
        },
        {
            "name": "🎉 样式相关",
            "value": "style"
        },
        {
            "name": "✅ 测试用例",
            "value": "test"
        },
        {
            "name": "🛠 构建工具",
            "value": "chore"
        },
        {
            "name": "📝 文档变更",
            "value": "docs"
        },
        {
            "name": "⏪ 回滚",
            "value": "revert"
        }
    ],
    scopes: [
        { name: 'worker' },
        { name: 'web' },
        { name: 'server' },
        { name: 'all' },
    ],
    // override the messages, defaults are as follows
    messages: {
        type: '选择一种你的提交类型:',
        scope: '选择影响模块:',
        // used if allowCustomScopes is true
        // customScope: 'Denote the SCOPE of this change:',
        subject: '简要说明:',
        body: '提交内容详细说明，使用"|"换行(可选): ',
        breaking: '非兼容性说明 (可选):',
        footer: '关联关闭的issue,例如: #31, #34(可选):',
        confirmCommit: '确定提交说明?'
    },

    allowCustomScopes: true,
    allowBreakingChanges: ['特性', '修复'],

    // limit subject length
    subjectLimit: 100

}