#! /usr/bin/env bash

go build -o resumefilter_$(go env GOOS)_$(go env GOARCH) .