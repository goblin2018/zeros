#!/usr/bin/env bash

tables=$2
modeldir=.

host=127.0.0.1
port=3306
dbname=$1
username=root
passwd=xy2089

echo "开始创建数据库 $dbname 的表 $2"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}" -dir="${modeldir}" --cache=true 