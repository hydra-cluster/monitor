#!/bin/bash

#copy python 
cp -n ../lib/execCommand.py /$GOPATH/bin/

#build and install server and agent executables
cd /$GOPATH/src/github.com/hydra-cluster/monitor/cmd/agent
go build -o monitor-agent
cp -n monitor-agent /$GOPATH/bin/

cd /$GOPATH/src/github.com/hydra-cluster/monitor/cmd/server
go build -o monitor-server
cp -n monitor-server /$GOPATH/bin/

echo "Hydra Cluster Monitor - Installed"
