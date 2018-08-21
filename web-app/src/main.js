// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import ServerSocket from './socket'
import store from './store'
import moment from 'moment'
import _ from 'lodash'
import 'jquery'
import 'peity'

import 'bulma/css/bulma.css'
import 'bulma-extensions/bulma-pageloader/dist/css/bulma-pageloader.min.css'
import 'bulma-extensions/bulma-badge/dist/css/bulma-badge.min.css'
import 'bulma-extensions/bulma-switch/dist/css/bulma-switch.min.css'
import 'bulma-extensions/bulma-divider/dist/css/bulma-divider.min.css'
import 'font-awesome/css/font-awesome.min.css'

Vue.config.productionTip = false

const server = new ServerSocket('ws://192.168.15.32:5000', 5000)
server.connect()
server.handleMessage = function (message) {
  switch (message.action) {
    case 'registered_agents':
      store.state.agents = message.content.registered
      break
    case 'update_agent_data':
      store.commit('updateAgentData', message.content)
      break
    case 'execute_task':
      store.commit('updateTask', message.content)
      break
  }
}

Vue.prototype.$server = server
Vue.prototype.$moment = moment
Vue.prototype.$_ = _

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  template: '<App/>',
  components: { App }
})
