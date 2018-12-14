#!/bin/sh
dir="/var/jenkins_home/workspace/{jobName}"
cd $dir/k8s/build/docker
docker build -f  ./php/build/Dockerfile -t {jobName} .

cd $dir/k8s/build/docker/nginx/build
docker build -t test-nginx .