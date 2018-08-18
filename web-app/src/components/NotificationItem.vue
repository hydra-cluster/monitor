<template>
<div class="dropdown-item" >
    <article class="message" :class="status">
        <div class="message-body">
          <button class="delete" @click="removeTask(task.id)" :class="task.status === 'Requested' ? 'is-hidden' : ''"></button>
          <p class="is-size-7">
            {{this.$moment(this.task.start).format('DD-MM-YY HH:mm:ss')}} - <strong>{{task.status}}</strong>
          </p>
          <p>
            Task to <strong>{{this.task.action}}</strong> agent(s): {{this.task.target}}
          </p>
          <p class="has-text-right is-size-7">
            <a>View output message</a>
          </p>
        </div>
    </article>
</div>
</template>

<script>
export default {
  name: 'notification-item',
  computed: {
    task () {
      return this.$store.getters.getTaskByID(this.$vnode.key)
    },
    status () {
      switch (this.task.status) {
        case 'Requested':
          return 'is-info'
        case 'Processing':
          return 'is-warning'
        case 'Done':
          return 'is-primary'
        default:
          return 'is-danger'
      }
    }
  },
  methods: {
    removeTask (id) {
      this.$store.commit('removeTask', id)
    }
  }
}
</script>

<style scoped>
.delete {
  position: absolute;
  top: 10px;
  left: 295px;
}
</style>