<template>
<div class="container is-fluid" style="margin: 0px;">
  <div class="columns is-multiline" :class="agents.length < 5 ? 'is-centered' : ''">
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
