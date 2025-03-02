#!/bin/bash
set -e

# Wait for Elasticsearch to be ready
echo "Waiting for Elasticsearch..."
until curl -s -o /dev/null "http://elasticsearch:9200"; do
  sleep 5
done

echo "Elasticsearch is up. Creating users..."

# Create Kibana user
curl -X POST -u elastic:$ELASTIC_PASSWORD "http://elasticsearch:9200/_security/user/kibana_system" -H "Content-Type: application/json" -d '{
  "password": "kibanapassword",
  "roles": ["kibana_system"]
}'

# Create APM user
curl -X POST -u elastic:$ELASTIC_PASSWORD "http://elasticsearch:9200/_security/user/apm_system" -H "Content-Type: application/json" -d '{
  "password": "apmpassword",
  "roles": ["apm_user"]
}'

echo "Users created successfully."
