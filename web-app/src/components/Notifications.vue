<template>
<div class="level-item">
  <div class="dropdown is-right" :class="active ? 'is-active' : ''">
    <div class="dropdown-trigger">
      <button class="button" aria-haspopup="true" @click="toggleDropDown" :disabled="tasks.length === 0">
        <span class="icon is-small badge is-badge-link" :data-badge="tasks.length">
          <i class="fa fa-bell"></i>
        </span>
      </button>
    </div>
    <div class="dropdown-menu" role="menu">
      <div class="dropdown-content">
        <notification-item v-for="task in tasks" :key="task.id"></notification-item>
        <hr class="dropdown-divider">
        <a class="dropdown-item" @click="clearAllTasks">
          <span class="icon is-medium">
            <i class="fa fa-lg fa-trash"></i>
          </span>
          Clear all tasks
        </a>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import NotificationItem from './NotificationItem'
export default {
  name: 'notifications',
  components: {
    NotificationItem
  },
  data: function () {
    return {
      active: false
    }
  },
  computed: {
    tasks () {
      return this.$store.getters.getTasks
    }
  },
  methods: {
    toggleDropDown () {
      if (this.tasks.length > 0) {
        this.active = !this.active
      }
    },
    clearAllTasks () {
      this.$store.commit('removeAllTasks')
      this.active = false
    }
  },
  watch: {
    tasks: function (value) {
      if (value.length <= 0) {
        this.active = false
      }
    }
  }
}
</script>

<style scope='local'>
.dropdown-menu {
    min-width: 21rem;
}
</style>
