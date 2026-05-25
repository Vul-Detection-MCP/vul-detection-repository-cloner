#!/bin/zsh

docker build --platform linux/amd64,linux/arm64 -t biggujo/vul-detection-repository-cloner:1.0.1 .
docker push biggujo/vul-detection-repository-cloner:1.0.1
