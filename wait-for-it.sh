#!/usr/bin/env bash
# wait-for-it.sh script

# Example usage: ./wait-for-it.sh host:port -- <command>
# Waits until the MySQL server at `host:port` is available before running the specified command.

host=$1
port=$2
shift 2
cmd="$@"

until nc -z $host $port; do
  echo "Waiting for $host:$port to be available..."
  sleep 1
done

echo "$host:$port is available"
exec $cmd
