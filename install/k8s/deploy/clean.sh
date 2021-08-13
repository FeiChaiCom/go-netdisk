#!/bin/bash

kubectl get all

NS=${1:-default}

echo "namespace: ${NS}"

kubectl delete deployment --all -n="${NS}"
kubectl delete replicaset --all -n="${NS}"
kubectl delete service --all -n="${NS}"
kubectl delete secret --all -n="${NS}"
kubectl delete configmap --all -n="${NS}"
kubectl delete pod --all -n="${NS}"
kubectl delete namespace "${NS}"
