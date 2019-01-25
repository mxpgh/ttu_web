#!/bin/sh
ps -ef | grep ttu_web | grep -v grep | awk '{print $2}' | xargs kill -9

rm -rf ttu_web.log

