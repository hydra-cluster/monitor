<template>
<section class="hero is-dark is-fullheight">
  <div class="pageloader is-info" :class="server.connected ? '' : 'is-active'" id="hydra-loader">
    <span class="title">Connecting to Cluster Server... {{server.attempts}}</span>
  </div>
  <div class="hero-head">
    <cluster-menu></cluster-menu>
  </div>
  <div class="hero-body">
    <cluster></cluster>
  </div>
  <div class="hero-foot">
    <div class="content has-text-centered">
      <p>Hydra Raspberry Pi Cluster Monitor - <strong>v1.0</strong></p>
    </div>
  </div>
</section>
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
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
