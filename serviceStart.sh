#!/bin/bash
ps -fe|grep "SmsCallback_Go -addr" |grep -v grep
if [ $? -ne 0 ]
then
	echo "start server ...."
	#echo $cmd_1
	#`$cmd_1`
	#`$cmd_2`
    SmsCallback_Go -addr=":8001" >> /tmp/logs/SmsCallback_Go.log 2>&1 &
    SmsCallback_Go -addr=":8002" >> /tmp/logs/SmsCallback_Go.log 2>&1 &
	echo "finish SmsCallback_Go on 8001,8002 port"

else
	i=1
	echo "runing....."
	pids=`pgrep -f "SmsCallback_Go -addr"`
	OLD_IFS="$IFS"
	IFS=" "
	arr=($pids)
	IFS="$OLD_IFS"
	for pid in ${arr[@]}
	do
		kill $pid
		echo "killing ""$pid"" ..."
		sleep 7
		echo "restart server ..."
		Addr=":800"$i
        SmsCallback_Go -addr="$Addr" >> /tmp/logs/SmsCallback_Go.log 2>&1 &
		i=$(($i+1))
		echo "restart finish ..."
	done
fi

