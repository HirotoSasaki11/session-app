#!/bin/sh

cd `dirname $0`/..
wd=`pwd`


docker-compose -f ${wd}/docker/radies/docker-compose.yml down -v
docker-compose -f ${wd}/docker/radies/docker-compose.yml up