#!/usr/bin/env bash
cd ..
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ncgo
mv ncgo ~/go/bin