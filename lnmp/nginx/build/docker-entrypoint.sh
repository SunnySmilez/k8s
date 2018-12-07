#!/bin/bash

# 配置文件变更
rm -f /etc/nginx/conf.d/default.conf
#mkdir -p /data/logs
nginx -g "daemon off;"