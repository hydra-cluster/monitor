import store from './store'

const uuid = 'web' + Math.floor(new Date().valueOf() * Math.random()) + '.hydra'
const urlParams = '/ws?id=' + uuid + '&mode=web'

export default class ServerSocket {
  constructor (serverURL, timeToReconnect) {
    this.url = serverURL + urlParams
    this.reconnectInterval = timeToReconnect
  }

  send (content) {
    console.log(content)
    this.socket.send(JSON.stringify(content))
  }

  connect () {
    var self = this
    this.socket = new WebSocket(this.url)
    store.state.clientID = uuid

    // Connection stablished
    this.socket.onopen = function () {
      self.socket.send(
        JSON.stringify({
          action: 'registered_agents',
          to: 'server',
          from: uuid,
          content: ''
        })
      )
      store.state.server.connected = true
      console.log('Connected to the server')
    }

    // Receiving message
    this.socket.onmessage = function (msg) {
      const message = JSON.parse(msg.data)
      self.handleMessage(message)
    }

    // Connection closed
    this.socket.onclose = function (event) {
      store.state.server.connected = false
      self.reconnect()
    }
  }

  reconnect () {
    console.log('Reconnecting in ' + this.reconnectInterval / 1000 + 's')
    store.state.server.attempts++
    setTimeout(this.connect.bind(this), this.reconnectInterval)
  }

  handleMessage (message) {}
}
