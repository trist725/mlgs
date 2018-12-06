#!/bin/sh

kill -9 $(ps -ef | grep mlgs.* | grep -v grep  | head -n 1 | awk '{print $2}')


cd bin
(./mlgs* &)
