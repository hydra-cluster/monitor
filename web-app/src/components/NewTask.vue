<template>
  <div class="level-item">
    <a class="button is-link" :class="taskInProgress ? 'is-loading' : ''" @click="openModal" >
      <span class="icon is-small">
        <i class="fa fa-terminal"></i>
      </span>
    </a>
    <div class="modal" :class="active ? 'is-active' : ''">
      <div class="modal-background" @click="closeModal"></div>
      <div class="card is-size-7">
        <div class="card-content">
          <div>
            <div class="columns">
              <div class="column is-one-third has-text-grey-dark hydra-column-modal">
                <div class="field" v-for="agent in agents" :key="agent.hostname">
                  <input :id="agent.hostname" type="checkbox" :value="agent.hostname" class="switch is-rounded" v-model="checkedAgents">
                  <label :for="agent.hostname">{{agent.hostname}}</label>
                </div>
              </div>
              <div class="column hydra-column-modal">
                <div class="field">
                  <div class="control has-text-grey-dark">
                    <label class="radio">
                      <input type="radio" name="action" value="reboot" v-model="action">
                      Reboot
                    </label>
                    <label class="radio">
                      <input type="radio" name="action" value="shutdown" v-model="action">
                      Shutdown
                    </label>
                    <label class="radio">
                      <input type="radio" name="action" value="custom" v-model="action">
                      Custom
                    </label>
                  </div>
                </div>
                <div class="field" :class="action === 'custom' ? '' : 'is-hidden' ">
                  <div class="control">
                    <textarea class="textarea" placeholder="Custom Command" v-model="command" rows="2"></textarea>
                  </div>
                </div>
                <div class="field">
                  <p class="control has-icons-left">
                    <input class="input" type="password" placeholder="Password" v-model="password">
                    <span class="icon is-small is-left">
                      <i class="fa fa-lock"></i>
                    </span>
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
        <footer class="card-foot" style="padding-left: 15px">
          <button class="button is-small is-success" @click="newTask">Create task</button>
          <button class="button is-small" @click="closeModal">Cancel</button>
        </footer>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'new-task',
  data: function () {
    return {
      active: false,
      checkedAgents: [],
      action: '',
      command: '',
      password: ''
    }
  },
  computed: {
    agents () {
      return this.$store.getters.getAgents
    },
    customCommand () {
      return this.action === 'custom'
    },
    taskInProgress () {
      return this.$store.getters.taskInProgress
    }
  },
  methods: {
    openModal () {
      this.active = true
    },
    closeModal () {
      this.active = false
    },
    newTask () {
      const task = {
        id: 'task.' + Math.floor(new Date().valueOf() * Math.random()),
        owner: this.$store.state.clientID,
        action: this.customCommand ? 'execute custom command at' : this.action.toLowerCase(),
        command: this.customCommand ? this.command : this.action,
        target: this.$_.join(this.checkedAgents, ','),
        status: 'Requested',
        start: this.$moment().format()
      }
      this.$store.commit('updateTask', task)
      this.active = false

      // Emit message to the server
      const message = {
        to: 'server',
        from: task.owner,
        action: 'execute_task',
        content: {
          password: this.password,
          task: task
        }
      }
      this.$server.send(message)
    }
  }
}
</script>

<style>
.checkbox + .checkbox {
    margin-left: .5em;
}
.hydra-column-modal {
  padding: 15px;
}
</style>
