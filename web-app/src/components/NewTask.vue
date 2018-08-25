<template>
<div class="level-item">
  <a class="button is-link" :class="taskInProgress ? 'is-loading' : ''" @click="openModal" >
    <span class="icon is-small">
      <i class="fa fa-terminal"></i>
    </span>
  </a>
  <div class="modal" :class="active ? 'is-active' : ''">
    <div class="modal-background" @click="closeModal"></div>
    <div class="card" style="padding: 15px;">
      <div class="card-content">
        <div style="border-bottom: 1px solid #dadada; margin-bottom: 10px; padding-bottom: 15px;">
          <label class="label is-size-7">Select Agent(s)</label>
          <div class="field is-size-7" v-for="agent in agents" :key="agent.hostname">
            <input :id="agent.hostname" type="checkbox" :value="agent.hostname" class="switch is-rounded is-small" v-model="checkedAgents" :disabled="agent.status === 'Offline'">
            <label :for="agent.hostname" style="padding-top: 0px;">{{agent.hostname}}</label>
          </div>
        </div>
        <div class="field">
          <div class="control has-text-grey-dark">
            <label class="radio is-size-7">
              <input type="radio" name="action" value="Reboot" v-model="action">
              Reboot
            </label>
            <label class="radio is-size-7">
              <input type="radio" name="action" value="Shutdown" v-model="action">
              Shutdown
            </label>
            <label class="radio is-size-7">
              <input type="radio" name="action" value="custom" v-model="action">
              Custom
            </label>
          </div>
        </div>
        <div class="field" :class="action === 'custom' ? '' : 'is-hidden' ">
          <div class="control">
            <textarea class="textarea is-small" placeholder="Custom Command" v-model="command" rows="2"></textarea>
          </div>
        </div>
        <div class="field">
          <div class="control has-icons-left has-icons-right">
            <input class="input is-small" type="password" placeholder="Password" v-model="password">
            <span class="icon is-small is-left">
              <i class="fa fa-unlock-alt"></i>
            </span>
            <span class="icon is-small is-right">
              <i class="fa fa-check"></i>
            </span>
          </div>
        </div>
        <div class="control">
          <button class="button is-primary is-small" @click="newTask" :disabled="disableSubmit">New task</button>
          <button class="button is-small" @click="closeModal">Cancel</button>
        </div>
      </div>
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
    },
    disableSubmit () {
      var customCommand = true
      if (this.action === 'custom' && this.command === '') {
        customCommand = false
      }
      return !(this.action !== '' && this.password !== '' && customCommand && this.checkedAgents.length > 0)
    }
  },
  methods: {
    openModal () {
      this.active = true
    },
    closeModal () {
      this.checkedAgents = []
      this.action = ''
      this.command = ''
      this.password = ''
      this.active = false
    },
    newTask () {
      const task = {
        id: 'task.' + Math.floor(new Date().valueOf() * Math.random()),
        owner: this.$store.state.clientID,
        action: this.customCommand ? 'Custom command' : this.action.toLowerCase(),
        command: this.customCommand ? this.command : this.action,
        target: this.$_.join(this.checkedAgents, ','),
        status: 'Requested',
        start: this.$moment().format()
      }
      this.$store.commit('updateTask', task)

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
      this.closeModal()
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
