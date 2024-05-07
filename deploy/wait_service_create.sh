#!/bin/sh

set -e

mysql_host="$1"
redis_host="$2"
shift
shift
cmd="$@"

until mysqladmin ping -h"$mysql_host" --silent; do
  echo 'waiting for mysql...'
  sleep 5
done

until echo PING | redis-cli -h $redis_host; do
  echo 'waiting for redis...'
  sleep 5
done

exec $cmd