module.exports = {
  publicPath: './',
  configureWebpack: {
    externals: {
      // https://github.com/apertureless/vue-chartjs/issues/124
      moment: 'moment',
    },
  },
  devServer: {
    proxy: {
      '^/oHelp': {
        target: 'http://192.168.20.162:9762',
      },
    },
  },
}
