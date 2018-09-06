module.exports = {
  lintOnSave: false,
  configureWebpack: {
    resolve: {
      extensions: ['.js', '.json', '.vue'],
    },
  },
    devServer: {
        disableHostCheck: true,
    },
};
