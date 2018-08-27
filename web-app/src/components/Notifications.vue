<template>
<div class="level-item">
  <button class="button" aria-haspopup="true" @click="toggleSidenav" :disabled="tasks.length === 0">
    <span class="icon is-small badge is-badge-link" :data-badge="tasks.length">
      <i class="fa fa-bell"></i>
    </span>
  </button>
  <div id="notifications" class="sidenav has-text-grey-dark">
    <notification-item v-for="task in tasks" :key="task.id"></notification-item>
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
    toggleSidenav () {
      if (!this.active) {
        this.openSidenav()
      } else {
        this.closeSidenav()
      }
    },
    openSidenav () {
      this.active = true
      document.getElementById('notifications').style.width = '350px'
      document.getElementById('main').style.marginRight = '350px'
    },
    closeSidenav () {
      this.active = false
      document.getElementById('notifications').style.width = '0'
      document.getElementById('main').style.marginRight = '0'
    },
    clearAllTasks () {
      this.$store.commit('removeAllTasks')
      this.closeSidenav()
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
  top: 40px!important;
}
#main {
  transition: margin-right .5s;
}
.sidenav {
  height: 100%;
  width: 0;
  position: fixed;
  z-index: 31;
  top: 0;
  right: 0;
  overflow-x: hidden;
  overflow-y: auto;
  transition: 0.5s;
  margin-top: 60px;
  border-top-left-radius: 6px;
  background-color: rgba(255,255,255,0.95);
  padding-top: 1rem;
}
</style>
