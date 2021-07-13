#!/bin/bash

NS=${1:-default}

if [[ $NS != "default" ]]; then
  kubectl create namespace "${NS}"
fi

kubectl apply -f ./ --namespace="${NS}"
