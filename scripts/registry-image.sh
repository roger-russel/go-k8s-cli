#!/bin/bash

version=$1

minikube docker-env
eval $(minikube -p minikube docker-env)
docker build -t go-k8s-cli:$version -f k8s/Dockerfile .