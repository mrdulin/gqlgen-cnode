import webpack from 'webpack';
import path from 'path';
import CleanWebpackPlugin from 'clean-webpack-plugin';
import HtmlWebpackPlugin from 'html-webpack-plugin';

const cwd = process.cwd();

const src = path.resolve(cwd, 'src');
const dist = path.resolve(cwd, 'dist');
const port = 3000;

const config: webpack.Configuration = {
  entry: {
    app: src,
    vendors: ['react', 'react-dom', 'react-router-dom', 'react-apollo']
  },
  output: {
    path: dist,
    filename: '[name].js',
    publicPath: '/',
    pathinfo: true
  },
  resolve: {
    alias: {
      gqlMod: path.resolve(src, 'graphql'),
      services: path.resolve(src, 'services')
    },
    extensions: ['', '']
  },
  devtool: 'source-map',
  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/,
        include: [src],
        use: 'ts-loader'
      },
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader']
      },
      {
        test: /\.(graphql|gql)$/,
        exclude: /node_modules/,
        use: 'graphql-tag/loader'
      }
    ]
  },
  plugins: [
    new CleanWebpackPlugin(dist),
    new HtmlWebpackPlugin({
      template: src + '/index.html'
    })
  ]
};

export default config;
