const themeDir = __dirname + '/../../';
const postcssImport = require('postcss-import');
const postcssCustomMedia = require('postcss-custom-media');
const postcssNestedModule = require('postcss-nested');
const postcssNested = postcssNestedModule.default || postcssNestedModule;

module.exports = {
  plugins: [
    postcssImport({
        path: [themeDir]
    }),
    postcssCustomMedia({
        path: [themeDir]
    }),
    postcssNested({
        path: [themeDir]
    }),
  ]
}
