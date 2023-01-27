module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: ['plugin:react/recommended', 'google', 'standard-with-typescript', 'prettier'],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    project: ['tsconfig.json'],
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 'latest',
    sourceType: 'module',
  },
  plugins: ['react', '@typescript-eslint', 'import', 'prettier'],
  settings: {
    react: {
      version: 'detect',
    },
  },
  ignorePatterns: ['next-env.d.ts', 'next.config.js', '.prettierrc.js', '.eslintrc.js'],
  rules: {
    'react/react-in-jsx-scope': 'off',
    'require-jsdoc': 'off',
    'react/no-unescaped-entities': 0,
    'import/order': [
      'error',
      {
        groups: ['index', 'sibling', 'parent', 'internal', 'external', 'builtin', 'object', 'type'],
        pathGroupsExcludedImportTypes: ['react'],
        alphabetize: { order: 'asc', caseInsensitive: true },
      },
    ],
    'space-before-function-paren': ['error', 'never'],
    'prettier/prettier': 'error',
  },
}
