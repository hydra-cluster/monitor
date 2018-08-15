// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import ServerSocket from './socket'
import store from './store'

require('bulma/css/bulma.css')
require('bulma-extensions/bulma-pageloader/dist/css/bulma-pageloader.min.css')
require('font-awesome/css/font-awesome.min.css')

Vue.config.productionTip = false

const server = new ServerSocket(5000)
server.connect()
server.handleMessage = function (message) {
  if (message.action === 'registered_agents') {
    store.state.agents = message.content.registered
  }
}

Vue.prototype.$server = server

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  template: '<App/>',
  components: { App }
})
