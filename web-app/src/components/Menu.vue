<template>
  <nav class="level is-mobile">
    <div class="level-left">
      <div class="level-item has-text-centered">
        <h1 class="title">HYDRA</h1>
      </div>
    </div>
    <div class="level-right">
      <task></task>
      <notifications></notifications>
    </div>
  </nav>
</template>

<script>
import Notifications from './Notifications'
import Task from './Task'
export default {
  name: 'cluster-menu',
  components: {
    Notifications,
    Task
  }
}
</script>
