#!/bin/sh

cd `dirname $0`/..
wd=`pwd`


docker-compose -f ${wd}/docker/memcached/docker-compose.yml down -v
docker-compose -f ${wd}/docker/memcached/docker-compose.yml up