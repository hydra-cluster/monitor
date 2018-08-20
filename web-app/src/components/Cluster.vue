<template>
<div class="container is-fluid">
  <div class="columns is-multiline is-centered">
    <agent v-for="agent in agents" :key="agent.hostname"></agent>
  </div>
</div>
</template>

<script>
import Agent from './Agent'
export default {
  components: {
    Agent
  },
  computed: {
    agents () {
      return this.$store.getters.getAgents
    }
  }
}
</script>

<style scoped>
.is-fluid {
  margin: 0px;
}
</style>
