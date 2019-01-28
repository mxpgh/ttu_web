#!/bin/sh

ps -ef | grep ttu_web | grep -v grep
if [ $? -ne 0 ]
then
	echo "start ttu_web ......"
	nohup ./ttu_web > ttu_web.log 2>&1 & 
else
	echo "ttu_web is running ......"
fi


