module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  plugins: ["perfectionist", "unused-imports", "@typescript-eslint", "prettier"],
  parser: "@typescript-eslint/parser",
  extends: [
    "airbnb",
    "airbnb-typescript",
    "airbnb/hooks",
    "prettier",
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react/recommended",
    "plugin:react-hooks/recommended",
    "plugin:prettier/recommended",
  ],
  parserOptions: {
    sourceType: "module",
    ecmaVersion: "latest",
    ecmaFeatures: { jsx: true },
    project: "tsconfig.json",
  },
  settings: {
    "import/resolver": {
      typescript: {
        project: "./tsconfig.json",
      },
    },
  },
  /**
    - 0 ~ 'off'
    - 1 ~ 'warn'
    - 2 ~ 'error'
   */
  rules: {
    // general
    "no-alert": 0,
    camelcase: 0,
    radix: 0,
    "no-bitwise": 0,
    "no-console": 0,
    "no-unused-vars": 0,
    "no-nested-ternary": 0,
    "no-param-reassign": 0,
    "no-underscore-dangle": 0,
    "no-restricted-exports": 0,
    "no-plusplus": 0,
    "no-promise-executor-return": 0,
    "import/extensions": [
      "error",
      "ignorePackages",
      {
        ts: "never",
        tsx: "never",
      },
    ],
    "import/prefer-default-export": 0,
    "import/no-extraneous-dependencies": 0,
    "import/no-cycle": 0,
    "array-callback-return": 0,
    "consistent-return": 0,
    "prefer-destructuring": [1, { object: true, array: false }],
    // typescript
    "@typescript-eslint/no-shadow": 0,
    "@typescript-eslint/naming-convention": 0,
    "@typescript-eslint/no-use-before-define": 0,
    "@typescript-eslint/consistent-type-exports": 1,
    "@typescript-eslint/consistent-type-imports": 1,
    "@typescript-eslint/no-unused-vars": [1, { varsIgnorePattern: "^_", argsIgnorePattern: "^_" }],
    // react
    "react/display-name": 0,
    "react/no-children-prop": 0,
    "react/react-in-jsx-scope": 0,
    "react/no-array-index-key": 0,
    "react/require-default-props": 0,
    "react/jsx-props-no-spreading": 0,
    "react/function-component-definition": 0,
    "react/jsx-no-duplicate-props": [1, { ignoreCase: false }],
    "react/jsx-no-useless-fragment": [1, { allowExpressions: true }],
    "react/no-unstable-nested-components": [1, { allowAsProps: true }],
    "react/no-unused-prop-types": 0,
    "react/prop-types": 0,
    // jsx-a11y
    "jsx-a11y/anchor-is-valid": 0,
    "jsx-a11y/control-has-associated-label": 0,
    // unused imports
    "unused-imports/no-unused-imports": 0,
    "unused-imports/no-unused-vars": [
      0,
      { vars: "all", varsIgnorePattern: "^_", args: "after-used", argsIgnorePattern: "^_" },
    ],
    // prettier
    "prettier/prettier": [2, { singleQuote: false, jsxBracketSameLine: false, arrowParens: "avoid" }],
    // perfectionist
    "perfectionist/sort-exports": [1, { order: "asc", type: "line-length" }],
    "perfectionist/sort-named-imports": [1, { order: "asc", type: "line-length" }],
    "perfectionist/sort-named-exports": [1, { order: "asc", type: "line-length" }],
    "no-restricted-imports": [
      "error",
      {
        patterns: [
          {
            group: ["../"],
            message: "親ディレクトリからの相対インポートは禁止です。@/common などのエイリアスを使用してください。",
          },
        ],
      },
    ],
  },
  "include": [
    "src/**/*",
    ".eslintrc.cjs"
  ]
};
