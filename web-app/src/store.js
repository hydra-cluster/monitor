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
    invalidPassword: false,
    taskInProgress: false,
    agentsContentVisible: []
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
    },
    getTaskByID: (state) => (id) => {
      return state.tasks.find(task => task.id === id)
    },
    taskInProgress (state) {
      return state.taskInProgress
    },
    getAgentContentVisible: (state) => (hostname) => {
      return state.agentsContentVisible.find(agent => agent.hostname === hostname)
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
      state.tasks.unshift(task)
      if (task.status === 'Requested' || task.status === 'Processing') {
        state.taskInProgress = true
      } else {
        state.taskInProgress = false
      }
    },
    removeTask (state, id) {
      const taskIndex = _.findIndex(state.tasks, {'id': id})
      if (taskIndex !== -1) {
        state.tasks.splice(taskIndex, 1)
      }
    },
    removeAllTasks (state) {
      state.tasks = []
      state.taskInProgress = false
    },
    setAgentContentVisible (state, agent) {
      const agentIndex = _.findIndex(state.agentsContentVisible, {'hostname': agent.hostname})
      if (agentIndex !== -1) {
        state.agentsContentVisible.splice(agentIndex, 1)
      }
      state.agentsContentVisible.push(agent)
    },
    toogleAgentsContentVisible (state, show) {
      _.forEach(state.agentsContentVisible, function (value) {
        value.show = show
      })
    }
  }
})
