import Vue from 'vue'
import Vuex from 'vuex'
import _ from 'lodash'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    agents: [],
    tasks: [],
    clientID: '',
    server: {
      connected: false,
      attempts: 0
    },
    invalidPassword: false
  },
  getters: {
    getServer (state) {
      return state.server
    },
    getAgents (state) {
      return _.sortBy(state.agents, ['hostname'])
    },
    getAgentByHostname: (state) => (hostname) => {
      return state.agents.find(agent => agent.hostname === hostname)
    },
    getTasks (state) {
      return state.tasks
    }
  },
  mutations: {
    updateAgentData (state, agent) {
      const oldAgentIndex = _.findIndex(state.agents, {'hostname': agent.hostname})
      if (oldAgentIndex !== -1) {
        state.agents.splice(oldAgentIndex, 1)
      }
      state.agents.push(agent)
    },
    updateTask (state, task) {
      const taskIndex = _.findIndex(state.tasks, {'id': task.id})
      if (taskIndex !== -1) {
        state.tasks.splice(taskIndex, 1)
      }
      state.tasks.push(task)
    },
    removeTask (state, id) {
      const taskIndex = _.findIndex(state.tasks, {'id': id})
      if (taskIndex !== -1) {
        state.tasks.splice(taskIndex, 1)
      }
    },
    removeAllTasks (state) {
      state.tasks = []
    }
  }
})
