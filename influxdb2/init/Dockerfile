FROM influxdb:2.0

# Install jq for JSON parsing
RUN apt-get update && apt-get install -y jq

# Copy the initialization script
COPY ./* /docker-entrypoint-initdb.d/

# Ensure the script has execution permissions
RUN chmod +x /docker-entrypoint-initdb.d/token-gen.sh