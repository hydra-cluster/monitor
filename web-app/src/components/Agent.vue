<template>
<div class="column is-4-widescreen is-6-desktop is-6-tablet is-12-mobile">
  <div class="card has-background-black-ter has-text-grey">
    <header class="card-header" @click="toggleContent" style="cursor: pointer;">
      <p class="card-header-title">
        <span class="has-text-white-ter">
          <span class="icon">
            <i class="fa" :class="contentVisible ? 'fa-chevron-up' : 'fa-chevron-down'"></i>
          </span>
          {{agent.hostname}}
        </span>
      </p>
      <p class="card-header-icon">
        <span class="is-size-7 has-text-right" style="padding-right: 12px;">{{agent.status}}<br>{{agent.ip}}</span>
        <br>
        <span class="icon">
          <i class="fa fa-2x fa-server" :class="status"></i>
        </span>
      </p>
    </header>
    <div class="card-content" :class="contentVisible ? '' : 'is-hidden'">
      <div> <!-- CPU | Network -->
        <div class="columns is-mobile is-size-7"> 
          <div class="column is-6 has-text-centered">
            <i class="fa fa-bolt is-size-5 has-text-white-ter"></i>
            <span style="margin-right: 10px;">{{cpu.freq}}</span>
            <i class="fa fa-thermometer-half is-size-5" :class="cpu.tempColor"></i>
            <span>{{cpu.tempLabel}}</span>
          </div>
          <div class="column is-6">
            <progress class="progress is-small" :class="cpu.loadColor" style="margin: 0px; height: 3px!important;" :value="cpu.loadPerc" max="100"></progress>
            <span>Load average: {{cpu.loadLabel}}</span>
          </div>
        </div>
      </div>
      <div class="group"> <!-- RAM -->
        <div class="columns is-mobile">
          <div class="column is-6 has-text-white-ter has-text-centered">
            <span class="donut-peity" :data-peity="mem.ramDonut" style="display: none;">{{mem.ramUsed}},{{mem.ramFree}}</span>
            <p style="padding-bottom: 15px;">RAM</p>
          </div>
          <div class="column is-6 has-text-white-ter has-text-centered border-left">
            <span class="donut-peity" :data-peity="mem.swapDonut" style="display: none;">{{mem.swapUsed}},{{mem.swapFree}}</span>
            <p style="padding-bottom: 15px;">Swap</p>
          </div>
        </div>
        <div class="columns is-mobile is-size-7">
          <div class="column is-3 has-text-centered">
              <small><i class="fa fa-fw fa-circle" :style="mem.ramColor"></i> Used</small>
              <h5>{{mem.ramUsed}} {{mem.unit}}</h5>
              <p>{{mem.ramUsedPerc.toFixed(1)}}%</p>
          </div>
          <div class="column is-3 has-text-centered">
              <small><i class="fa fa-fw fa-circle" style="color:#222D33"></i> Free</small>
              <h5>{{mem.ramFree}} {{mem.unit}}</h5>
              <p>{{mem.ramFreePerc.toFixed(1)}}%</p>
          </div>
          <div class="column is-3 has-text-centered border-left">
              <small><i class="fa fa-fw fa-circle" :style="mem.swapColor"></i> Used</small>
              <h5>{{mem.swapUsed}} {{mem.unit}}</h5>
              <p>{{mem.swapUsedPerc.toFixed(1)}}%</p>
          </div>
          <div class="column is-3 has-text-centered">
              <small><i class="fa fa-fw fa-circle" style="color:#222D33"></i> Free</small>
              <h5>{{mem.swapFree}} {{mem.unit}}</h5>
              <p>{{mem.swapFreePerc.toFixed(1)}}%</p>
          </div>
        </div>
      </div>
      <div class="group" v-for="vol in volumes" :key="vol.label"> <!-- Volumes -->
        <div class="columns is-mobile is-size-7">
          <div class="column is-3">
            <p>{{vol.label}}</p>
            <p class="is-size-4 is-size-6-mobile has-text-white-ter">
              {{vol.total}}
              <span class="is-size-6 is-size-7-mobile">GB</span>
            </p>
          </div>
          <div class="column is-3">
            <span class="donut-peity" :data-peity="vol.donut" style="display: none;">{{vol.used}},{{vol.free}}</span>
          </div>
          <div class="column is-3 has-text-centered">
              <small><i class="fa fa-fw fa-circle" :style="vol.color"></i> Used</small>
              <h5>{{vol.used}}GB</h5>
              <p>{{vol.usedPerc}}%</p>
          </div>
          <div class="column is-3 has-text-centered">
              <small><i class="fa fa-fw fa-circle" style="color:#222D33"></i> Free</small>
              <h5>{{vol.free}}GB</h5>
              <p>{{vol.freePerc}}%</p>
          </div>
        </div>
      </div>
      <div class="group"> <!-- Distro, Kernel, Model -->
        <div class="columns is-mobile">
          <div class="column is-6">
            <p class="is-size-7">Distro | Kernel</p>
            <p class="has-text-white-ter is-size-7-mobile is-size-6-tablet">{{agent.distro}} | {{agent.kernel}}</p>
          </div>
          <div class="column is-6">
            <p class="is-size-7">Model</p>
            <p class="has-text-white-ter is-size-7-mobile is-size-6-tablet">{{agent.model}}</p>
          </div>
        </div>
      </div>
      <div class="group"> <!-- System OS -->
        <div class="columns is-mobile">
          <div class="column">
            <p class="is-size-7">Network</p>
            <div class="is-size-7-mobile is-size-6-tablet">
              <p class="has-text-white-ter" v-for="net in agent.network_interfaces" :key="net.ip">
                {{net.name}}: {{net.ip}} ({{net.mac_addr}})
              </p>
            </div>
          </div>
        </div>
        <div class="columns is-mobile" style="padding-top: 5px;">
          <div class="column">
            <p class="is-size-7">Agent uptime</p>
            <p class="has-text-white-ter is-size-7-mobile is-size-6-tablet">{{agentUptime}}</p>
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
  computed: {
    agent () {
      return this.$store.getters.getAgentByHostname(this.$vnode.key)
    },
    status () {
      if (this.agent.status === 'Online') {
        return 'has-text-primary'
      } else if (this.agent.status === 'Offline') {
        return 'has-text-danger'
      }
      return 'has-text-warning'
    },
    cpu () {
      const t = this.$_.find(this.agent.params, ['label', 'CPU Temp'])
      const l = this.$_.find(this.agent.params, ['label', 'CPU Load'])
      const f = this.$_.find(this.agent.params, ['label', 'CPU Freq'])
      var loadClass = 'is-link'
      if (parseFloat(l.value) >= parseFloat(l.danger)) {
        loadClass = 'is-danger'
      } else if (parseFloat(l.value) >= parseFloat(l.warning)) {
        loadClass = 'is-warning'
      }
      var tempClass = 'has-text-link'
      if (parseFloat(t.value) >= parseFloat(t.danger)) {
        tempClass = 'has-text-danger'
      } else if (parseFloat(t.value) >= parseFloat(t.warning)) {
        tempClass = 'has-text-warning'
      }
      return {
        tempLabel: t.value + t.unit,
        loadLabel: l.value,
        loadPerc: parseFloat(l.value) > 1 ? 100 : parseFloat(l.value) * 100,
        loadColor: loadClass,
        tempColor: tempClass,
        freq: (f.value / 1000 / 1000).toFixed(1) + 'G' + f.unit
      }
    },
    mem () {
      const rt = this.$_.find(this.agent.params, ['label', 'RAM Total'])
      const ru = this.$_.find(this.agent.params, ['label', 'RAM Used'])
      const free = (parseFloat(rt.value) - parseFloat(ru.value))
      const freePerc = (free / parseFloat(rt.value) * 100)
      const usedPerc = (parseFloat(ru.value) / parseFloat(rt.value) * 100)
      var ramColor = '#2D99DC'
      if (usedPerc >= parseFloat(rt.danger)) {
        ramColor = '#FF3860'
      } else if (usedPerc >= parseFloat(rt.warning)) {
        ramColor = '#FFDD57'
      }
      const st = this.$_.find(this.agent.params, ['label', 'SWAP Total'])
      const su = this.$_.find(this.agent.params, ['label', 'SWAP Used'])
      const swapFree = (parseFloat(st.value) - parseFloat(su.value))
      const swapFreePerc = parseFloat(su.value) === 0 ? 100 : (free / parseFloat(st.value) * 100)
      const swapUsedPerc = (parseFloat(su.value) / parseFloat(st.value) * 100)
      var swapColor = '#2D99DC'
      if (swapUsedPerc >= parseFloat(st.danger)) {
        swapColor = '#FF3860'
      } else if (swapUsedPerc >= parseFloat(st.warning)) {
        swapColor = '#FFDD57'
      }
      return {
        unit: ru.unit,
        ramUsed: ru.value,
        ramUsedPerc: usedPerc,
        ramFree: free,
        ramFreePerc: freePerc,
        ramColor: 'color:' + ramColor,
        ramDonut: '{ "fill": ["' + ramColor + '", "#222D33"], "innerRadius": 15, "radius": 25 }',
        swapUsed: su.value,
        swapUsedPerc: swapUsedPerc,
        swapFree: swapFree,
        swapFreePerc: swapFreePerc,
        swapColor: 'color:' + swapColor,
        swapDonut: '{ "fill": ["' + swapColor + '", "#222D33"], "innerRadius": 15, "radius": 25 }'
      }
    },
    volumes () {
      var volumes = []
      const sdct = this.$_.find(this.agent.params, ['label', 'SD Card Total'])
      const sdcu = this.$_.find(this.agent.params, ['label', 'SD Card Used'])
      const sdcfree = (sdct.value - sdcu.value)
      const cardUsedPerc = (sdcu.value / sdct.value * 100)
      var cardColor = '#2D99DC'
      if (cardUsedPerc >= parseFloat(sdct.danger)) {
        cardColor = '#FF3860'
      } else if (cardUsedPerc >= parseFloat(sdct.warning)) {
        cardColor = '#FFDD57'
      }
      var card = {
        label: 'SD Card',
        total: (sdct.value / 1024 / 1024).toFixed(0),
        used: (sdcu.value / 1024 / 1024).toFixed(1),
        usedPerc: cardUsedPerc.toFixed(1),
        free: (sdcfree / 1024 / 1024).toFixed(1),
        freePerc: (sdcfree / sdct.value * 100).toFixed(1),
        color: 'color:' + cardColor,
        donut: '{ "fill": ["' + cardColor + '", "#222D33"], "innerRadius": 15, "radius": 25 }'
      }
      volumes.push(card)
      const stgt = this.$_.find(this.agent.params, ['label', 'Storage Total'])
      const stgu = this.$_.find(this.agent.params, ['label', 'Storage Used'])
      const stgfree = (stgt.value - stgu.value)
      const storageUsedPerc = (stgu.value / stgt.value * 100)
      var storageColor = '#2D99DC'
      if (storageUsedPerc >= parseFloat(stgt.danger)) {
        storageColor = '#FF3860'
      } else if (storageUsedPerc >= parseFloat(stgt.warning)) {
        storageColor = '#FFDD57'
      }
      var storage = {
        label: 'Storage',
        total: (stgt.value / 1024 / 1024).toFixed(0),
        used: (stgu.value / 1024 / 1024).toFixed(1),
        usedPerc: storageUsedPerc.toFixed(1),
        free: (stgfree / 1024 / 1024).toFixed(1),
        freePerc: (stgfree / stgt.value * 100).toFixed(1),
        color: 'color:' + storageColor,
        donut: '{ "fill": ["' + storageColor + '", "#222D33"], "innerRadius": 15, "radius": 25 }'
      }
      if (storage.total > 0) {
        volumes.push(storage)
      }
      return volumes
    },
    agentUptime () {
      return this.agent.status === 'Offline' ? 'N/D' : this.$moment(this.agent.last_connection_at).fromNow(true)
    },
    contentVisible () {
      return this.$store.getters.getAgentContentVisible(this.agent.hostname).show
    }
  },
  methods: {
    toggleContent () {
      this.$store.commit('setAgentContentVisible', {hostname: this.agent.hostname, show: !this.contentVisible})
    }
  },
  mounted: function () {
    $('.donut-peity').peity('donut')
  },
  created: function () {
    this.$store.commit('setAgentContentVisible', {hostname: this.agent.hostname, show: true})
  }
}
</script>

<style scoped>
.card {
  border-radius: 6px;
}
.card-content {
  padding: 1rem;
}
.group {
  border-top: 1px solid #333131;
  padding-top: 10px;
  margin: 10px 0 10px 0;
}
.columns {
  margin: 0px;
}
.column {
  padding-top: 0px;
  padding-bottom: 0px;
}
.border-left {
  border-left: 1px solid #333131;
}
</style>