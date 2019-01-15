#!/bin/bash
pids=`pgrep -f "SmsCallback_Go -addr"`
OLD_IFS="$IFS"
IFS=" "
arr=($pids)
IFS="$OLD_IFS"
for pid in ${arr[@]}
do
	kill $pid
done


