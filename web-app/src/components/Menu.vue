<template>
  <nav class="navbar is-fixed-top level is-mobile">
    <div class="level-left">
      <div class="level-item has-text-centered" style="margin-top: 10px; margin-left: 10px;">
        <h1 class="title hydra-title">hydra</h1>
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

<style>
@import url('https://fonts.googleapis.com/css?family=Audiowide|Orbitron');

.hydra-title {
  font-family: 'Orbitron', cursive;
  font-size: 42px;
  font-style: normal;
  font-weight: 200;
  text-shadow: 3px 3px 8px #181818;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  padding: 10px, 0 0 10px;
}
</style>

