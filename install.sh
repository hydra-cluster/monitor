#!/bin/bash
echo "[Hydra Cluster Monitor] Installing"

#copy python lib
sudo rm /usr/local/bin/execCommand.py
cd $GOPATH/src/github.com/hydra-cluster/monitor/lib
sudo cp -n execCommand.py /usr/local/bin/

echo "   Downloading go dependencies..."

go get -v -t -d ./...

echo "   Building Monitor Agent..."

#build and install server and agent executables
sudo rm /usr/local/bin/monitor-agent
cd $GOPATH/src/github.com/hydra-cluster/monitor/cmd/agent
go build -ldflags "-X main.libFolder=/usr/local/bin/" -o monitor-agent
sudo cp -n monitor-agent /usr/local/bin/

sudo rm /etc/init.d/hydra-monitor-agent
sudo cp -n hydra-monitor-agent /etc/init.d
sudo chmod +x /etc/init.d/hydra-monitor-agent

echo "   Building Monitor Server..."

sudo rm /usr/local/bin/monitor-server
cd $GOPATH/src/github.com/hydra-cluster/monitor/cmd/server
go build -ldflags "-X main.libFolder=/usr/local/bin/" -o monitor-server
sudo cp -n monitor-server /usr/local/bin/

sudo rm /etc/init.d/hydra-monitor-server
sudo cp -n hydra-monitor-server /etc/init.d
sudo chmod +x /etc/init.d/hydra-monitor-server

echo "[Hydra Cluster Monitor] Installed"
