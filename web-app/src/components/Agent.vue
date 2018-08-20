<template>
<div class="column is-4-widescreen is-6-desktop is-6-tablet is-12-mobile">
  <div class="card has-background-black-ter has-text-grey">
    <header class="card-header" @click="toggleContent" style="cursor: pointer;">
      <p class="card-header-title">
        <span class="has-text-white-ter">
          <span class="icon">
            <i class="fa" :class="contentOpen ? 'fa-chevron-up' : 'fa-chevron-down'"></i>
          </span>
          head01.hydra
        </span>
      </p>
      <p class="card-header-icon">
        <span class="is-size-7" style="padding-right: 5px;">Online</span>
        <br>
        <span class="icon">
          <i class="fa fa-lg fa-server has-text-primary"></i>
        </span>
      </p>
    </header>
    <div class="card-content" :class="contentOpen ? '' : 'is-hidden'">
      <div> <!-- CPU -->
        <div class="columns is-mobile"> 
          <div class="column is-2 has-text-white-ter">
            CPU
          </div>
          <div class="column is-3-desktop is-4-mobile is-size-7">
            <span class="icon is-size-6">
              <i class="fa fa-thermometer-half"></i>
            </span>
            56.7c
          </div>
          <div class="column is-7-desktop is-6-mobile is-size-7">
            <span>Load average: 45%</span>
            <progress class="progress is-small is-link" style="margin: 0px; height: 3px!important;" value="45" max="100"></progress>
          </div>
        </div>
      </div>
      <div> <!-- RAM -->
        <div class="group has-text-white-ter">
          Memory RAM
        </div>
        <div class="columns is-mobile is-size-7">
          <div class="column is-3">
            <span class="donut-peity" data-peity='{ "fill": ["#08A5E1", "#2A88C5", "#0058A1", "#222D33"],  "innerRadius": 18, "radius": 28 }' style="display: none;">31,24,46,0</span>
          </div>
          <div class="column is-3">
              <small><i class="fa fa-fw fa-circle" style="color:#08A5E1"></i> Using</small>
              <h5>5079 MB</h5>
              <p>31%</p>
          </div>
          <div class="column is-3">
              <small><i class="fa fa-fw fa-circle" style="color:#2A88C5"></i> Cache</small>
              <h5>3789 MB</h5>
              <p>24%</p>
          </div>
          <div class="column is-3">
              <small><i class="fa fa-fw fa-circle" style="color:#0058A1"></i> Free</small>
              <h5>3591 MB</h5>
              <p>46%</p>
          </div>
        </div>
      </div>
      <div> <!-- Volumes -->
        <div class="group has-text-white-ter">
          Volume status
        </div>
        <div>
          <div class="columns is-mobile is-size-7">
            <div class="column is-3">
              <p class="is-size-4 is-size-6-mobile has-text-white-ter">
                32
                <span class="is-size-6 is-size-7-mobile">GB</span></p>
              <p>SD Card</p>
            </div>
            <div class="column is-3">
              <span class="donut-peity" data-peity='{ "fill": ["#2D99DC", "#222D33"], "innerRadius": 18, "radius": 28 }' style="display: none;">2/7</span>
            </div>
            <div class="column is-3">
                <small><i class="fa fa-fw fa-circle" style="color:#2D99DC"></i> Used</small>
                <h5>5079 MB</h5>
                <p>31%</p>
            </div>
            <div class="column is-3">
                <small><i class="fa fa-fw fa-circle" style="color:#222D33"></i> Free</small>
                <h5>3789 MB</h5>
                <p>24%</p>
            </div>
          </div>
        </div>
        <div>
          <div class="columns is-mobile is-size-7">
            <div class="column is-3">
              <p class="is-size-4 is-size-6-mobile has-text-white-ter">
                128
                <span class="is-size-6 is-size-7-mobile">GB</span>
              </p>
              <p>Storage</p>
            </div>
            <div class="column is-3">
              <span class="donut-peity" data-peity='{ "fill": ["#2D99DC", "#222D33"], "innerRadius": 18, "radius": 28 }' style="display: none;">2/7</span>
            </div>
            <div class="column is-3">
                <small><i class="fa fa-fw fa-circle" style="color:#2D99DC"></i> Used</small>
                <h5>5079 MB</h5>
                <p>31%</p>
            </div>
            <div class="column is-3">
                <small><i class="fa fa-fw fa-circle" style="color:#222D33"></i> Free</small>
                <h5>3789 MB</h5>
                <p>24%</p>
            </div>
          </div>
        </div>
      </div>
      <div class="group"> <!-- System OS -->
        <div class="columns is-mobile is-size-7">
          <div class="column is-4">
            <p>Distro</p>
            <p class="has-text-white-ter is-size-6">Raspbian strech lite</p>
          </div>
          <div class="column is-4">
            <p>Kernel</p>
            <p class="has-text-white-ter is-size-6">4.14 ARMv7</p>
          </div>
          <div class="column is-4">
            <p>Model</p>
            <p class="has-text-white-ter is-size-6">RPi 3 model B+</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: 'agent',
  data: function () {
    return {
      contentOpen: false
    }
  },
  computed: {
    agent () {
      return this.$store.getters.getAgentByHostname(this.$vnode.key)
    }
  },
  methods: {
    toggleContent () {
      this.contentOpen = !this.contentOpen
    }
  },
  mounted: function () {
    $('.donut-peity').peity('donut')
  }
}
</script>

<style scoped>
.card {
  border-radius: 6px;
}
.group {
  border-top: 1px solid #333131;
  padding-top: 10px;
  margin: 10px 0 10px 0;
}
</style>