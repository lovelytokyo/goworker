#!/usr/bin/env bash

for i in `seq 1 $1`
    do
        tmp='{"class":"MyClass","args":["a123456", "__ID__", "http://banner-mtb.dspcdn.com/mtbimg/video/bb5/bb59adddc40801a2f9fa10f2116d4185c56a0213"]}'
        cmd=${tmp//__ID__/z${i}}
        echo $cmd

        redis-cli  RPUSH resque:queue:default "$cmd"

    done

