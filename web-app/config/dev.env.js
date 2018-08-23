'use strict'

const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  SOCKET_SERVER_URL: '"ws://192.168.15.32:5000"'
})
