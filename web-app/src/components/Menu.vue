<template>
  <nav class="level is-mobile">
    <div class="level-left">
      <div class="level-item has-text-centered" style="margin-top: 10px; margin-left: 10px;">
        <h1 class="title">hydra</h1>
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

<style scoped>
@import url('https://fonts.googleapis.com/css?family=Audiowide|Orbitron');

.title {
  font-family: 'Orbitron', cursive;
  font-size: 42px;
  font-style: normal;
  font-weight: 200;
  text-shadow: 1px 1px 8px #333333;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  padding: 10px, 0 0 10px;
}
</style>

