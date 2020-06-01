let proxy = process.env.PROXY_HOST;

if (!proxy) {
  proxy = 'localhost';
}

module.exports = {
  devServer: {
    proxy: {
      '/select': {
        target: `http://${proxy}:8080`,
        changeOrigin: true,
      },
    },
  },
};
