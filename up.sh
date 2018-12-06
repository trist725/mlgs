#!/bin/sh

pid=`ps -ef | grep mlgs.* | grep -v grep  | head -n 1 | awk '{print $2}'`

if [ -n "$pid" ]
then

    kill -9 ${pid}
    echo "pid killed: " ${pid}

fi

cd bin
( ./mlgs* & )
