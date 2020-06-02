let proxyHost = process.env.PROXY_HOST;
let proxyPort = process.env.SERVER_PORT;

if (!proxyHost) {
  proxyHost = 'localhost';
}

if (!proxyPort) {
  proxyPort = '8080';
}

module.exports = {
  css: {
    loaderOptions: {
      sass: {
        prependData: `@import "@/variables.sass";`,
      },
    },
  },
  devServer: {
    proxy: {
      '/api': {
        target: `http://${proxyHost}:${proxyPort}`,
        changeOrigin: true,
        pathRewrite: {
          '^/api': '/',
        },
      },
    },
  },
};
