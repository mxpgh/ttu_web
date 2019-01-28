#!/bin/sh
basepath=$(cd `dirname $0`; pwd)
cd $basepath/src
go build -o ../ttu_web 
chmod a+x ../ttu_web


