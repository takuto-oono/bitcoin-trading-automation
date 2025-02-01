const config = {
  singleQuote: false,
  arrowParens: "avoid",
  printWidth: 120,
  tabWidth: 2,
  useTabs: false,
  semi: true,
  quoteProps: "as-needed",
  trailingComma: "es5",
  bracketSpacing: true,
  bracketSameLine: false,
  requirePragma: false,
  insertPragma: false,
  htmlWhitespaceSensitivity: "css",
  vueIndentScriptAndStyle: false,
  endOfLine: "lf",
  embeddedLanguageFormatting: "auto",
  singleAttributePerLine: false,
  plugins: ["prettier-plugin-organize-attributes", "prettier-plugin-css-order"],
};

export default config;
