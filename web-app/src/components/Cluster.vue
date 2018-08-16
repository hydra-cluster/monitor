<template>
    <div class="container">
      <agent v-for="agent in agents" :key="agent.hostname"></agent>
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
