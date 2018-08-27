<template>
<article class="media">
  <div class="media-left">
    <span class="icon">
      <i class="fa fa-lg" :class="status"></i>
    </span>
  </div>
  <div class="media-content is-size-7">
    <div class="content">
      <p>
        <strong>{{this.task.action}}</strong> <small>{{formatDate(task.start)}}</small>
        <br>
        Agent(s): {{this.task.target}}
      </p>       
      <progress class="progress is-small hydra-progress" :class="invalidPassword ? 'is-hidden' : ''" :value="totalOutputs" :max="totalAgents"></progress>
      <p 
        class="has-text-right output-link" 
        :class="invalidPassword ? 'is-hidden' : ''" 
        @click="toggleOutput">{{viewOrHide}} output
      </p>
      <p :class="invalidPassword ? '' : 'is-hidden'">Task not executed, <strong>{{task.status}}</strong></p>
      <div :class="viewOutput ? '': 'is-hidden'">
        <p v-for="result in task.agentsOutput" :key="result.hostname">
          <strong>{{result.hostname}}</strong> <small>{{formatDate(result.end)}}</small>
          <br>
          {{result.output}}
        </p>
      </div>
    </div>
  </div>
  <div class="media-right">
    <span 
      class="icon has-text-grey" 
      :class="task.status === 'Requested' ? 'is-hidden' : ''" 
      style="cursor: pointer;"
      @click="removeTask(task.id)"> 
      <i class="fa fa-times-circle"></i>
    </span>
  </div>
</article>
</template>

<script>
export default {
  name: 'notification-item',
  data: function () {
    return {
      viewOutput: false
    }
  },
  computed: {
    task () {
      return this.$store.getters.getTaskByID(this.$vnode.key)
    },
    totalOutputs () {
      return this.$_.isEmpty(this.task.agentsOutput) ? 0 : this.task.agentsOutput.length
    },
    totalAgents () {
      return this.$_.split(this.task.target, ',').length
    },
    status () {
      switch (this.task.status) {
        case 'Requested':
          return 'fa-check-circle has-text-info'
        case 'Processing':
          return 'fa-check-circle has-text-warning'
        case 'Done':
          return 'fa-check-circle has-text-primary'
        default:
          return 'fa-check-circle has-text-danger'
      }
    },
    viewOrHide () {
      return this.viewOutput ? 'Hide' : 'View'
    },
    invalidPassword () {
      return this.task.status === 'Invalid password'
    }
  },
  methods: {
    removeTask (id) {
      this.$store.commit('removeTask', id)
    },
    formatDate (date) {
      return this.$moment(date).format('DD-MM-YY HH:mm:ss')
    },
    toggleOutput () {
      this.viewOutput = !this.viewOutput
    }
  }
}
</script>

<style>
.output-link {
  margin-top: 8px;
  cursor: pointer;
}
.hydra-progress {
  margin: 0px!important; 
  height: 3px!important;
}
.sidenav .media-left {
  padding: 25px 0 0 5px;
}
</style>
