#!/bin/bash
#1:
docker network create pro-test
docker run -d \
  -p 9104:9104 \
  --network pro-test  \
  -e DATA_SOURCE_NAME="xxxxxxx:xxxxxxxxxxxx@(rm-xxxxxxxxxxxxxx.mysql.rds.aliyuncs.com:3306)/" \
  prom/mysqld-exporter

