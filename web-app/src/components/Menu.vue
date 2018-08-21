<template>
  <nav class="level is-mobile">
    <div class="level-left">
      <div class="level-item has-text-centered">
        <h1 class="title">HYDRA</h1>
      </div>
    </div>
    <div class="level-right">
      <div class="level-item">
        <a class="button" @click="toggleAllContents">
          <span class="icon is-small">
            <i class="fa fa-2x has-text-primary" :class="openAllContents ? 'fa-chevron-circle-up' : 'fa-chevron-circle-down'"></i>
          </span>
        </a>
      </div>
      <new-task></new-task>
      <notifications></notifications>
    </div>
  </nav>
</template>

<script>
import Notifications from './Notifications'
import NewTask from './NewTask'
export default {
  name: 'cluster-menu',
  components: {
    Notifications,
    NewTask
  },
  computed: {
    openAllContents () {
      return this.$store.state.allAgentsContentOpen
    }
  },
  methods: {
    toggleAllContents () {
      this.$store.state.allAgentsContentOpen = !this.$store.state.allAgentsContentOpen
    }
  }
}
</script>
