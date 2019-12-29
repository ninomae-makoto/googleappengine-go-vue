module.exports = {
  publicPath: './',
  // https://cli.vuejs.org/config/
  chainWebpack: config => {
    // ライブラリの外部化 htmlで読み込む必要がある
    config.externals({
      "axios": "axios",
      "vue": "Vue",
      'element-ui': 'ElementUI'
    })
  },
  pages: {
    index: {
      // entry for the page
      entry: 'src/main.ts',
      // the source template
      template: 'public/index.html',
      // output as dist/index.html
      filename: 'index.html',
      // when using title option,
      // template title tag needs to be <title><%= htmlWebpackPlugin.options.title %></title>
      title: '',
      // chunks to include on this page, by default includes
      // extracted common chunks and vendor chunks.
      // chunks: ['chunk-vendors', 'chunk-common', 'index']
    }
  }
  // output: {
  //   path: path.resolve(__dirname, './'),
  //   publicPath: '/',
  //   filename: 'build.js'
  // },
}
