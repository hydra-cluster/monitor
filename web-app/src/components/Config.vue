<template>
<div class="level-item">
  <div class="dropdown is-right" :class="active ? 'is-active' : ''">
    <div class="dropdown-trigger">
      <button class="button" aria-haspopup="true" @click="toggleDropDown">
        <span class="icon is-small">
          <i class="fa fa-lg has-text-primary fa-cog"></i>
        </span>
      </button>
    </div>
    <div class="dropdown-menu" id="dropdown-config" role="menu">
      <div class="dropdown-content">
        <a href="#" class="dropdown-item" @click="toggleAgentsCard(true)">
          Open all Agents
        </a>
        <a href="#" class="dropdown-item" @click="toggleAgentsCard(false)">
          Close all Agents
        </a>
        <a href="#" class="dropdown-item is-disabled">
          Configurations
        </a>
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: 'config',
  data: function () {
    return {
      active: false
    }
  },
  methods: {
    toggleAgentsCard (action) {
      this.$store.commit('toogleAgentsContentVisible', action)
      this.active = false
    },
    toggleDropDown () {
      this.active = !this.active
    }
  }
}
</script>

<style>
.dropdown-menu {
  min-width: 5rem;
}
.is-disabled{
  pointer-events: none;
  opacity: .65;
}
</style>
