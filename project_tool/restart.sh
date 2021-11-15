#!/usr/bin/env bash
name="staking_server"
pid=`ps -ef|grep ${name}|grep -v "grep"| awk '{print $2}'`
if [ ! -n "$pid" ];then
  echo "${name} not running"
  nohup ./${name} >> prog.log 2>&1  &
else
  echo "old pid : ${pid}"
  kill -9 ${pid}
  nohup ./${name} >> prog.log 2>&1 &
fi
pid2=`ps -ef|grep ${name}|grep -v "grep"| awk '{print $2}'`
echo "new pid is : ${pid2}"