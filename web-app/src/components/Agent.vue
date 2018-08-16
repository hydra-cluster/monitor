<template>
    <div class="tile">
        {{agent.hostname}}
    </div>
</template>

<script>
export default {
  name: 'agent',
  computed: {
    agent () {
      return this.$store.getters.getAgentByHostname(this.$vnode.key)
    }
  }
}
</script>
