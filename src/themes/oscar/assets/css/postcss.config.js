const themeDir = __dirname + '/../../';

module.exports = {
  plugins: [
    require('postcss-import')({
        path: [themeDir]
    }),
    require('postcss-custom-media')({
        path: [themeDir]
    }),
    require('postcss-nested')({
        path: [themeDir]
    }),
  ]
}
