<template>
  <nav class="level is-mobile">
    <div class="level-left">
      <div class="level-item has-text-centered">
        <h1 class="title">HYDRA</h1>
      </div>
    </div>
    <div class="level-right">
      <config></config>
      <new-task></new-task>
      <notifications></notifications>
    </div>
  </nav>
</template>

<script>
import Config from './Config'
import NewTask from './NewTask'
import Notifications from './Notifications'

export default {
  name: 'cluster-menu',
  components: {
    Config,
    NewTask,
    Notifications
  }
}
</script>
