module.exports = {
  devtool: 'source-map',
  entry: {
    'user/login': './src/js/user/login.js',
    'user/register': './src/js/user/register.js'
  },
  output: {
    filename: '[name].js'
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader'
        }
      }
    ]
  },
  optimization: {
    minimize: true
  }
}
