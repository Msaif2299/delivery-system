#!/bin/bash
if [ ! -f /var/lib/influxdb2/meta/meta.db ]; then
    echo "Creating token"
    mkdir -p /var/lib/influxdb2/secrets
    BUCKET_ID=$(influx bucket list | grep "$DOCKER_INFLUXDB_INIT_BUCKET" | awk '{print $1}')
    NEW_TOKEN=$(influx auth create --org="$DOCKER_INFLUXDB_INIT_ORG" \
                    --description "API token" \
                    --read-bucket "$BUCKET_ID" \
                    --write-bucket "$BUCKET_ID" \
                    --token "$DOCKER_INFLUXDB_INIT_ADMIN_TOKEN" \
                    --json | jq -r '.token')
    echo "$NEW_TOKEN" > /var/lib/influxdb2/secrets/api_token.txt
    mkdir -p /var/lib/influxdb2/shared_config
    echo "$NEW_TOKEN" > /var/lib/influxdb2/shared_config/influxdb_config.txt
    export INFLUXDB_TOKEN="$NEW_TOKEN"
else
    echo "InfluxDB already initialized"
    export INFLUXDB_TOKEN=$(cat /var/lib/influxdb2/secrets/api_token.txt)
fi