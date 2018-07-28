#!/bin/bash
echo "[Hydra Cluster Monitor] Installing"

#copy python lib
cd $GOPATH/src/github.com/hydra-cluster/monitor/lib
cp -n execCommand.py $GOPATH/bin/

echo "Building Monitor Agent..."

#build and install server and agent executables
cd $GOPATH/src/github.com/hydra-cluster/monitor/cmd/agent
go build -ldflags "-X main.libFolder=/hydra/storage/local/go/bin/" -o monitor-agent
cp -n monitor-agent $GOPATH/bin/

echo "Building Monitor Server..."

cd $GOPATH/src/github.com/hydra-cluster/monitor/cmd/server
go build -ldflags "-X main.libFolder=/hydra/storage/local/go/bin/" -o monitor-server
cp -n monitor-server $GOPATH/bin/

echo "[Hydra Cluster Monitor] Installed"
