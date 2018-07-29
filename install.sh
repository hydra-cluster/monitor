#!/bin/bash

echo "[Hydra Cluster Monitor] Installing"

general() {
    sudo mkdir -p /var/log/hydra/

    echo "   Configuring python lib..."

    #copy python lib
    sudo mkdir -p /usr/local/var/lib/hydra/monitor/
    sudo rm -f /usr/local/var/lib/hydra/monitor/execCommand.py
    cd $GOPATH/src/github.com/hydra-cluster/monitor/lib
    sudo cp -n execCommand.py /usr/local/var/lib/hydra/monitor/

    echo "   Downloading go dependencies..."

    cd $GOPATH/src/github.com/hydra-cluster/monitor
    go get -v -t -d ./...
}

agent() {
    echo "   Building Monitor Agent..."

    #build and install server and agent executables
    sudo rm -f /usr/local/bin/monitor-agent
    cd $GOPATH/src/github.com/hydra-cluster/monitor/cmd/agent
    go build -ldflags "-X main.libFolder=/usr/local/var/lib/hydra/monitor/" -o monitor-agent
    sudo cp -n monitor-agent /usr/local/bin/

    sudo rm -f /etc/init.d/hydra-monitor-agent
    sudo cp -n hydra-monitor-agent /etc/init.d
    sudo chmod +x /etc/init.d/hydra-monitor-agent

    sudo update-rc.d hydra-moitor-agent defaults 
}

server() {
    echo "   Building Monitor Server..."

    sudo rm -f /usr/local/bin/monitor-server
    cd $GOPATH/src/github.com/hydra-cluster/monitor/cmd/server
    go build -ldflags "-X main.libFolder=/usr/local/var/lib/hydra/monitor/" -o monitor-server
    sudo cp -n monitor-server /usr/local/bin/

    sudo rm -f /etc/init.d/hydra-monitor-server
    sudo cp -n hydra-monitor-server /etc/init.d
    sudo chmod +x /etc/init.d/hydra-monitor-server

    sudo update-rc.d hydra-monitor-server defaults 
}

case "$1" in
	'all')
		general
        agent
        server
		;;
	'agent')
        general
		agent
		;;
	'server')
		general
        server
		;;
	*)
		echo "Usage: $0 { all | agent | server }"
		exit 1
		;;
esac
 
echo "[Hydra Cluster Monitor] Installed"

exit 0

