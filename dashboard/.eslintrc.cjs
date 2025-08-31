/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution')

module.exports = {
  root: true,
  extends: [
    'plugin:vue/vue3-recommended',
    'eslint:recommended',
    '@vue/eslint-config-typescript',
    '@vue/eslint-config-prettier',
    './.eslintrc-auto-import.json',
  ],
  parser: 'vue-eslint-parser',
  parserOptions: {
    ecmaVersion: 'latest',
  },
  plugins: ['simple-import-sort'],
  rules: {
    // 允许被禁止的类型 https://typescript-eslint.io/rules/ban-types/
    '@typescript-eslint/ban-types': 'off',
    // 允许any类型 https://typescript-eslint.io/rules/no-explicit-any/
    '@typescript-eslint/no-explicit-any': 'off',
    // 允许任何TS指令注释 https://typescript-eslint.io/rules/ban-ts-comment
    '@typescript-eslint/ban-ts-comment': 'off',
    // 允许空函数 https://typescript-eslint.io/rules/no-empty-function/
    '@typescript-eslint/no-empty-function': 'off',
    // 禁止函数无返回值类型 https://typescript-eslint.io/rules/explicit-function-return-type/
    '@typescript-eslint/explicit-function-return-type': 'error',
    // 禁止未使用变量，除了下划线开头 https://typescript-eslint.io/rules/no-unused-vars/
    '@typescript-eslint/no-unused-vars': [
      'error',
      {
        argsIgnorePattern: '^_',
        varsIgnorePattern: '^_',
      },
    ],
    // 关闭强制多词组件名 https://eslint.vuejs.org/rules/multi-word-component-names.html
    'vue/multi-word-component-names': 'off',
    // 强制执行自我关闭风格 https://eslint.vuejs.org/rules/html-self-closing.html
    'vue/html-self-closing': [
      'error',
      {
        html: {
          void: 'always',
          normal: 'never',
          component: 'always',
        },
        svg: 'always',
        math: 'always',
      },
    ],
    // 禁止未使用变量，除了下划线开头 https://eslint.org/docs/latest/rules/no-unused-vars
    'no-unused-vars': [
      'error',
      {
        argsIgnorePattern: '^_',
        varsIgnorePattern: '^_',
      },
    ],
    'simple-import-sort/imports': 'error',
    'simple-import-sort/exports': 'error',
  },
}
