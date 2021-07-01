module.exports = {
  indexPath: "../templates/index.html",
  outputDir: "../static",
  publicPath: process.env.PUBLIC_PATH,
  assetsDir: "",
  devServer: {
    port: 8080,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:5000',
        // target: 'http://127.0.0.1:6010',
        changeOrigin: true,
        pathRewrite: {
          '^/api': '/api'
        }
      }
    }
  },
  lintOnSave: false,
  transpileDependencies: [
    'vue-echarts',
    'resize-detector'
  ]
}

console.log(module.exports);
console.log(process.env);
