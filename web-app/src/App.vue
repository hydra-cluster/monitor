<template>
<div class="hero is-hydra is-primary is-fullheight is-bold">
  <div class="pageloader is-link" :class="server.connected ? '' : 'is-active'" id="hydra-loader">
    <span class="title">Connecting to Cluster Server... {{server.attempts}}</span>
  </div>
  <div class="hero-head">
    <cluster-menu></cluster-menu>
  </div>
  <div class="hero-body">
    <cluster></cluster>
  </div>
  <div class="hero-foot">
    <p class="content has-text-centered is-size-7 has-text-white">
      Hydra Raspberry Pi Cluster Monitor - <strong>v1.0</strong>
    </p>
  </div>
</div>
</template>

<script>
import ClusterMenu from './components/Menu'
import Cluster from './components/Cluster'

export default {
  name: 'app',
  components: {
    ClusterMenu,
    Cluster
  },
  computed: {
    server () {
      return this.$store.getters.getServer
    }
  }
}
</script>

<style>
body {
  margin: 0;
}

#app {
  font-family: -apple-system,"Helvetica Neue",Helvetica,"Segoe UI",Arial,sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

.hero-head {
  margin: 10px 10px 0 10px;
}
</style>
