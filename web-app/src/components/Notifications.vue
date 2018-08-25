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
    <div>
    <div class="dropdown-menu" style="min-width: 23rem;" role="menu">
      <div class="dropdown-content has-background-white-bis has-text-grey-dark" style="padding: 0.5rem;">
        <notification-item v-for="task in tasks" :key="task.id"></notification-item>
      </div>
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

<style>
.dropdown-menu {
  top: 45px!important;
}
</style>
