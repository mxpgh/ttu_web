#!/bin/sh

pid=$(pidof ttu_web)
if [ $? -eq 0 ] 
then
	kill -9 $pid 
fi

rm -rf ttu_web.log
echo "ttu_web stopped"


