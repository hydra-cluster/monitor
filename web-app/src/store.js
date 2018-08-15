import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    agents: [],
    server: {
      connected: false,
      attempts: 0
    }
  },
  getters: {
    getServer (state) {
      return state.server
    }
  }
})
