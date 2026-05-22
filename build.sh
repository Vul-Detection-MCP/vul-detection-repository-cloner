#!/bin/zsh

docker build -t biggujo/vul-detection-repository-cloner:1.0.1 .
docker push biggujo/vul-detection-repository-cloner:1.0.1
